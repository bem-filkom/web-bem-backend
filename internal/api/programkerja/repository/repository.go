package repository

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type programKerjaRepository struct {
	db *sqlx.DB
}

type IProgramKerjaRepository interface {
	CreateProgramKerja(ctx context.Context, programKerja *entity.ProgramKerja) (uuid.UUID, error)
	GetProgramKerjaByID(ctx context.Context, id uuid.UUID) (*entity.ProgramKerja, error)
	GetProgramKerjasByKemenbiroID(ctx context.Context, kemenbiroID uuid.UUID) ([]*entity.ProgramKerja, error)
}

func NewProgramKerjaRepository(db *sqlx.DB) IProgramKerjaRepository {
	return &programKerjaRepository{db: db}
}
