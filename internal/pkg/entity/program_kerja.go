package entity

import (
	"database/sql"
	"github.com/google/uuid"
)

type ProgramKerja struct {
	ID              uuid.UUID      `json:"id,omitempty"`
	Slug            string         `json:"slug,omitempty"`
	Name            string         `json:"name,omitempty"`
	KemenbiroID     uuid.UUID      `json:"kemenbiro_id,omitempty" db:"kemenbiro_id"`
	Kemenbiro       *Kemenbiro     `json:"kemenbiro,omitempty"`
	Description     sql.NullString `json:"description,omitempty"`
	PenanggungJawab []*BemMember   `json:"penanggung_jawab,omitempty" db:"penanggung_jawab"`
}
