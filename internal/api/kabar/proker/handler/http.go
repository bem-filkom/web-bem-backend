package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker/service"
	"github.com/bem-filkom/web-bem-backend/internal/middleware"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/gofiber/fiber/v2"
)

type KabarProkerHandler struct {
	s service.IKabarProkerService
}

func NewKabarProkerHandler(s service.IKabarProkerService) *KabarProkerHandler {
	return &KabarProkerHandler{s: s}
}

func (h *KabarProkerHandler) Start(router fiber.Router) {
	router = router.Group("/v2/kabars/proker")
	router.Post("",
		middleware.Authenticate(),
		middleware.RequireRole(entity.RoleBemMember),
		h.CreateKabarProker(),
	)
	router.Get("", h.GetKabarProkerByQuery())
}
