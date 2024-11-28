package repository

import (
	"context"
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strings"
)

func (r *kabarProkerRepository) createKabarProker(ctx context.Context, tx sqlx.ExtContext, kabarProker *entity.KabarProker) error {
	_, err := tx.ExecContext(ctx, createKabarProkerQuery, kabarProker.ID, kabarProker.ProgramKerjaID, kabarProker.Title, kabarProker.Content)
	return err
}

func (r *kabarProkerRepository) CreateKabarProker(ctx context.Context, kabarProker *entity.KabarProker) error {
	return r.createKabarProker(ctx, r.db, kabarProker)
}

func (r *kabarProkerRepository) getKabarProkerByID(ctx context.Context, tx sqlx.ExtContext, id string) (*entity.KabarProker, error) {
	var row proker.GetKabarProkerQueryRow
	if err := tx.QueryRowxContext(ctx, getKabarProkerByIDQuery, id).StructScan(&row); err != nil {
		return nil, err
	}

	return getKabarProkerFromRow(&row), nil
}

func (r *kabarProkerRepository) GetKabarProkerByID(ctx context.Context, id string) (*entity.KabarProker, error) {
	return r.getKabarProkerByID(ctx, r.db, id)
}

func (r *kabarProkerRepository) getKabarProkerByQuery(
	ctx context.Context,
	tx sqlx.ExtContext,
	conditions *proker.GetKabarProkerByQueryRequest,
	offset uint,
) ([]*entity.KabarProker, int64, error) {
	var queryConditions []string
	var queryParams []interface{}
	paramIndex := 1 // SQLX uses 1-based indexing for parameters in PostgreSQL.

	// Initialize query conditions
	if conditions.ProgramKerjaID != (uuid.UUID{}) {
		queryConditions = append(queryConditions, fmt.Sprintf("kp.program_kerja_id = $%d", paramIndex))
		queryParams = append(queryParams, conditions.ProgramKerjaID)
		paramIndex++
	}

	if conditions.KemenbiroID != (uuid.UUID{}) {
		queryConditions = append(queryConditions, fmt.Sprintf("pk.kemenbiro_id = $%d", paramIndex))
		queryParams = append(queryParams, conditions.KemenbiroID)
		paramIndex++
	}

	// Fallback if no conditions are provided
	whereClause := "1=1"
	if len(queryConditions) > 0 {
		whereClause = strings.Join(queryConditions, " AND ")
	}

	countQuery := fmt.Sprintf(getKabarProkerCountQuery, whereClause)
	// Execute the count query
	var count int64
	if err := sqlx.GetContext(ctx, tx, &count, countQuery, queryParams...); err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf(getKabarProkerQuery, whereClause, paramIndex, paramIndex+1)

	// Add limit parameter
	queryParams = append(queryParams, conditions.Limit, offset)

	// Execute the query
	var rows []*proker.GetKabarProkerQueryRow
	if err := sqlx.SelectContext(ctx, tx, &rows, query, queryParams...); err != nil {
		return nil, 0, err
	}

	return getKabarProkersFromRows(rows), count, nil
}

func (r *kabarProkerRepository) GetKabarProkerByQuery(ctx context.Context, conditions *proker.GetKabarProkerByQueryRequest, offset uint) ([]*entity.KabarProker, int64, error) {
	return r.getKabarProkerByQuery(ctx, r.db, conditions, offset)
}
