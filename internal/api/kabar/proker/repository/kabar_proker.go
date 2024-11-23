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

func (r *kabarProkerRepository) getKabarProkerByQuery(
	ctx context.Context,
	tx sqlx.ExtContext,
	conditions *proker.GetKabarProkerByQueryRequest,
) ([]*entity.KabarProker, error) {
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

	if !conditions.Before.IsZero() {
		queryConditions = append(queryConditions, fmt.Sprintf("kp.created_at < $%d", paramIndex))
		queryParams = append(queryParams, conditions.Before)
		paramIndex++
	}

	// Fallback if no conditions are provided
	whereClause := "1=1"
	if len(queryConditions) > 0 {
		whereClause = strings.Join(queryConditions, " AND ")
	}

	query := fmt.Sprintf(getKabarProkerQuery, whereClause, paramIndex)

	// Add limit parameter
	queryParams = append(queryParams, conditions.Limit)

	// Execute the query
	var rows []*proker.GetKabarProkerQueryRow
	if err := sqlx.SelectContext(ctx, tx, &rows, query, queryParams...); err != nil {
		return nil, err
	}

	return getKabarProkersFromRow(rows), nil
}

func (r *kabarProkerRepository) GetKabarProkerByQuery(ctx context.Context, conditions *proker.GetKabarProkerByQueryRequest) ([]*entity.KabarProker, error) {
	return r.getKabarProkerByQuery(ctx, r.db, conditions)
}
