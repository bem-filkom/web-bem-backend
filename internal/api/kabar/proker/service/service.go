package service

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker/repository"
	prokerRepo "github.com/bem-filkom/web-bem-backend/internal/api/programkerja/repository"
)

type kabarProkerService struct {
	r   repository.IKabarProkerRepository
	pkr prokerRepo.IProgramKerjaRepository
}

type IKabarProkerService interface {
	CreateKabarProker(ctx context.Context, req *proker.CreateKabarProkerRequest) error
}

func NewKabarProkerService(r repository.IKabarProkerRepository, pkr prokerRepo.IProgramKerjaRepository) IKabarProkerService {
	return &kabarProkerService{r: r, pkr: pkr}
}
