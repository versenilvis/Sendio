package repository

import (
	"context"

	"github.com/versenilvis/sendio/internal/database/dbgen"
)

type AuthRepository interface {
	GetAPIKeyByHash(ctx context.Context, hash string) (dbgen.ApiKey, error)
}

type authRepo struct {
	q *dbgen.Queries
}

func NewAuthRepository(q *dbgen.Queries) AuthRepository {
	return &authRepo{q: q}
}

func (r *authRepo) GetAPIKeyByHash(ctx context.Context, hash string) (dbgen.ApiKey, error) {
	return r.q.GetAPIKeyByHash(ctx, hash)
}
