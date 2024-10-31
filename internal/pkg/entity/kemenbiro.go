package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kemenbiro struct {
	ID           uuid.UUID      `json:"id,omitempty" gorm:"primaryKey"`
	Name         string         `json:"name,omitempty" gorm:"type:varchar(255);not null"`
	Abbreviation string         `json:"abbreviation,omitempty" gorm:"type:varchar(255);uniqueIndex;not null"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
