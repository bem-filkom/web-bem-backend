package auth

import (
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"net/http"
)

var (
	ErrInvalidCredentials = response.NewError(http.StatusBadRequest).
				WithRefCode("INVALID_CREDENTIALS").
				WithMessage("Email/NIM dan Password kamu nggak cocok nih. Coba dicek lagi ya!")

	ErrEmptyToken = response.NewError(http.StatusUnauthorized).
			WithRefCode("EMPTY_TOKEN").
			WithMessage("Kamu belum log in.")

	ErrInvalidToken = response.NewError(http.StatusUnauthorized).
			WithRefCode("INVALID_TOKEN").
			WithMessage("Silakan log in ulang.")
)
