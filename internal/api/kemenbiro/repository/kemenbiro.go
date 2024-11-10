package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
)

func (r *kemenbiroRepository) createKemenbiro(ctx context.Context, tx sqlx.ExtContext, kemenbiro *entity.Kemenbiro) (*entity.Kemenbiro, error) {
	var kemenbiroObj entity.Kemenbiro

	err := tx.QueryRowxContext(ctx, createKemenbiroQuery, kemenbiro.Abbreviation, kemenbiro.Name).StructScan(&kemenbiroObj)
	if err != nil {
		return nil, err
	}

	return &kemenbiroObj, nil
}

func (r *kemenbiroRepository) CreateKemenbiro(ctx context.Context, kemenbiro *entity.Kemenbiro) (*entity.Kemenbiro, error) {
	return r.createKemenbiro(ctx, r.db, kemenbiro)
}
