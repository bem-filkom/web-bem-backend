package service

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker/repository"
	prokerRepo "github.com/bem-filkom/web-bem-backend/internal/api/programkerja/repository"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/pagination"
)

type kabarProkerService struct {
	r   repository.IKabarProkerRepository
	pkr prokerRepo.IProgramKerjaRepository
}

type IKabarProkerService interface {
	CreateKabarProker(ctx context.Context, req *proker.CreateKabarProkerRequest) error
	GetKabarProkerByQuery(ctx context.Context, req *proker.GetKabarProkerByQueryRequest) ([]*entity.KabarProker, *pagination.Response, error)
}

func NewKabarProkerService(r repository.IKabarProkerRepository, pkr prokerRepo.IProgramKerjaRepository) IKabarProkerService {
	return &kabarProkerService{r: r, pkr: pkr}
}
