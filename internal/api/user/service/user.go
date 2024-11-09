package service

import (
	"context"
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
)

func (s *userService) SaveUser(ctx context.Context, req *user.SaveUserRequest) error {
	if valErr := validator.GetValidator().Validate(req); valErr != nil {
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
	if valErr := validator.GetValidator().Validate(req); valErr != nil {
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
