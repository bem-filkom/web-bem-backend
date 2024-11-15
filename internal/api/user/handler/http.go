package handler

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/user/service"
	"github.com/bem-filkom/web-bem-backend/internal/middleware"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"time"
)

type UserHandler struct {
	s service.IUserService
}

func NewUserHandler(s service.IUserService) *UserHandler {
	return &UserHandler{s: s}
}

func (h *UserHandler) Start(router fiber.Router) {
	router = router.Group("/v2/users")

	router.Post("/bem-member",
		middleware.Authenticate(),
		middleware.RequireRole(entity.RoleBemMember),
		timeout.NewWithContext(h.CreateBemMember(), 5*time.Second),
	)

	router.Patch("/bem-member/:nim",
		middleware.Authenticate(),
		middleware.RequireRole(entity.RoleBemMember),
		timeout.NewWithContext(h.UpdateBemMember(), 5*time.Second))

	router.Delete("/bem-member/:nim",
		middleware.Authenticate(),
		middleware.RequireRole(entity.RoleBemMember),
		timeout.NewWithContext(h.DeleteBemMember(), 5*time.Second))
}
