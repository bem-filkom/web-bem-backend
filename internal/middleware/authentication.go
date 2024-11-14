package middleware

import (
	"github.com/bem-filkom/web-bem-backend/internal/api/auth"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearer := c.Get("Authorization")
		if bearer == "" {
			return auth.ErrEmptyToken
		}

		tokenSlice := strings.Split(bearer, " ")
		if len(tokenSlice) != 2 {
			return auth.ErrInvalidToken
		}

		token := tokenSlice[1]

		var claims jwt.Claims
		if err := jwt.DecodeAccessToken(token, &claims); err != nil {
			return auth.ErrInvalidToken
		}

		c.Locals("user.id", claims.Subject)
		c.Locals("user.role", claims.Role)

		// BEM Member
		c.Locals("user.kemenbiro_id", claims.KemenbiroID)
		c.Locals("user.kemenbiro_abbreviation", claims.KemenbiroAbbreviation)
		c.Locals("user.is_super_admin", claims.IsSuperAdmin)

		return c.Next()
	}
}
