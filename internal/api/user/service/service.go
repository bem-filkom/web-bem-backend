package service

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/api/user/repository"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
)

type userService struct {
	r repository.IUserRepository
}

type IUserService interface {
	SaveUser(ctx context.Context, req *user.SaveUserRequest) error
	SaveStudent(ctx context.Context, req *user.SaveStudentRequest) error
	CreateBemMember(ctx context.Context, req *user.CreateBemMemberRequest) error
	GetBemMemberByNIM(ctx context.Context, req *user.GetUserRequest) (*entity.BemMember, error)
	GetRole(ctx context.Context, req *user.GetUserRequest) (entity.UserRole, error)
}

func NewUserService(r repository.IUserRepository) IUserService {
	return &userService{
		r: r,
	}
}
