package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramKerja struct {
	ID          uuid.UUID      `json:"id,omitempty" gorm:"primaryKey"`
	Slug        string         `json:"slug,omitempty" gorm:"type:varchar(255);uniqueIndex;not null"`
	Name        string         `json:"name,omitempty" gorm:"type:varchar(255);not null"`
	KemenbiroID uuid.UUID      `json:"kemenbiro_id,omitempty" gorm:"not null"`
	Kemenbiro   Kemenbiro      `json:"kemenbiro,omitempty" gorm:"foreignKey:KemenbiroID;references:ID;constraint:OnDelete:CASCADE;"`
	Description string         `json:"description,omitempty" gorm:"type:text"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
