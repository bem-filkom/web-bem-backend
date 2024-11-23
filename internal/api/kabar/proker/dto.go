package proker

import "github.com/google/uuid"

type CreateKabarProkerRequest struct {
	ID             string    `json:"id" validate:"required,slug,max=255"`
	ProgramKerjaID uuid.UUID `json:"program_kerja_id" validate:"required,uuid"`
	Title          string    `json:"title" validate:"required,max=255"`
	Content        string    `json:"content" validate:"required"`
}
