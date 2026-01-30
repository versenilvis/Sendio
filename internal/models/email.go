// Package models - email
package models

import (
	"time"

	"github.com/google/uuid"
)

type EmailStatus string

const (
	EmailPending    EmailStatus = "pending"
	EmailQueued     EmailStatus = "queued"
	EmailProcessing EmailStatus = "processing"
	EmailSent       EmailStatus = "sent"
	EmailDelivered  EmailStatus = "delivered"
	EmailFailed     EmailStatus = "failed"
	EmailBounced    EmailStatus = "bounced"
	EmailCancelled  EmailStatus = "cancelled"
)

type Email struct {
	ID             uuid.UUID      `json:"id"`
	APIKeyID       uuid.UUID      `json:"api_key_id"`
	IdempotencyKey *string        `json:"idempotency_key,omitempty"`
	FromEmail      string         `json:"from_email"`
	FromName       *string        `json:"from_name,omitempty"`
	ToEmail        string         `json:"to_email"`
	ToName         *string        `json:"to_name,omitempty"`
	Subject        string         `json:"subject"`
	TemplateID     *uuid.UUID     `json:"template_id,omitempty"`
	TemplateVars   map[string]any `json:"template_vars,omitempty"`
	BodyHTML       *string        `json:"body_html,omitempty"`
	BodyText       *string        `json:"body_text,omitempty"`
	Tags           []string       `json:"tags,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Status         EmailStatus    `json:"status"`
	RetryCount     int            `json:"retry_count"`
	MaxRetries     int            `json:"max_retries"`
	CreatedAt      time.Time      `json:"created_at"`
	ScheduledAt    *time.Time     `json:"scheduled_at,omitempty"`
	QueuedAt       *time.Time     `json:"queued_at,omitempty"`
	SentAt         *time.Time     `json:"sent_at,omitempty"`
	DeliveredAt    *time.Time     `json:"delivered_at,omitempty"`
}

type EmailEvent struct {
	ID          uuid.UUID      `json:"id"`
	EmailID     uuid.UUID      `json:"email_id"`
	EventType   string         `json:"event_type"`
	Description *string        `json:"description,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
}

type EmailDeliveryAttempt struct {
	ID           uuid.UUID  `json:"id"`
	EmailID      uuid.UUID  `json:"email_id"`
	AttemptNum   int        `json:"attempt_num"`
	Status       string     `json:"status"`
	SMTPCode     *string    `json:"smtp_code,omitempty"`
	SMTPMessage  *string    `json:"smtp_message,omitempty"`
	ErrorMessage *string    `json:"error_message,omitempty"`
	SMTPConfigID *uuid.UUID `json:"smtp_config_id,omitempty"`
	DurationMs   *int       `json:"duration_ms,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}
