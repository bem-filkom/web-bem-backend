package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type KabarProker struct {
	ID             string         `json:"id,omitempty" gorm:"primaryKey;type:varchar(255)"` // slug
	ProgramKerjaID uuid.UUID      `json:"program_kerja_id,omitempty" gorm:"not null"`
	ProgramKerja   ProgramKerja   `json:"program_kerja,omitempty" gorm:"foreignKey:ProgramKerjaID;references:ID;constraint:OnDelete:CASCADE;"`
	Title          string         `json:"title,omitempty" gorm:"type:varchar(255);not null"`
	Content        string         `json:"content,omitempty" gorm:"type:text"`
	CreatedAt      time.Time      `json:"created_at,omitempty"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type KabarProkerImage struct {
	ID            uuid.UUID      `json:"id,omitempty" gorm:"primaryKey"`
	KabarProkerID string         `json:"kabar_proker_id,omitempty" gorm:"type:varchar(255);not null"`
	KabarProker   KabarProker    `json:"kabar_proker,omitempty" gorm:"foreignKey:KabarProkerID;references:ID;constraint:OnDelete:CASCADE;"`
	URL           string         `json:"url,omitempty" gorm:"type:text;not null"`
	Description   string         `json:"description,omitempty" gorm:"type:varchar(255)"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
