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
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	SaveStudent(ctx context.Context, student *entity.Student) error
	GetStudentByNIM(ctx context.Context, nim string) (*entity.Student, error)
	CreateBemMember(ctx context.Context, bemMember *entity.BemMember) error
	GetBemMemberByNIM(ctx context.Context, nim string) (*entity.BemMember, error)
	UpdateBemMember(ctx context.Context, updates *entity.BemMember) error
	DeleteBemMember(ctx context.Context, nim string) error
	GetRole(ctx context.Context, nim string) (entity.UserRole, error)
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &userRepository{db: db}
}
