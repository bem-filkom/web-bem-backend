package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro/service"
	"github.com/bem-filkom/web-bem-backend/internal/middleware"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"time"
)

type KemenbiroHandler struct {
	s service.IKemenbiroService
}

func NewKemenbiroHandler(s service.IKemenbiroService) *KemenbiroHandler {
	return &KemenbiroHandler{s: s}
}

func (h *KemenbiroHandler) Start(router fiber.Router) {
	router = router.Group("/v2/kemenbiros")
	router.Post("",
		middleware.Authenticate(),
		middleware.RequireSuperAdmin(),
		timeout.NewWithContext(h.CreateKemenbiro(), 5*time.Second),
	)
	router.Get("/:id", timeout.NewWithContext(h.GetKemenbiroByID(), 5*time.Second))
	router.Get("", timeout.NewWithContext(h.GetKemenbiroWithQuery(), 5*time.Second))
	router.Patch("/:abbreviationAsID",
		middleware.Authenticate(),
		middleware.RequireRole(entity.RoleBemMember),
		timeout.NewWithContext(h.UpdateKemenbiro(), 5*time.Second),
	)
	router.Delete("/:abbreviation",
		middleware.Authenticate(),
		middleware.RequireSuperAdmin(),
		timeout.NewWithContext(h.DeleteKemenbiro(), 5*time.Second),
	)
}
