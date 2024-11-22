package repository

import (
	"context"
	"database/sql"
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func createProgramKerja(ctx context.Context, tx sqlx.ExtContext, programKerja *entity.ProgramKerja) (uuid.UUID, error) {
	var id uuid.UUID

	if err := tx.QueryRowxContext(ctx, createProgramKerjaQuery,
		programKerja.Slug,
		programKerja.Name,
		programKerja.KemenbiroID,
		programKerja.Description,
	).Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *programKerjaRepository) createPenanggungJawab(ctx context.Context, tx sqlx.PreparerContext, programKerja *entity.ProgramKerja) error {
	stmt, err := tx.PrepareContext(ctx, createPenanggungJawabQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, pj := range programKerja.PenanggungJawabs {
		if _, err := stmt.ExecContext(ctx, pj.NIM, programKerja.ID); err != nil {
			return err
		}
	}

	return nil
}

func (r *programKerjaRepository) CreateProgramKerja(ctx context.Context, programKerja *entity.ProgramKerja) (uuid.UUID, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()

	prokerID, err := createProgramKerja(ctx, tx, programKerja)
	if err != nil {
		return uuid.Nil, err
	}

	programKerja.ID = prokerID

	if err := r.createPenanggungJawab(ctx, tx, programKerja); err != nil {
		return uuid.Nil, err
	}

	return prokerID, tx.Commit()
}

func (r *programKerjaRepository) getProgramKerjaByID(ctx context.Context, tx sqlx.ExtContext, id uuid.UUID) (*entity.ProgramKerja, error) {
	var rows []*programkerja.Row
	if err := sqlx.SelectContext(ctx, tx, &rows, getProgramKerjaByIDQuery, id); err != nil {
		return nil, err
	}

	programKerjas := getProgramKerjasFromRow(rows)
	if len(programKerjas) == 0 {
		return nil, sql.ErrNoRows
	}

	return programKerjas[0], nil
}

func (r *programKerjaRepository) GetProgramKerjaByID(ctx context.Context, id uuid.UUID) (*entity.ProgramKerja, error) {
	return r.getProgramKerjaByID(ctx, r.db, id)
}

func (r *programKerjaRepository) getProgramKerjasByKemenbiroID(ctx context.Context, tx sqlx.ExtContext, kemenbiroID uuid.UUID) ([]*entity.ProgramKerja, error) {
	var rows []*programkerja.Row
	if err := sqlx.SelectContext(ctx, tx, &rows, getProgramKerjasByKemenbiroIDQuery, kemenbiroID); err != nil {
		return nil, err
	}

	return getProgramKerjasFromRow(rows), nil
}

func (r *programKerjaRepository) GetProgramKerjasByKemenbiroID(ctx context.Context, kemenbiroID uuid.UUID) ([]*entity.ProgramKerja, error) {
	return r.getProgramKerjasByKemenbiroID(ctx, r.db, kemenbiroID)
}
