package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja/repository"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type kabarProkerRepository struct {
	db  *sqlx.DB
	pkr repository.IProgramKerjaRepository
}

type IKabarProkerRepository interface {
	CreateKabarProker(ctx context.Context, kabarProker *entity.KabarProker) error
	GetKabarProkerByQuery(ctx context.Context, conditions *proker.GetKabarProkerByQueryRequest) ([]*entity.KabarProker, error)
}

func NewKabarProkerRepository(db *sqlx.DB) IKabarProkerRepository {
	return &kabarProkerRepository{db: db}
}
