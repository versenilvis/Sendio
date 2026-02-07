// Package handlers - Email
package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/versenilvis/sendio/internal/email"
)

type EmailHandler struct {
	emailSvc email.Service
}

func NewEmailHandler(emailSvc email.Service) *EmailHandler {
	return &EmailHandler{
		emailSvc: emailSvc,
	}
}

type SendEmailRequest struct {
	FromEmail    string         `json:"from_email"`
	FromName     *string        `json:"from_name"`
	ToEmail      string         `json:"to_email"`
	ToName       *string        `json:"to_name"`
	Subject      string         `json:"subject"`
	TemplateID   *string        `json:"template_id"`
	TemplateVars map[string]any `json:"template_vars"`
	BodyHTML     *string        `json:"body_html"`
	BodyText     *string        `json:"body_text"`
}

// SendEmail handles the HTTP request for sending an email
// at HTTP/API layer, it doesnt care about the database
// it only cares about the request and response
func (h *EmailHandler) SendEmail(c fiber.Ctx) error {
	var req SendEmailRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.ToEmail == "" || req.Subject == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "to_email and subject are required",
		})
	}

	// get API Key ID from context (set by AuthRequired middleware)
	apiKeyID := fiber.Locals[[16]byte](c, "apiKeyID")
	if apiKeyID == [16]byte{} {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "API Key ID not found or invalid",
		})
	}

	emailData, err := h.emailSvc.SendEmail(c, email.SendEmailParams{
		APIKeyID:     apiKeyID,
		FromEmail:    req.FromEmail,
		FromName:     req.FromName,
		ToEmail:      req.ToEmail,
		ToName:       req.ToName,
		Subject:      req.Subject,
		TemplateID:   req.TemplateID,
		TemplateVars: req.TemplateVars,
		BodyHTML:     req.BodyHTML,
		BodyText:     req.BodyText,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"id":     uuid.UUID(emailData.ID.Bytes).String(),
		"status": "queued",
	})
}
