package auth

import (
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"net/http"
)

var (
	ErrInvalidCredentials = response.NewError(http.StatusBadRequest).
		WithRefCode("INVALID_CREDENTIALS").
		WithMessage("Email/NIM dan Password kamu nggak cocok nih. Coba dicek lagi ya!")
)
