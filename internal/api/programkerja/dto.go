package programkerja

import (
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
