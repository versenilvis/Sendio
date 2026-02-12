package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/versenilvis/sendio/internal/oauth"
)

func SessionAuth(sessionManager *oauth.SessionManager) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Cookies("session_token")
		if token == "" {
			return c.Next()
		}

		user, err := sessionManager.ValidateSession(c.Context(), token)
		if err != nil {
			c.ClearCookie("session_token")
			return c.Next()
		}

		c.Locals("user", user)
		return c.Next()
	}
}

// RequireAuth middleware ensures a user is logged in
func RequireAuth(c fiber.Ctx) error {
	if c.Locals("user") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	return c.Next()
}
