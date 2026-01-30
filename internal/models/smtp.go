package models

import (
	"time"

	"github.com/google/uuid"
)

type SMTPConfig struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Host         string     `json:"host"`
	Port         int        `json:"port"`
	Username     *string    `json:"username,omitempty"`
	PasswordEnc  *string    `json:"-"`
	UseTLS       bool       `json:"use_tls"`
	UseSSL       bool       `json:"use_ssl"`
	IsDefault    bool       `json:"is_default"`
	IsActive     bool       `json:"is_active"`
	LastUsedAt   *time.Time `json:"last_used_at,omitempty"`
	FailureCount int        `json:"failure_count"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
