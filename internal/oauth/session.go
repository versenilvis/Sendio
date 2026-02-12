package oauth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/versenilvis/sendio/internal/config"
	"github.com/versenilvis/sendio/internal/database/dbgen"
	"github.com/versenilvis/sendio/internal/repository"
)

type SessionManager struct {
	repo repository.UserRepository
	cfg  *config.Config
}

func NewSessionManager(repo repository.UserRepository, cfg *config.Config) *SessionManager {
	return &SessionManager{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *SessionManager) CreateSession(ctx context.Context, userID pgtype.UUID) (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", fmt.Errorf("failed to generate session token: %w", err)
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	hash := sha256.New()
	hash.Write([]byte(token + s.cfg.Security.SessionSecret))
	tokenHash := hex.EncodeToString(hash.Sum(nil))

	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	err := s.repo.CreateSession(ctx, dbgen.CreateSessionParams{
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create session in db: %w", err)
	}

	return token, nil
}

func (s *SessionManager) ValidateSession(ctx context.Context, token string) (*dbgen.GetSessionByTokenHashRow, error) {
	hash := sha256.New()
	hash.Write([]byte(token + s.cfg.Security.SessionSecret))
	tokenHash := hex.EncodeToString(hash.Sum(nil))

	sessionUser, err := s.repo.GetSessionByTokenHash(ctx, tokenHash)
	if err != nil {
		return nil, err
	}

	return &sessionUser, nil
}

func (s *SessionManager) DeleteSession(ctx context.Context, token string) error {
	hash := sha256.New()
	hash.Write([]byte(token + s.cfg.Security.SessionSecret))
	tokenHash := hex.EncodeToString(hash.Sum(nil))

	return s.repo.DeleteSessionByTokenHash(ctx, tokenHash)
}
