package service

import (
	"context"
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *userService) SaveUser(ctx context.Context, req *user.SaveUserRequest) error {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][SaveUser] validation error")
		return valErr
	}

	if err := s.r.SaveUser(ctx, &entity.User{
		ID:       req.ID,
		Email:    req.Email,
		FullName: req.FullName,
	}); err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   err.Error(),
			"request": req,
		}).Error("[UserService][SaveUser] fail to save user to database")
		return response.ErrInternalServerError
	}
	return nil
}

func (s *userService) SaveStudent(ctx context.Context, req *user.SaveStudentRequest) error {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][SaveStudent] validation error")
		return valErr
	}

	if err := s.r.SaveStudent(ctx, &entity.Student{
		NIM: req.NIM,
		User: &entity.User{
			ID:       req.NIM,
			Email:    req.Email,
			FullName: req.FullName,
		},
		ProgramStudi: req.ProgramStudi,
		Fakultas:     req.Fakultas,
	}); err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   err.Error(),
			"request": req,
		}).Error("[UserService][SaveStudent] fail to save student to database")

		if errors.Is(err, context.DeadlineExceeded) {
			return response.ErrTimeout
		}

		return response.ErrInternalServerError
	}
	return nil
}

func (s *userService) CreateBemMember(ctx context.Context, req *user.CreateBemMemberRequest) error {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][CreateBemMember] validation error")
		return response.ErrValidation.WithDetail(valErr)
	}

	if err := s.r.CreateBemMember(ctx, &entity.BemMember{
		NIM:         req.NIM,
		KemenbiroID: req.KemenbiroID,
		Position:    req.Position,
	}); err != nil {
		var pgErr *pgconn.PgError
		_ = errors.As(err, &pgErr)
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23503" {
				return response.ErrNotFound.WithMessage("Kemenbiro tidak ditemukan.")
			}
		}

		log.GetLogger().WithFields(map[string]any{
			"error":   err.Error(),
			"request": req,
		}).Error("[UserService][CreateBemMember] fail to promote student to BEM member")
		return response.ErrInternalServerError
	}
	return nil
}
