package worker

import (
	"context"
	"encoding/json"
	"fmt"
	netmail "net/mail"
	"time"

	"github.com/hibiken/asynq"
	"github.com/versenilvis/sendio/internal/logger"
	"github.com/versenilvis/sendio/internal/models"
	"github.com/versenilvis/sendio/internal/queue"
	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
)

// ProcessTask handles email:send tasks from the queue
func (p *Processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload queue.EmailSendPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("could not unmarshal payload: %w", err)
	}

	// ensure only ONE worker can process this email ID
	locked, err := p.repo.Email.TryLockForProcessing(ctx, payload.EmailID)
	if err != nil {
		return fmt.Errorf("failed to lock email for processing: %w", err)
	}
	if !locked {
		logger.Log.Warn("email already being processed or sent, skipping", zap.String("email_id", payload.EmailID))
		return nil
	}

	dbCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	email, err := p.repo.Email.GetByID(dbCtx, payload.EmailID)
	if err != nil {
		return fmt.Errorf("failed to fetch email %s: %w", payload.EmailID, err)
	}

	smtpConfig, err := p.repo.SMTP.GetDefault(dbCtx)
	if err != nil {
		return fmt.Errorf("failed to fetch smtp config: %w", err)
	}

	htmlBody, textBody := p.prepareEmailBody(email)

	msg := p.buildMailMessage(email, htmlBody, textBody)

	if err := p.sendEmail(smtpConfig, msg); err != nil {
		_ = p.repo.Email.UpdateStatus(ctx, email.ID.String(), models.EmailFailed, "")
		logger.Log.Error("failed to send email", zap.Error(err), zap.String("email_id", email.ID.String()))
		return fmt.Errorf("failed to send email: %w", err)
	}

	if err := p.repo.Email.UpdateStatus(ctx, email.ID.String(), models.EmailSent, "sent_at"); err != nil {
		logger.Log.Error("failed to update email status", zap.Error(err))
	}

	logger.Log.Info("email sent successfully", zap.String("email_id", email.ID.String()))
	return nil
}

// prepareEmailBody extracts or renders email body content
func (p *Processor) prepareEmailBody(email *models.Email) (htmlBody, textBody string) {
	if email.TemplateID != nil {
		rendered, err := p.renderer.Render("welcome.html", email.TemplateVars)
		if err == nil {
			htmlBody = rendered
		}
	} else if email.BodyHTML != nil {
		htmlBody = *email.BodyHTML
	}

	if email.BodyText != nil {
		textBody = *email.BodyText
	}

	return htmlBody, textBody
}

// buildMailMessage creates a go-mail message from email data
func (p *Processor) buildMailMessage(email *models.Email, htmlBody, textBody string) *mail.Msg {
	m := mail.NewMsg()

	// use pre-parsed addresses to skip DNS validation
	fromAddr := &netmail.Address{Address: email.FromEmail}
	if email.FromName != nil {
		fromAddr.Name = *email.FromName
	}
	toAddr := &netmail.Address{Address: email.ToEmail}
	if email.ToName != nil {
		toAddr.Name = *email.ToName
	}

	m.FromMailAddress(fromAddr)
	m.ToMailAddress(toAddr)
	m.EnvelopeFromMailAddress(fromAddr)
	m.Subject(email.Subject)

	if htmlBody != "" {
		m.SetBodyString(mail.TypeTextHTML, htmlBody)
	} else if textBody != "" {
		m.SetBodyString(mail.TypeTextPlain, textBody)
	}

	return m
}
