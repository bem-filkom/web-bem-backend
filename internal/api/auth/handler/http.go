package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/auth/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"time"
)

type AuthHandler struct {
	s service.IAuthService
}

func NewAuthHandler(s service.IAuthService) *AuthHandler {
	return &AuthHandler{s: s}
}

func (h *AuthHandler) Start(router fiber.Router) {
	router = router.Group("/v2/auth")
	router.Post("/ub/login", timeout.NewWithContext(h.LoginUB(), 10*time.Second))
}
