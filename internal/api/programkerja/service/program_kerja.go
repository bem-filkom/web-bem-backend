package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/api/programkerja"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/utils"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *programKerjaService) CreateProgramKerja(ctx context.Context, req *programkerja.CreateProgramKerjaRequest) (*entity.ProgramKerja, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	if err := utils.RequireKemenbiroID(ctx, req.KemenbiroID); err != nil {
		return nil, err
	}

	programKerjaObj := &entity.ProgramKerja{
		Slug:             req.Slug,
		Name:             req.Name,
		KemenbiroID:      req.KemenbiroID,
		Description:      sql.NullString{String: req.Description, Valid: req.Description != ""},
		PenanggungJawabs: req.PenanggungJawabs,
	}

	id, err := s.r.CreateProgramKerja(ctx, programKerjaObj)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			switch pgErr.ConstraintName {
			case "program_kerjas_kemenbiro_id_fkey":
				return nil, response.ErrNotFound.WithMessage("Kemenbiro tidak ditemukan")
			case "program_kerjas_slug_key":
				return nil, response.ErrConflictSlug
			case "program_kerja_penanggung_jawabs_nim_fkey":
				return nil, programkerja.ErrPenanggungJawabBemMemberNotExists
			}
		}

		log.GetLogger().WithFields(map[string]interface{}{
			"error":   err,
			"request": req,
		}).Errorln("[ProgramKerjaService][CreateProgramKerja] fail to create program kerja")
		return nil, response.ErrInternalServerError
	}

	return &entity.ProgramKerja{ID: id}, nil
}

func (s *programKerjaService) GetProgramKerjasByKemenbiroID(ctx context.Context, req *programkerja.GetProgramKerjasByKemenbiroIDRequest) ([]*entity.ProgramKerja, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	programKerjas, err := s.r.GetProgramKerjasByKemenbiroID(ctx, req.KemenbiroID)
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":        err,
			"kemenbiro_id": req.KemenbiroID,
		}).Errorln("[ProgramKerjaService][GetProgramKerjasByKemenbiroID] fail to get program kerjas by kemenbiro id")
		return nil, response.ErrInternalServerError
	}

	return programKerjas, nil
}
