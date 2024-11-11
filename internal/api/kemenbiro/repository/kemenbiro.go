package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *kemenbiroRepository) updateKemenbiro(ctx context.Context, tx sqlx.ExtContext, abbreviationAsID string, kemenbiro *entity.Kemenbiro) error {
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

	args = append(args, abbreviationAsID)

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

func (r *kemenbiroRepository) UpdateKemenbiro(ctx context.Context, abbreviationAsID string, kemenbiro *entity.Kemenbiro) error {
	return r.updateKemenbiro(ctx, r.db, abbreviationAsID, kemenbiro)
}
