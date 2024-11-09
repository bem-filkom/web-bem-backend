package service

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/auth"
	"github.com/bem-filkom/web-bem-backend/internal/api/user/service"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/ubauth"
)

type authService struct {
	us     service.IUserService
	ubAuth ubauth.IUBAuth
}

type IAuthService interface {
	LoginAuthUb(ctx context.Context, req *auth.LoginRequest) (string, error)
}

func NewAuthService(us service.IUserService, ubAuth ubauth.IUBAuth) IAuthService {
	return &authService{
		us:     us,
		ubAuth: ubAuth,
	}
}
