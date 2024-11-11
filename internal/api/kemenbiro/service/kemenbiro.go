package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/api/kemenbiro"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *kemenbiroService) CreateKemenbiro(ctx context.Context, req *kemenbiro.CreateKemenbiroRequest) (*entity.Kemenbiro, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	kemenbiroObj := &entity.Kemenbiro{
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
	}
	if req.Description == "" {
		kemenbiroObj.Description = sql.NullString{}
	}

	id, err := s.r.CreateKemenbiro(ctx, kemenbiroObj)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, kemenbiro.ErrAbbreviationAlreadyExists
		}

		log.GetLogger().WithFields(map[string]any{
			"error":   err,
			"request": req,
		}).Errorln("[KemenbiroService][CreateKemenbiro] fail to create kemenbiro")
		return nil, response.ErrInternalServerError
	}

	return id, nil
}

func (s *kemenbiroService) GetKemenbiroByAbbreviation(ctx context.Context, req *kemenbiro.GetKemenbiroByAbbreviationRequest) (*entity.Kemenbiro, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	kemenbiroObj, err := s.r.GetKemenbiroByAbbreviation(ctx, req.Abbreviation)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.ErrNotFound
		}

		log.GetLogger().WithFields(map[string]any{
			"error": err,
			"id":    req.Abbreviation,
		}).Errorln("[KemenbiroService][GetKemenbiroByAbbreviation] fail to get kemenbiro by abbreviation")
		return nil, response.ErrInternalServerError
	}

	return kemenbiroObj, nil
}

func (s *kemenbiroService) UpdateKemenbiro(ctx context.Context, req *kemenbiro.UpdateKemenbiroRequest) error {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return response.ErrValidation.WithDetail(err)
	}

	updatedKemenbiro := &entity.Kemenbiro{
		Name:         req.Name,
		Abbreviation: req.Abbreviation,
		Description:  sql.NullString{String: req.Description, Valid: req.Description != ""},
	}

	err := s.r.UpdateKemenbiro(ctx, req.AbbreviationAsID, updatedKemenbiro)
	if err != nil {
		if err.Error() == "no fields to update" {
			return response.ErrNoUpdatedField
		}

		if err.Error() == "no rows affected" {
			return response.ErrNotFound
		}

		log.GetLogger().WithFields(map[string]any{
			"error":   err,
			"request": req,
		}).Errorln("[KemenbiroService][UpdateKemenbiro] fail to update kemenbiro")
		return response.ErrInternalServerError
	}

	return nil
}