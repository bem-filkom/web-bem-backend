package middleware

import (
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// RequireRole Dependencies: [Authenticate]
func RequireRole(role entity.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("user.role").(entity.UserRole)
		if userRole != role {
			return response.ErrForbiddenRole
		}

		return c.Next()
	}
}

// RequireSuperAdmin Dependencies: [Authenticate]
func RequireSuperAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isSuperAdmin := c.Locals("user.is_super_admin").(bool)
		if !isSuperAdmin {
			return response.ErrForbiddenSuperAdmin
		}

		return c.Next()
	}
}
