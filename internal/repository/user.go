// Package Repository - User
// convert text to pgtype
package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/versenilvis/sendio/internal/database/dbgen"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (dbgen.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (dbgen.User, error)
	CreateUser(ctx context.Context, email string, name *string, avatarURL *string) (dbgen.User, error)
	UpdateUser(ctx context.Context, id pgtype.UUID, name *string, avatarURL *string) (dbgen.User, error)
	UpsertOAuthAccount(ctx context.Context, params dbgen.UpsertOAuthAccountParams) error
	CreateSession(ctx context.Context, params dbgen.CreateSessionParams) error
	GetSessionByTokenHash(ctx context.Context, tokenHash string) (dbgen.GetSessionByTokenHashRow, error)
	DeleteSessionByTokenHash(ctx context.Context, tokenHash string) error
}

type userRepo struct {
	q *dbgen.Queries
}

func NewUserRepository(q *dbgen.Queries) UserRepository {
	return &userRepo{q: q}
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (dbgen.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}

func (r *userRepo) GetUserByID(ctx context.Context, id pgtype.UUID) (dbgen.User, error) {
	return r.q.GetUserByID(ctx, id)
}

func (r *userRepo) CreateUser(ctx context.Context, email string, name *string, avatarURL *string) (dbgen.User, error) {
	return r.q.CreateUser(ctx, dbgen.CreateUserParams{
		Email:     email,
		Name:      textToPgType(name),
		AvatarUrl: textToPgType(avatarURL),
	})
}

func (r *userRepo) UpdateUser(ctx context.Context, id pgtype.UUID, name *string, avatarURL *string) (dbgen.User, error) {
	return r.q.UpdateUser(ctx, dbgen.UpdateUserParams{
		ID:        id,
		Name:      textToPgType(name),
		AvatarUrl: textToPgType(avatarURL),
	})
}

func (r *userRepo) UpsertOAuthAccount(ctx context.Context, params dbgen.UpsertOAuthAccountParams) error {
	return r.q.UpsertOAuthAccount(ctx, params)
}

func (r *userRepo) CreateSession(ctx context.Context, params dbgen.CreateSessionParams) error {
	return r.q.CreateSession(ctx, params)
}

func (r *userRepo) GetSessionByTokenHash(ctx context.Context, tokenHash string) (dbgen.GetSessionByTokenHashRow, error) {
	return r.q.GetSessionByTokenHash(ctx, tokenHash)
}

func (r *userRepo) DeleteSessionByTokenHash(ctx context.Context, tokenHash string) error {
	return r.q.DeleteSessionByTokenHash(ctx, tokenHash)
}

func textToPgType(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}
