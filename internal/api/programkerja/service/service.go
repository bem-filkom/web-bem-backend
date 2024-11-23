package service

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja"
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja/repository"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/google/uuid"
)

type programKerjaService struct {
	r repository.IProgramKerjaRepository
}

type IProgramKerjaService interface {
	CreateProgramKerja(ctx context.Context, req *programkerja.CreateProgramKerjaRequest) (*entity.ProgramKerja, error)
	GetProgramKerjaByID(ctx context.Context, req *programkerja.GetProgramKerjaByIDRequest) (*entity.ProgramKerja, error)
	GetProgramKerjasByKemenbiroID(ctx context.Context, req *programkerja.GetProgramKerjasByKemenbiroIDRequest) ([]*entity.ProgramKerja, error)
	GetKemenbiroIDByProgramKerjaID(ctx context.Context, prokerID uuid.UUID) (uuid.UUID, error)
}

func NewProgramKerjaService(r repository.IProgramKerjaRepository) IProgramKerjaService {
	return &programKerjaService{r: r}
}
