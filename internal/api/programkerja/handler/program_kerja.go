package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *ProgramKerjaHandler) CreateProgramKerja() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req programkerja.CreateProgramKerjaRequest
		if err := c.BodyParser(&req); err != nil {
			return response.ErrUnprocessableEntity
		}

		programkerjaObj, err := h.s.CreateProgramKerja(c.Context(), &req)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
			"program_kerja": programkerjaObj,
		})
	}
}
