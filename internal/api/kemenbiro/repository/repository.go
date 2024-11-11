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
	GetKemenbiroByAbbreviation(ctx context.Context, abbreviation string) (*entity.Kemenbiro, error)
	UpdateKemenbiro(ctx context.Context, abbreviationAsID string, kemenbiro *entity.Kemenbiro) error
	DeleteKemenbiro(ctx context.Context, abbreviation string) error
}

func NewKemenbiroRepository(db *sqlx.DB) IKemenbiroRepository {
	return &kemenbiroRepository{db: db}
}
