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
	if err := validator.GetValidator().Validate(req); err != nil {
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
