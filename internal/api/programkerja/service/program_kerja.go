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
	"github.com/google/uuid"
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

func (s *programKerjaService) GetProgramKerjaByID(ctx context.Context, req *programkerja.GetProgramKerjaByIDRequest) (*entity.ProgramKerja, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	programKerjaObj, err := s.r.GetProgramKerjaByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.ErrNotFound.WithMessage("Program kerja tidak ditemukan")
		}

		log.GetLogger().WithFields(map[string]interface{}{
			"error": err,
			"id":    req.ID,
		}).Errorln("[ProgramKerjaService][GetProgramKerjaByID] fail to get program kerja by id")
		return nil, response.ErrInternalServerError
	}

	return programKerjaObj, nil
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

func (s *programKerjaService) GetKemenbiroIDByProgramKerjaID(ctx context.Context, prokerID uuid.UUID) (uuid.UUID, error) {
	if err := validator.GetValidator().ValidateVariable(prokerID, "required,uuid"); err != nil {
		return uuid.Nil, response.ErrValidation.WithDetail(err)
	}

	kemenbiroID, err := s.r.GetKemenbiroIDByProgramKerjaID(ctx, prokerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, response.ErrNotFound.WithMessage("Program kerja tidak ditemukan")
		}

		log.GetLogger().WithFields(map[string]interface{}{
			"error": err,
			"id":    prokerID,
		}).Errorln("[ProgramKerjaService][GetKemenbiroIDByProgramKerjaID] fail to get kemenbiro id by program kerja id")
		return uuid.Nil, response.ErrInternalServerError
	}

	return kemenbiroID, nil
}

func (s *programKerjaService) UpdateProgramKerja(ctx context.Context, req *programkerja.UpdateProgramKerjaRequest) error {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return response.ErrValidation.WithDetail(err)
	}

	kemenbiroID, err := s.GetKemenbiroIDByProgramKerjaID(ctx, req.ID)
	if err != nil {
		return err
	}

	if err := utils.RequireKemenbiroID(ctx, kemenbiroID); err != nil {
		return err
	}

	if req.KemenbiroID != uuid.Nil {
		if err := utils.RequireKemenbiroID(ctx, req.KemenbiroID); err != nil {
			return err
		}
	}

	if err := s.r.UpdateProgramKerja(ctx, req); err != nil {
		if err.Error() == "no fields to update" {
			return response.ErrNoFieldsToUpdate
		}
		if err.Error() == "no rows affected" {
			return response.ErrNotFound
		}

		log.GetLogger().WithFields(map[string]interface{}{
			"error":   err,
			"request": req,
		}).Errorln("[ProgramKerjaService][UpdateProgramKerja] fail to update program kerja")
		return response.ErrInternalServerError
	}

	return nil
}
