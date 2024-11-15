package entity

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
)

type Kemenbiro struct {
	ID           uuid.UUID      `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Abbreviation string         `json:"abbreviation,omitempty"`
	Description  sql.NullString `json:"description,omitempty"`
}

func (k Kemenbiro) MarshalJSON() ([]byte, error) {
	type Alias Kemenbiro
	aux := &struct {
		ID          *uuid.UUID `json:"id,omitempty"`
		Description *string    `json:"description,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&k),
	}

	if k.ID != uuid.Nil {
		aux.ID = &k.ID
	}

	if k.Description.Valid {
		aux.Description = &k.Description.String
	}

	return json.Marshal(aux)
}
