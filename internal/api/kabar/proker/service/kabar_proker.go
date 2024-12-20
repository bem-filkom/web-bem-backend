package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/pagination"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/utils"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *kabarProkerService) CreateKabarProker(ctx context.Context, req *proker.CreateKabarProkerRequest) error {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return response.ErrValidation.WithDetail(err)
	}

	kemenbiroID, err := s.pkr.GetKemenbiroIDByProgramKerjaID(ctx, req.ProgramKerjaID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrNotFound.WithMessage("Program kerja tidak ditemukan")
		}

		log.GetLogger().WithFields(map[string]any{
			"error":   err,
			"request": req,
		}).Errorln("[KabarProkerService][CreateKabarProker] fail to get kemenbiro ID by program kerja ID")
	}

	if err := utils.RequireKemenbiroID(ctx, kemenbiroID); err != nil {
		return err
	}

	kabarProker := &entity.KabarProker{
		ID:             req.ID,
		ProgramKerjaID: req.ProgramKerjaID,
		Title:          req.Title,
		Content:        req.Content,
	}

	err = s.r.CreateKabarProker(ctx, kabarProker)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			switch pgErr.ConstraintName {
			case "kabar_prokers_program_kerja_id_fkey":
				return response.ErrNotFound.WithMessage("Program kerja tidak ditemukan")
			case "kabar_prokers_pkey":
				return response.ErrConflictSlug
			}
		}

		log.GetLogger().WithFields(map[string]any{
			"error":   err,
			"request": req,
		}).Errorln("[KabarProkerService][CreateKabarProker] fail to create kabar proker")
		return response.ErrInternalServerError
	}

	return nil
}

func (s *kabarProkerService) GetKabarProkerByID(ctx context.Context, req proker.GetKabarProkerByIDRequest) (*entity.KabarProker, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	kabarProker, err := s.r.GetKabarProkerByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.ErrNotFound.WithMessage("Kabar proker tidak ditemukan")
		}

		log.GetLogger().WithFields(map[string]any{
			"error":   err,
			"request": req,
		}).Errorln("[KabarProkerService][GetKabarProkerByID] fail to get kabar proker by ID")
		return nil, response.ErrInternalServerError
	}

	return kabarProker, nil
}

func (s *kabarProkerService) GetKabarProkerByQuery(ctx context.Context, req *proker.GetKabarProkerByQueryRequest) ([]*entity.KabarProker, *pagination.Response, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, nil, response.ErrValidation.WithDetail(err)
	}

	offset := (req.Page - 1) * req.Limit

	kabarProkers, count, err := s.r.GetKabarProkerByQuery(ctx, req, offset)
	if err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   err,
			"request": req,
		}).Errorln("[KabarProkerService][GetKabarProkerByQuery] fail to get kabar proker by query")
		return nil, nil, response.ErrInternalServerError
	}

	paginationRes := pagination.NewPagination(count, req.Page, req.Limit)

	return kabarProkers, paginationRes, nil
}
