package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type kabarProkerRepository struct {
	db *sqlx.DB
}

type IKabarProkerRepository interface {
	CreateKabarProker(ctx context.Context, kabarProker *entity.KabarProker) error
}

func NewKabarProkerRepository(db *sqlx.DB) IKabarProkerRepository {
	return &kabarProkerRepository{db: db}
}
