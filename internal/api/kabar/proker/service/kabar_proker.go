package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/api/kabar/proker"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
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

func (s *kabarProkerService) GetKabarProkerByQuery(ctx context.Context, req *proker.GetKabarProkerByQueryRequest) ([]*entity.KabarProker, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	kabarProkers, err := s.r.GetKabarProkerByQuery(ctx, req)
	if err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   err,
			"request": req,
		}).Errorln("[KabarProkerService][GetKabarProkerByQuery] fail to get kabar proker by query")
		return nil, response.ErrInternalServerError
	}

	return kabarProkers, nil
}
