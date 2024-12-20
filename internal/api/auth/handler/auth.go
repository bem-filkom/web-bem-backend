package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/auth"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) LoginUB() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req auth.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		res, err := h.s.LoginUB(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}
