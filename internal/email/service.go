// Package email - Service
// call repo to store email record in db
// enqueue email send task
// update email status
package email

import (
	"context"
	"fmt"

	"github.com/versenilvis/sendio/internal/database/dbgen"
	"github.com/versenilvis/sendio/internal/models"
	"github.com/versenilvis/sendio/internal/queue"
	"github.com/versenilvis/sendio/internal/repository"
)

type Service interface {
	SendEmail(ctx context.Context, params SendEmailParams) (*dbgen.Email, error)
}

type service struct {
	repo  *repository.Registry
	queue *queue.QueueClient
}

// SendEmailParams uses pointer here to detect if user want to fill their name blank or don't want to fill it
type SendEmailParams struct {
	APIKeyID     [16]byte
	FromEmail    string
	FromName     *string
	ToEmail      string
	ToName       *string
	Subject      string
	TemplateID   *string
	TemplateVars map[string]any
	BodyHTML     *string
	BodyText     *string
}

func NewService(repo *repository.Registry, queue *queue.QueueClient) Service {
	return &service{
		repo:  repo,
		queue: queue,
	}
}

// SendEmail creates a new email record and enqueues it for sending
func (s *service) SendEmail(ctx context.Context, params SendEmailParams) (*dbgen.Email, error) {
	// NOTE: convert UUID string to [16]byte if needed, but here it's already [16]byte from context

	email, err := s.repo.Email.Create(ctx, repository.CreateEmailParams{
		APIKeyID:     params.APIKeyID,
		FromEmail:    params.FromEmail,
		FromName:     params.FromName,
		ToEmail:      params.ToEmail,
		ToName:       params.ToName,
		Subject:      params.Subject,
		TemplateID:   params.TemplateID,
		TemplateVars: params.TemplateVars,
		BodyHTML:     params.BodyHTML,
		BodyText:     params.BodyText,
		Status:       models.EmailPending,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create email record: %w", err)
	}

	if err := s.queue.EnqueueEmailSend(email.ID.String()); err != nil {
		// TASK: Update status to failed? or let it be pending for a cron retry?
		// TODO:
		// for now, return error
		// we will make a cron job to scan email that is on pending for more than 5 mins
		return nil, fmt.Errorf("failed to enqueue email task: %w", err)
	}

	if err := s.repo.Email.UpdateStatus(ctx, email.ID.String(), models.EmailQueued, "queued_at"); err != nil {
		return nil, fmt.Errorf("failed to update status to queued: %w", err)
	}

	return email, nil
}
