package repository

import (
	"context"
	"database/sql"
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

func (r *programKerjaRepository) getProgramKerjasByKemenbiroID(ctx context.Context, tx sqlx.ExtContext, kemenbiroID uuid.UUID) ([]*entity.ProgramKerja, error) {
	type Row struct {
		ProkerID              uuid.UUID      `db:"proker_id"`
		ProkerSlug            string         `db:"proker_slug"`
		ProkerName            string         `db:"proker_name"`
		ProkerKemenbiroID     uuid.UUID      `db:"proker_kemenbiro_id"`
		ProkerDescription     sql.NullString `db:"proker_description"`
		KemenbiroAbbreviation string         `db:"kemenbiro_abbreviation"`
		KemenbiroName         string         `db:"kemenbiro_name"`
		PjNim                 sql.NullString `db:"pj_nim"`
		PjProdi               sql.NullString `db:"pj_prodi"`
		PjFullName            sql.NullString `db:"pj_full_name"`
	}

	var rows []*Row
	if err := sqlx.SelectContext(ctx, tx, &rows, getProgramKerjasByKemenbiroIDQuery, kemenbiroID); err != nil {
		return nil, err
	}

	programKerjasMap := make(map[uuid.UUID]*entity.ProgramKerja)

	for _, row := range rows {
		// Check if ProgramKerja already exists in the map
		proker, exists := programKerjasMap[row.ProkerID]
		if !exists {
			proker = &entity.ProgramKerja{
				ID:          row.ProkerID,
				Slug:        row.ProkerSlug,
				Name:        row.ProkerName,
				KemenbiroID: row.ProkerKemenbiroID,
				Kemenbiro: &entity.Kemenbiro{
					Name:         row.KemenbiroName,
					Abbreviation: row.KemenbiroAbbreviation,
				},
				Description: row.ProkerDescription,
			}
			programKerjasMap[row.ProkerID] = proker
		}

		// Add PenanggungJawab only if it's valid
		if row.PjNim.Valid {
			bemMember := &entity.BemMember{
				NIM: row.PjNim.String,
				Student: &entity.Student{
					ProgramStudi: row.PjProdi.String,
					User: &entity.User{
						FullName: row.PjFullName.String,
					},
				},
			}
			proker.PenanggungJawabs = append(proker.PenanggungJawabs, bemMember)
		}
	}

	// Convert map to slice
	prokerTotal := len(programKerjasMap)
	programKerjas := make([]*entity.ProgramKerja, prokerTotal)
	for i := 0; i < prokerTotal; i++ {
		programKerjas[i] = programKerjasMap[rows[i].ProkerID]
	}

	return programKerjas, nil
}

func (r *programKerjaRepository) GetProgramKerjasByKemenbiroID(ctx context.Context, kemenbiroID uuid.UUID) ([]*entity.ProgramKerja, error) {
	return r.getProgramKerjasByKemenbiroID(ctx, r.db, kemenbiroID)
}
