package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *KemenbiroHandler) CreateKemenbiro() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req kemenbiro.CreateKemenbiroRequest
		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		kemenbiroObj, err := h.s.CreateKemenbiro(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
			"kemenbiro": kemenbiroObj,
		})
	}
}

func (h *KemenbiroHandler) GetModifiableKemenbiro() fiber.Handler {
	return func(c *fiber.Ctx) error {
		kemenbiros, err := h.s.GetModifiableKemenbiros(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(map[string]interface{}{
			"kemenbiros": kemenbiros,
		})
	}
}

func (h *KemenbiroHandler) GetKemenbiroByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req kemenbiro.GetKemenbiroByIDRequest
		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		kemenbiroObj, err := h.s.GetKemenbiroByID(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.JSON(map[string]any{
			"kemenbiro": kemenbiroObj,
		})
	}
}

func (h *KemenbiroHandler) GetKemenbiroWithQuery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req kemenbiro.GetKemenbiroByAbbreviationRequest
		if err := c.QueryParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if req.Abbreviation == "" {
			kemenbiros, err := h.s.GetAllKemenbiros(c.Context())
			if err != nil {
				return err
			}

			return c.JSON(map[string]interface{}{
				"kemenbiros": kemenbiros,
			})
		}

		kemenbiroObj, err := h.s.GetKemenbiroByAbbreviation(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.JSON(map[string]any{
			"kemenbiro": kemenbiroObj,
		})
	}
}

func (h *KemenbiroHandler) UpdateKemenbiro() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req kemenbiro.UpdateKemenbiroRequest

		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if err := h.s.UpdateKemenbiro(c.Context(), &req); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func (h *KemenbiroHandler) DeleteKemenbiro() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req kemenbiro.DeleteKemenbiroRequest
		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		if err := h.s.DeleteKemenbiro(c.Context(), &req); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
