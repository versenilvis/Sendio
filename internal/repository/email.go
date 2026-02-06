// Package repository - Email
// provides the data access layer for the application
// this file specifically manages email related data operations
package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/versenilvis/sendio/internal/database/dbgen"
	"github.com/versenilvis/sendio/internal/models"
)

// EmailRepository defines the contract for managing email records in the database
type EmailRepository interface {
	GetByID(ctx context.Context, id string) (*models.Email, error)
	UpdateStatus(ctx context.Context, id string, status models.EmailStatus, field string) error
	TryLockForProcessing(ctx context.Context, id string) (bool, error)
}

type emailRepo struct {
	q *dbgen.Queries
}

func NewEmailRepository(q *dbgen.Queries) EmailRepository {
	return &emailRepo{q: q}
}

// GetByID fetches an email record, converting the input string ID into a PostgreSQL UUID type
func (r *emailRepo) GetByID(ctx context.Context, id string) (*models.Email, error) {
	var pgid pgtype.UUID
	if err := pgid.Scan(id); err != nil {
		return nil, fmt.Errorf("invalid uuid format: %w", err)
	}
	// GetEmailByID is the function name from our email.sql file
	// it returns a dbgen.Email struct
	dbEmail, err := r.q.GetEmailByID(ctx, pgid)
	if err != nil {
		return nil, fmt.Errorf("failed to get email by id from database: %w", err)
	}

	return mapDBEmailToModel(dbEmail), nil
}

// UpdateStatus changes the lifecycle state of an email
// the 'field' parameter determines if additional timestamps (sent_at/queued_at) should be set
func (r *emailRepo) UpdateStatus(ctx context.Context, id string, status models.EmailStatus, field string) error {
	var pgid pgtype.UUID
	// check if the id is a valid uuid
	if err := pgid.Scan(id); err != nil {
		return err
	}

	// we need to convert the status type from go to sqlc type
	pgStatus := dbgen.NullEmailStatus{
		EmailStatus: dbgen.EmailStatus(status),
		Valid:       true,
	}

	// simply means that we want to exec which function depends on the field
	switch field {
	case "sent_at":
		// if it is "sent_at", change status and attach sent_at = NOW()
		// means sent email successfully
		return r.q.UpdateEmailSent(ctx, dbgen.UpdateEmailSentParams{
			Status: pgStatus,
			ID:     pgid,
		})
	case "queued_at":
		// if it is "queued_at", change status and attach queued_at = NOW()
		// means email is queued in redis
		return r.q.UpdateEmailQueued(ctx, dbgen.UpdateEmailQueuedParams{
			Status: pgStatus,
			ID:     pgid,
		})
	default:
		return r.q.UpdateEmailStatus(ctx, dbgen.UpdateEmailStatusParams{
			Status: pgStatus,
			ID:     pgid,
		})
	}
}

// TryLockForProcessing implements the "Idempotent Consumer" pattern
// it ensures that even if multiple workers receive the same task, only one can proceed
func (r *emailRepo) TryLockForProcessing(ctx context.Context, id string) (bool, error) {
	var pgid pgtype.UUID
	if err := pgid.Scan(id); err != nil {
		return false, err
	}

	pgStatus := dbgen.NullEmailStatus{
		EmailStatus: dbgen.EmailStatusProcessing,
		Valid:       true,
	}

	// this calls the SQL with: WHERE id = $1 AND status IN ('pending', 'queued')
	result, err := r.q.UpdateEmailStatusAtomic(ctx, dbgen.UpdateEmailStatusAtomicParams{
		Status: pgStatus,
		ID:     pgid,
	})
	if err != nil {
		return false, err
	}

	return result.RowsAffected() > 0, nil
}

// mapDBEmailToModel converts the raw database record into a clean domain model
// this is essential for keeping postgres-specific logic (pgtype) out of the main business logic
func mapDBEmailToModel(dbEmail dbgen.Email) *models.Email {
	email := &models.Email{
		ID:         uuid.UUID(dbEmail.ID.Bytes),
		APIKeyID:   uuid.UUID(dbEmail.ApiKeyID.Bytes),
		FromEmail:  dbEmail.FromEmail,
		Subject:    dbEmail.Subject,
		ToEmail:    dbEmail.ToEmail,
		Status:     models.EmailStatus(dbEmail.Status.EmailStatus),
		RetryCount: int(dbEmail.RetryCount.Int32),
		MaxRetries: int(dbEmail.MaxRetries.Int32),
		CreatedAt:  dbEmail.CreatedAt.Time,
	}

	// handle all optional/nullable database fields
	if dbEmail.IdempotencyKey.Valid {
		email.IdempotencyKey = &dbEmail.IdempotencyKey.String
	}
	if dbEmail.FromName.Valid {
		email.FromName = &dbEmail.FromName.String
	}
	if dbEmail.ToName.Valid {
		email.ToName = &dbEmail.ToName.String
	}
	if dbEmail.BodyHtml.Valid {
		email.BodyHTML = &dbEmail.BodyHtml.String
	}
	if dbEmail.BodyText.Valid {
		email.BodyText = &dbEmail.BodyText.String
	}
	if dbEmail.SentAt.Valid {
		email.SentAt = &dbEmail.SentAt.Time
	}

	// TODO: add more if db expanded

	return email
}
