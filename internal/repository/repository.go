// Package Repository - Repository
// provides the data access layer for the application
package repository

import (
	"github.com/versenilvis/sendio/internal/database"
	"github.com/versenilvis/sendio/internal/database/dbgen"
)

type Registry struct {
	Q     *dbgen.Queries
	Email EmailRepository
	SMTP  SMTPRepository
	Auth  AuthRepository
}

// NewRegistry creates a new instance of Registry
func NewRegistry(db *database.DB) *Registry {
	q := dbgen.New(db.Pool)
	return &Registry{
		Q:     q,
		Email: NewEmailRepository(q),
		SMTP:  NewSMTPRepository(q),
		Auth:  NewAuthRepository(q),
	}
}
