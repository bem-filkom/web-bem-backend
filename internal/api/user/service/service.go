package service

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/api/user/repository"
)

type userService struct {
	r repository.IUserRepository
}

type IUserService interface {
	SaveUser(ctx context.Context, req *user.SaveUserRequest) error
	SaveStudent(ctx context.Context, req *user.SaveStudentRequest) error
}

func NewUserService(r repository.IUserRepository) IUserService {
	return &userService{
		r: r,
	}
}
