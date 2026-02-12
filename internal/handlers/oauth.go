package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	pgtype "github.com/jackc/pgx/v5/pgtype"
	"github.com/versenilvis/sendio/internal/config"
	"github.com/versenilvis/sendio/internal/database/dbgen"
	"github.com/versenilvis/sendio/internal/oauth"
	"github.com/versenilvis/sendio/internal/repository"
)

type OAuthHandler struct {
	userRepo       repository.UserRepository
	google         oauth.Provider
	github         oauth.Provider
	sessionManager *oauth.SessionManager
	cfg            *config.Config
}

func NewOAuthHandler(userRepo repository.UserRepository, google oauth.Provider, github oauth.Provider, sessionManager *oauth.SessionManager, cfg *config.Config) *OAuthHandler {
	return &OAuthHandler{
		userRepo:       userRepo,
		google:         google,
		github:         github,
		sessionManager: sessionManager,
		cfg:            cfg,
	}
}

// --- Login Redirects ---

func (h *OAuthHandler) GoogleLogin(c fiber.Ctx) error {
	state := h.generateState(c)
	url := h.google.AuthURL(state)
	return c.Redirect().To(url)
}

func (h *OAuthHandler) GitHubLogin(c fiber.Ctx) error {
	state := h.generateState(c)
	url := h.github.AuthURL(state)
	return c.Redirect().To(url)
}

// --- Callbacks ---

func (h *OAuthHandler) GoogleCallback(c fiber.Ctx) error {
	return h.handleCallback(c, h.google, "google")
}

func (h *OAuthHandler) GitHubCallback(c fiber.Ctx) error {
	return h.handleCallback(c, h.github, "github")
}

func (h *OAuthHandler) handleCallback(c fiber.Ctx, provider oauth.Provider, providerName string) error {
	state := c.Query("state")
	cookieState := c.Cookies("oauth_state")
	if state != cookieState {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state"})
	}
	// clear state cookie because it's only valid for one request
	c.ClearCookie("oauth_state")

	code := c.Query("code")
	userInfo, err := provider.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	existingUser, err := h.userRepo.GetUserByEmail(c.Context(), userInfo.Email)
	var userID pgtype.UUID

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// user not found -> create new user
			name := &userInfo.Name
			avatar := &userInfo.AvatarURL
			newUser, err := h.userRepo.CreateUser(c.Context(), userInfo.Email, name, avatar)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
			}
			userID = newUser.ID
		} else {
			// real database error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
		}
	} else {
		// user found -> update info
		userID = existingUser.ID
	}

	err = h.userRepo.UpsertOAuthAccount(c.Context(), dbgen.UpsertOAuthAccountParams{
		UserID:       userID,
		Provider:     providerName,
		ProviderID:   userInfo.ProviderID,
		AccessToken:  pgtype.Text{String: "", Valid: true}, // we might not need to store token yet
		RefreshToken: pgtype.Text{String: "", Valid: true},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to link account"})
	}

	token, err := h.sessionManager.CreateSession(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create session"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   h.cfg.Env == "production",
		SameSite: "Lax",
		Path:     "/",
	})

	// close popup or redirect
	// since we are using a popup, we can close the window and reload parent
	// or just redirect to home if not popup
	return c.SendString("<script>window.opener.location.reload();window.close();</script>")
}

func (h *OAuthHandler) generateState(c fiber.Ctx) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
		Secure:   h.cfg.Env == "production",
	})
	return state
}

// --- Auth Endpoints ---

// Me returns the current user
// it uses the session token from the cookie to find the user
// if no session token is found, it returns unauthorized
//
// /auth/me
func (h *OAuthHandler) Me(c fiber.Ctx) error {
	// retrieve user from context (set by middleware)
	user, ok := c.Locals("user").(*dbgen.GetSessionByTokenHashRow)
	if !ok || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.JSON(fiber.Map{
		"name":      user.Name,
		"email":     user.Email,
		"avatarUrl": user.AvatarUrl,
	})
}

func (h *OAuthHandler) Logout(c fiber.Ctx) error {
	token := c.Cookies("session_token")
	if token != "" {
		h.sessionManager.DeleteSession(c.Context(), token)
	}

	c.ClearCookie("session_token")
	return c.JSON(fiber.Map{"message": "Logged out"})
}
