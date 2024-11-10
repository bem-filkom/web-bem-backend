package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro/service"
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
	router = router.Group("/v2/kemenbiro")
	router.Post("", timeout.NewWithContext(h.CreateKemenbiro(), 5*time.Second))
}
