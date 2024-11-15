package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

func (r *kemenbiroRepository) createKemenbiro(ctx context.Context, tx sqlx.ExtContext, kemenbiro *entity.Kemenbiro) (*entity.Kemenbiro, error) {
	var kemenbiroObj entity.Kemenbiro

	err := tx.QueryRowxContext(ctx, createKemenbiroQuery, kemenbiro.Abbreviation, kemenbiro.Name, kemenbiro.Description).StructScan(&kemenbiroObj)
	if err != nil {
		return nil, err
	}

	return &kemenbiroObj, nil
}

func (r *kemenbiroRepository) CreateKemenbiro(ctx context.Context, kemenbiro *entity.Kemenbiro) (*entity.Kemenbiro, error) {
	return r.createKemenbiro(ctx, r.db, kemenbiro)
}

func (r *kemenbiroRepository) getKemenbiroByID(ctx context.Context, tx sqlx.ExtContext, id uuid.UUID) (*entity.Kemenbiro, error) {
	var kemenbiroObj entity.Kemenbiro

	err := tx.QueryRowxContext(ctx, getKemenbiroByIDQuery, id).StructScan(&kemenbiroObj)
	if err != nil {
		return nil, err
	}

	return &kemenbiroObj, nil
}

func (r *kemenbiroRepository) GetKemenbiroByID(ctx context.Context, id uuid.UUID) (*entity.Kemenbiro, error) {
	return r.getKemenbiroByID(ctx, r.db, id)
}

func (r *kemenbiroRepository) getAllKemenbiros(ctx context.Context, tx sqlx.ExtContext) ([]*entity.Kemenbiro, error) {
	var kemenbiros []*entity.Kemenbiro

	err := sqlx.SelectContext(ctx, tx, &kemenbiros, getAllKemenbirosQuery)
	if err != nil {
		return nil, err
	}

	return kemenbiros, nil
}

func (r *kemenbiroRepository) GetAllKemenbiros(ctx context.Context) ([]*entity.Kemenbiro, error) {
	return r.getAllKemenbiros(ctx, r.db)
}

func (r *kemenbiroRepository) getKemenbiroByAbbreviation(ctx context.Context, tx sqlx.ExtContext, abbreviation string) (*entity.Kemenbiro, error) {
	var kemenbiroObj entity.Kemenbiro

	err := tx.QueryRowxContext(ctx, getKemenbiroByAbbreviationQuery, abbreviation).StructScan(&kemenbiroObj)
	if err != nil {
		return nil, err
	}

	return &kemenbiroObj, nil
}

func (r *kemenbiroRepository) GetKemenbiroByAbbreviation(ctx context.Context, abbreviation string) (*entity.Kemenbiro, error) {
	return r.getKemenbiroByAbbreviation(ctx, r.db, abbreviation)
}

func (r *kemenbiroRepository) updateKemenbiro(ctx context.Context, tx sqlx.ExtContext, kemenbiro *entity.Kemenbiro) error {
	var queryParts []string
	var args []any
	argIndex := 1

	if kemenbiro.Name != "" {
		queryParts = append(queryParts, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, kemenbiro.Name)
		argIndex++
	}
	if kemenbiro.Abbreviation != "" {
		queryParts = append(queryParts, fmt.Sprintf("abbreviation = $%d", argIndex))
		args = append(args, kemenbiro.Abbreviation)
		argIndex++
	}
	if kemenbiro.Description.Valid {
		queryParts = append(queryParts, fmt.Sprintf("description = $%d", argIndex))
		args = append(args, kemenbiro.Description.String)
		argIndex++
	}

	if len(queryParts) == 0 {
		return errors.New("no fields to update")
	}

	updateQuery := fmt.Sprintf(updateKemenbiroQuery,
		strings.Join(queryParts, ", "),
		argIndex,
	)

	args = append(args, kemenbiro.ID)

	result, err := tx.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *kemenbiroRepository) UpdateKemenbiro(ctx context.Context, kemenbiro *entity.Kemenbiro) error {
	return r.updateKemenbiro(ctx, r.db, kemenbiro)
}

func (r *kemenbiroRepository) deleteKemenbiro(ctx context.Context, tx sqlx.ExtContext, abbreviation string) error {
	result, err := tx.ExecContext(ctx, deleteKemenbiroQuery, abbreviation)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *kemenbiroRepository) DeleteKemenbiro(ctx context.Context, abbreviation string) error {
	return r.deleteKemenbiro(ctx, r.db, abbreviation)
}
