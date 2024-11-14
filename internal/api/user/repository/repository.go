package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type IUserRepository interface {
	SaveUser(ctx context.Context, user *entity.User) error
	SaveStudent(ctx context.Context, student *entity.Student) error
	CreateBemMember(ctx context.Context, bemMember *entity.BemMember) error
	GetBemMemberByNIM(ctx context.Context, nim string) (*entity.BemMember, error)
	GetRole(ctx context.Context, nim string) (entity.UserRole, error)
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &userRepository{db: db}
}
