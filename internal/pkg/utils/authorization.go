package utils

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/google/uuid"
)

// RequireKemenbiroID Dependencies: [middleware.Authenticate]
func RequireKemenbiroID(ctx context.Context, kemenbiroID uuid.UUID) error {
	if isSuperAdmin := ctx.Value("user.is_super_admin").(bool); isSuperAdmin {
		return nil
	}

	userKemenbiroID, ok := ctx.Value("user.kemenbiro_id").(uuid.UUID)
	if !ok {
		return response.ErrForbiddenRole
	}

	if userKemenbiroID != kemenbiroID {
		return response.ErrForbiddenKemenbiro
	}

	return nil
}
