package entity

import (
	"github.com/google/uuid"
)

type Kemenbiro struct {
	ID           uuid.UUID `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Abbreviation string    `json:"abbreviation,omitempty"`
}
