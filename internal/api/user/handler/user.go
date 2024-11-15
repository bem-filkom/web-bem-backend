package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) CreateBemMember() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req user.CreateBemMemberRequest
		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if err := h.s.CreateBemMember(c.Context(), &req); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func (h *UserHandler) UpdateBemMember() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req user.UpdateBemMemberRequest

		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if err := h.s.UpdateBemMember(c.Context(), &req); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func (h *UserHandler) DeleteBemMember() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req user.DeleteBemMemberRequest
		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if err := h.s.DeleteBemMember(c.Context(), &req); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
