package programkerja

import (
	"database/sql"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
)

type CreateProgramKerjaRequest struct {
	Slug             string              `json:"slug" validate:"required,max=255,slug"`
	Name             string              `json:"name" validate:"required,max=255"`
	KemenbiroID      uuid.UUID           `json:"kemenbiro_id" validate:"required,uuid"`
	Description      string              `json:"description" validate:"omitempty,max=2000"`
	PenanggungJawabs []*entity.BemMember `json:"penanggung_jawabs" validate:"required"`
}

type GetProgramKerjasByKemenbiroIDRequest struct {
	KemenbiroID uuid.UUID `query:"kemenbiro_id" validate:"required,uuid"`
}

type GetProgramKerjaByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required,uuid"`
}

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
