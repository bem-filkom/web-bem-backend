package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *KabarProkerHandler) CreateKabarProker() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req proker.CreateKabarProkerRequest
		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		err := h.s.CreateKabarProker(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func (h *KabarProkerHandler) GetKabarProkerByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req proker.GetKabarProkerByIDRequest
		if err := c.ParamsParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		kabarProker, err := h.s.GetKabarProkerByID(c.Context(), req)
		if err != nil {
			return err
		}

		return c.JSON(map[string]interface{}{
			"kabar_proker": kabarProker,
		})
	}
}

func (h *KabarProkerHandler) GetKabarProkerByQuery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req proker.GetKabarProkerByQueryRequest
		if err := c.QueryParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		kabarProkers, paginationRes, err := h.s.GetKabarProkerByQuery(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.JSON(map[string]interface{}{
			"kabar_prokers": kabarProkers,
			"pagination":    paginationRes,
		})
	}
}
