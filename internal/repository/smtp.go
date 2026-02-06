// Package repository - SMTP
// it does 2 jobs:
// 1. get raw data type from dbgen
// 2. map them to clean data type
package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/versenilvis/sendio/internal/database/dbgen"
	"github.com/versenilvis/sendio/internal/models"
)

// SMTPRepository defines the operations for managing SMTP configurations
// we use an interface here so that we can easily swap the implementation (e.g., for testing)
// no matter what database we use, we only need it to get the default SMTP configuration
type SMTPRepository interface {
	// GetDefault retrieves the default active SMTP configuration from the database
	GetDefault(ctx context.Context) (*models.SMTPConfig, error)
}

type smtpRepo struct {
	q *dbgen.Queries
}

// NewSMTPRepository creates a new instance of SMTPRepository
func NewSMTPRepository(q *dbgen.Queries) SMTPRepository {
	return &smtpRepo{q: q}
}

// GetDefault fetches the default SMTP configuration by calling the sqlc generated query
// it then converts the database-specific result into a clean model used by the rest of the application
func (r *smtpRepo) GetDefault(ctx context.Context) (*models.SMTPConfig, error) {
	// GetDefaultSMTPConfig is the function name from our smtp.sql file
	// it returns a dbgen.SmtpConfig struct
	dbConfig, err := r.q.GetDefaultSMTPConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get default smtp config: %w", err)
	}

	return mapDBSMTPToModel(dbConfig), nil
}

// mapDBSMTPToModel is a helper function that translates the raw database struct (dbgen.SmtpConfig)
// into our clean application model (models.SMTPConfig)
// this is where we handle the conversion from pgtype (used by sqlc/pgx) to standard Go types
func mapDBSMTPToModel(dbConfig dbgen.SmtpConfig) *models.SMTPConfig {
	config := &models.SMTPConfig{
		ID:           uuid.UUID(dbConfig.ID.Bytes),
		Name:         dbConfig.Name,
		Host:         dbConfig.Host,
		Port:         int(dbConfig.Port),
		UseTLS:       dbConfig.UseTls.Bool,
		UseSSL:       dbConfig.UseSsl.Bool,
		IsDefault:    dbConfig.IsDefault.Bool,
		IsActive:     dbConfig.IsActive.Bool,
		FailureCount: int(dbConfig.FailureCount.Int32),
		CreatedAt:    dbConfig.CreatedAt.Time,
		UpdatedAt:    dbConfig.UpdatedAt.Time,
	}

	if dbConfig.Username.Valid {
		config.Username = &dbConfig.Username.String
	}
	if dbConfig.PasswordEnc.Valid {
		config.PasswordEnc = &dbConfig.PasswordEnc.String
	}
	if dbConfig.LastUsedAt.Valid {
		config.LastUsedAt = &dbConfig.LastUsedAt.Time
	}

	return config
}
