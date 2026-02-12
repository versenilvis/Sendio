package oauth

import (
	"context"
)

type UserInfo struct {
	ProviderID string
	Email      string
	Name       string
	AvatarURL  string
}

type Provider interface {
	AuthURL(state string) string
	Exchange(ctx context.Context, code string) (*UserInfo, error)
}
