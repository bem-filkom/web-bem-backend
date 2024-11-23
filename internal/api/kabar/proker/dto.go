package proker

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type CreateKabarProkerRequest struct {
	ID             string    `json:"id" validate:"required,slug,max=255"`
	ProgramKerjaID uuid.UUID `json:"program_kerja_id" validate:"required,uuid"`
	Title          string    `json:"title" validate:"required,max=255"`
	Content        string    `json:"content" validate:"required"`
}

type GetKabarProkerByIDRequest struct {
	ID string `param:"id"`
}

type GetKabarProkerByQueryRequest struct {
	ProgramKerjaID uuid.UUID `query:"program_kerja_id" validate:"omitempty,uuid"`
	KemenbiroID    uuid.UUID `query:"kemenbiro_id" validate:"omitempty,uuid"`
	Before         time.Time `query:"before" validate:"omitempty"`
	Limit          uint      `query:"limit" validate:"required,max=30"`
}

type GetKabarProkerQueryRow struct {
	KabarProkerID         string         `db:"kabar_proker_id"`
	KabarProkerTitle      string         `db:"kabar_proker_title"`
	KabarProkerContent    string         `db:"kabar_proker_content"`
	KabarProkerCreatedAt  time.Time      `db:"kabar_proker_created_at"`
	KabarProkerUpdatedAt  time.Time      `db:"kabar_proker_updated_at"`
	ProkerID              uuid.UUID      `db:"proker_id"`
	ProkerSlug            string         `db:"proker_slug"`
	ProkerName            string         `db:"proker_name"`
	ProkerKemenbiroID     uuid.UUID      `db:"proker_kemenbiro_id"`
	KemenbiroAbbreviation string         `db:"kemenbiro_abbreviation"`
	KemenbiroName         string         `db:"kemenbiro_name"`
	PjNim                 sql.NullString `db:"pj_nim"`
	PjProdi               sql.NullString `db:"pj_prodi"`
	PjFullName            sql.NullString `db:"pj_full_name"`
}
