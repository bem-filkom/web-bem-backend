package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/auth"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req auth.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		token, err := h.s.LoginAuthUb(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(map[string]any{
			"access_token": token,
		})
	}
}
