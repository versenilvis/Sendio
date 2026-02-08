package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/versenilvis/sendio/internal/database/dbgen"
	"github.com/versenilvis/sendio/internal/repository"
)

type Service interface {
	VerifyAPIKey(ctx context.Context, apiKey string) (*dbgen.ApiKey, error)
}

type service struct {
	repo repository.AuthRepository
	salt string
}

func NewService(repo repository.AuthRepository, salt string) Service {
	return &service{
		repo: repo,
		salt: salt,
	}
}

func (s *service) VerifyAPIKey(ctx context.Context, apiKey string) (*dbgen.ApiKey, error) {
	hash := sha256.New()
	hash.Write([]byte(apiKey + s.salt))
	hashedKey := hex.EncodeToString(hash.Sum(nil))

	keyData, err := s.repo.GetAPIKeyByHash(ctx, hashedKey)
	if err != nil {
		return nil, fmt.Errorf("auth service: %w", err)
	}

	return &keyData, nil
}
