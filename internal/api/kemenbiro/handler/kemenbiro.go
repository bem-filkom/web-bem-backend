package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"net/url"
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

func (h *KemenbiroHandler) GetAllKemenbiros() fiber.Handler {
	return func(c *fiber.Ctx) error {
		kemenbiros, err := h.s.GetAllKemenbiros(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(map[string]interface{}{
			"kemenbiros": kemenbiros,
		})
	}
}

func (h *KemenbiroHandler) GetKemenbiroByAbbreviation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req kemenbiro.GetKemenbiroByAbbreviationRequest
		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		queryUnescapedAbbr, err := url.QueryUnescape(req.Abbreviation)
		if err != nil {
			return response.ErrUnprocessableEntity
		}
		req.Abbreviation = queryUnescapedAbbr

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

		queryUnescapedAbbr, err := url.QueryUnescape(req.AbbreviationAsID)
		if err != nil {
			return response.ErrUnprocessableEntity
		}
		req.AbbreviationAsID = queryUnescapedAbbr

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

		queryUnescapedAbbr, err := url.QueryUnescape(req.Abbreviation)
		if err != nil {
			return response.ErrUnprocessableEntity
		}
		req.Abbreviation = queryUnescapedAbbr

		if err := h.s.DeleteKemenbiro(c.Context(), &req); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
