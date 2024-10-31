package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `json:"id,omitempty" gorm:"primaryKey;type:varchar(15)"`
	Email     string         `json:"email,omitempty" gorm:"type:varchar(320);uniqueIndex;not null"`
	FullName  string         `json:"full_name,omitempty" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Student struct {
	NIM          string         `json:"nim,omitempty" gorm:"primaryKey;type:varchar(15)"`
	User         User           `json:"user,omitempty" gorm:"foreignKey:NIM;references:ID;constraint:OnDelete:CASCADE;"`
	ProgramStudi string         `json:"program_studi,omitempty" gorm:"type:varchar(255);not null"`
	Fakultas     string         `json:"fakultas,omitempty" gorm:"type:varchar(255);not null"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type BemMember struct {
	NIM         string         `json:"nim,omitempty" gorm:"primaryKey;type:varchar(15)"`
	Student     Student        `json:"student,omitempty" gorm:"foreignKey:NIM;references:NIM;constraint:OnDelete:CASCADE;"`
	KemenbiroID uuid.UUID      `json:"kemenbiro_id,omitempty" gorm:"not null"`
	Kemenbiro   Kemenbiro      `json:"kemenbiro,omitempty" gorm:"foreignKey:KemenbiroID;references:ID;constraint:OnDelete:SET NULL;"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
