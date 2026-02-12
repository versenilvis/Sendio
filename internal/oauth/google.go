package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/versenilvis/sendio/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleProvider struct {
	config *oauth2.Config
}

func NewGoogle(cfg *config.Config) *GoogleProvider {
	return &GoogleProvider{
		config: &oauth2.Config{
			ClientID:     cfg.OAuth.GoogleClientID,
			ClientSecret: cfg.OAuth.GoogleClientSecret,
			RedirectURL:  cfg.ServerURL + "/auth/callback/google",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (p *GoogleProvider) AuthURL(state string) string {
	return p.config.AuthCodeURL(state)
}

// Exchange code for access token and get user info
func (p *GoogleProvider) Exchange(ctx context.Context, code string) (*UserInfo, error) {
	token, err := p.config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	client := p.config.Client(ctx, token)
	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create userinfo request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// https://developers.google.com/identity/openid-connect/openid-connect#obtainuserinfo
	var googleUser struct {
		Sub     string `json:"sub"` // ID
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := json.Unmarshal(body, &googleUser); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	// remove the postfix "=s..." so we can get full resolution avatar
	avatarURL := regexp.MustCompile(`=s\d+-c$`).ReplaceAllString(googleUser.Picture, "")

	return &UserInfo{
		ProviderID: googleUser.Sub,
		Email:      googleUser.Email,
		Name:       googleUser.Name,
		AvatarURL:  avatarURL,
	}, nil
}
