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

func (h *KemenbiroHandler) GetKemenbiroByAbbreviation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req kemenbiro.GetKemenbiroByAbbreviationRequest
		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
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
