package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/user/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	s service.IUserService
}

func NewUserHandler(s service.IUserService) *UserHandler {
	return &UserHandler{s: s}
}

func (h *UserHandler) Start(router fiber.Router) {
	router = router.Group("/v2/users")

	router.Post("/bem-member", h.CreateBemMember())
}
