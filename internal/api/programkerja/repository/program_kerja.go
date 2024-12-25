package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/sqlutil"
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

func (r *programKerjaRepository) createPenanggungJawabs(ctx context.Context, tx sqlx.PreparerContext, pjs []*entity.BemMember, prokerID uuid.UUID) error {
	stmt, err := tx.PrepareContext(ctx, createPenanggungJawabQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, pj := range pjs {
		if _, err := stmt.ExecContext(ctx, pj.NIM, prokerID); err != nil {
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

	if err := r.createPenanggungJawabs(ctx, tx, programKerja.PenanggungJawabs, programKerja.ID); err != nil {
		return uuid.Nil, err
	}

	return prokerID, tx.Commit()
}

func (r *programKerjaRepository) getProgramKerjaByID(ctx context.Context, tx sqlx.ExtContext, id uuid.UUID) (*entity.ProgramKerja, error) {
	var rows []*programkerja.GetProgramKerjaQueryRow
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
	var rows []*programkerja.GetProgramKerjaQueryRow
	if err := sqlx.SelectContext(ctx, tx, &rows, getProgramKerjasByKemenbiroIDQuery, kemenbiroID); err != nil {
		return nil, err
	}

	return getProgramKerjasFromRow(rows), nil
}

func (r *programKerjaRepository) GetProgramKerjasByKemenbiroID(ctx context.Context, kemenbiroID uuid.UUID) ([]*entity.ProgramKerja, error) {
	return r.getProgramKerjasByKemenbiroID(ctx, r.db, kemenbiroID)
}

func (r *programKerjaRepository) getKemenbiroIDByProgramKerjaID(ctx context.Context, tx sqlx.ExtContext, programKerjaID uuid.UUID) (uuid.UUID, error) {
	var kemenbiroID uuid.UUID
	if err := tx.QueryRowxContext(ctx, getKemenbiroIDByProgramKerjaIDQuery, programKerjaID).Scan(&kemenbiroID); err != nil {
		return uuid.Nil, err
	}

	return kemenbiroID, nil
}

func (r *programKerjaRepository) GetKemenbiroIDByProgramKerjaID(ctx context.Context, prokerID uuid.UUID) (uuid.UUID, error) {
	return r.getKemenbiroIDByProgramKerjaID(ctx, r.db, prokerID)
}

func (r *programKerjaRepository) updatePenanggungJawabs(ctx context.Context, tx sqlx.PreparerContext, updates *programkerja.UpdateProgramKerjaRequest) error {
	if err := r.deletePenanggungJawabs(ctx, tx, updates.ID); err != nil {
		return err
	}

	if err := r.createPenanggungJawabs(ctx, tx, updates.PenanggungJawabs, updates.ID); err != nil {
		return err
	}

	return nil
}

func (r *programKerjaRepository) updateProgramKerja(ctx context.Context, tx *sqlx.Tx, updates *programkerja.UpdateProgramKerjaRequest) error {
	queryPart, args, argIndex, err := sqlutil.GenerateUpdateQueryPart(updates)
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   err,
			"updates": updates,
		}).Errorln("[ProgramKerjaRepository][updateProgramKerja] fail to generate update query parts")
	}

	if len(queryPart) == 0 && updates.PenanggungJawabs == nil {
		return errors.New("no fields to update")
	}

	if updates.PenanggungJawabs != nil {
		if err := r.updatePenanggungJawabs(ctx, tx, updates); err != nil {
			return err
		}
	}

	updateQuery := fmt.Sprintf(updateProgramKerjaQuery, queryPart, argIndex)
	args = append(args, updates.ID)

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

func (r *programKerjaRepository) UpdateProgramKerja(ctx context.Context, updates *programkerja.UpdateProgramKerjaRequest) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.updateProgramKerja(ctx, tx, updates); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *programKerjaRepository) deletePenanggungJawabs(ctx context.Context, tx sqlx.PreparerContext, id uuid.UUID) error {
	stmt, err := tx.PrepareContext(ctx, deletePenanggungJawabsQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
