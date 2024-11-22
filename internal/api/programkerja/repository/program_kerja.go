package repository

import (
	"context"
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

	for _, pj := range programKerja.PenanggungJawab {
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
