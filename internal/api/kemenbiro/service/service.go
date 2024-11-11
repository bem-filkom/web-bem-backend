package service

import (
	"context"
	"github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro"
	"github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro/repository"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
)

type kemenbiroService struct {
	r repository.IKemenbiroRepository
}

type IKemenbiroService interface {
	CreateKemenbiro(ctx context.Context, req *kemenbiro.CreateKemenbiroRequest) (*entity.Kemenbiro, error)
	GetKemenbiroByAbbreviation(ctx context.Context, req *kemenbiro.GetKemenbiroByAbbreviationRequest) (*entity.Kemenbiro, error)
	UpdateKemenbiro(ctx context.Context, req *kemenbiro.UpdateKemenbiroRequest) error
	DeleteKemenbiro(ctx context.Context, req *kemenbiro.DeleteKemenbiroRequest) error
}

func NewKemenbiroService(r repository.IKemenbiroRepository) IKemenbiroService {
	return &kemenbiroService{r: r}
}
