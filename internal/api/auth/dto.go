package auth

import "github.com/bem-filkom/web-bem-backend/internal/pkg/entity"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string            `json:"access_token"`
	Role        entity.UserRole   `json:"role"`
	User        *entity.User      `json:"user,omitempty"`
	Student     *entity.Student   `json:"student,omitempty"`
	BemMember   *entity.BemMember `json:"bem_member,omitempty"`
}
