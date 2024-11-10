package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type kemenbiroRepository struct {
	db *sqlx.DB
}

type IKemenbiroRepository interface {
	CreateKemenbiro(ctx context.Context, kemenbiro *entity.Kemenbiro) (*entity.Kemenbiro, error)
}

func NewKemenbiroRepository(db *sqlx.DB) IKemenbiroRepository {
	return &kemenbiroRepository{db: db}
}
