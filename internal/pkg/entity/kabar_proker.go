package entity

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type KabarProker struct {
	ID             string        `json:"id,omitempty"` // slug
	ProgramKerjaID uuid.UUID     `json:"program_kerja_id,omitempty" db:"program_kerja_id"`
	ProgramKerja   *ProgramKerja `json:"program_kerja,omitempty"`
	Title          string        `json:"title,omitempty"`
	Content        string        `json:"content,omitempty"`
	CreatedAt      time.Time     `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at,omitempty" db:"updated_at"`
}

type KabarProkerImage struct {
	ID            uuid.UUID      `json:"id,omitempty"`
	KabarProkerID string         `json:"kabar_proker_id,omitempty" db:"kabar_proker_id"`
	KabarProker   *KabarProker   `json:"kabar_proker,omitempty"`
	URL           string         `json:"url,omitempty"`
	Description   sql.NullString `json:"description,omitempty"`
}

func (k KabarProker) MarshalJSON() ([]byte, error) {
	type Alias KabarProker
	aux := &struct {
		ProgramKerjaID *uuid.UUID `json:"program_kerja_id,omitempty"`
		UpdatedAt      *time.Time `json:"updated_at,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&k),
	}

	if k.ProgramKerjaID != uuid.Nil {
		aux.ProgramKerjaID = &k.ProgramKerjaID
	}

	if !k.UpdatedAt.IsZero() {
		aux.UpdatedAt = &k.UpdatedAt
	}

	return json.Marshal(aux)
}
