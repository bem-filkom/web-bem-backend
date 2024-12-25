package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja/service"
	"github.com/bem-filkom/web-bem-backend/internal/middleware"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/gofiber/fiber/v2"
)

type ProgramKerjaHandler struct {
	s service.IProgramKerjaService
}

func NewProgramKerjaHandler(s service.IProgramKerjaService) *ProgramKerjaHandler {
	return &ProgramKerjaHandler{s: s}
}

func (h *ProgramKerjaHandler) Start(router fiber.Router) {
	router = router.Group("/v2/program-kerjas")
	router.Post("",
		middleware.Authenticate(),
		middleware.RequireRole(entity.RoleBemMember),
		h.CreateProgramKerja(),
	)
	router.Get("/:id", h.GetProgramKerjaByID())
	router.Get("", h.GetProgramKerjasByKemenbiroID())
	router.Patch("/:id",
		middleware.Authenticate(),
		middleware.RequireRole(entity.RoleBemMember),
		h.UpdateProgramKerja(),
	)
}
