package auth

import "github.com/bem-filkom/web-bem-backend/internal/pkg/response"

var (
	ErrInvalidCredentials = response.NewError(400, "INVALID_CREDENTIALS", "Email/NIM dan Password kamu nggak cocok nih. Coba dicek lagi ya!")
)
