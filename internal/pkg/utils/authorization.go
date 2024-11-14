package utils

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/google/uuid"
)

func RequireKemenbiroID(ctx context.Context, kemenbiroID uuid.UUID) error {
	userKemenbiroID, ok := ctx.Value("user.kemenbiro_id").(uuid.UUID)
	if !ok {
		return response.ErrForbiddenRole
	}

	if userKemenbiroID != kemenbiroID {
		return response.ErrForbiddenKemenbiro
	}

	return nil
}
