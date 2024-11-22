package entity

import (
	"database/sql"
	"encoding/json"
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

func (k ProgramKerja) MarshalJSON() ([]byte, error) {
	type Alias ProgramKerja
	aux := &struct {
		ID          *uuid.UUID `json:"id,omitempty"`
		KemenbiroID *uuid.UUID `json:"kemenbiro_id,omitempty"`
		Description *string    `json:"description,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&k),
	}

	if k.ID != uuid.Nil {
		aux.ID = &k.ID
	}

	if k.KemenbiroID != uuid.Nil {
		aux.KemenbiroID = &k.KemenbiroID
	}

	if k.Description.Valid {
		aux.Description = &k.Description.String
	}

	return json.Marshal(aux)
}
