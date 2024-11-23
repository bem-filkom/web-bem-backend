package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
)

func (r *kabarProkerRepository) createKabarProker(ctx context.Context, tx sqlx.ExtContext, kabarProker *entity.KabarProker) error {
	_, err := tx.ExecContext(ctx, createKabarProkerQuery, kabarProker.ID, kabarProker.ProgramKerjaID, kabarProker.Title, kabarProker.Content)
	return err
}

func (r *kabarProkerRepository) CreateKabarProker(ctx context.Context, kabarProker *entity.KabarProker) error {
	return r.createKabarProker(ctx, r.db, kabarProker)
}
