package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/versenilvis/sendio/internal/auth"
	"github.com/versenilvis/sendio/internal/logger"
	"go.uber.org/zap"
)

func AuthRequired(authSvc auth.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		// Header format: "Bearer sk_12345"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization format. Use 'Bearer <api_key>'",
			})
		}

		apiKey := parts[1]

		keyData, err := authSvc.VerifyAPIKey(c, apiKey)
		if err != nil {
			logger.Log.Warn("Invalid API Key attempt", zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or inactive API Key",
			})
		}

		// we don't need to define [16]byte type here because Go can automatically detect via keyData.ID.Bytes
		fiber.Locals(c, "apiKeyID", keyData.ID.Bytes)

		return c.Next()
	}
}
