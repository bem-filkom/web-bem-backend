package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/utils"
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

func (s *userService) GetUserByNIM(ctx context.Context, req *user.GetUserRequest) (*entity.User, error) {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][GetUserByNIM] validation error")
		return nil, response.ErrValidation.WithDetail(valErr)
	}

	userObj, err := s.r.GetUserByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.ErrNotFound
		}
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   err.Error(),
			"request": req,
		}).Error("[UserService][GetUserByNIM] fail to get user from database")
		return nil, response.ErrInternalServerError
	}
	return userObj, nil
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

func (s *userService) GetStudentByNIM(ctx context.Context, req *user.GetUserRequest) (*entity.Student, error) {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][GetStudentByNIM] validation error")
		return nil, response.ErrValidation.WithDetail(valErr)
	}

	student, err := s.r.GetStudentByNIM(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.ErrNotFound
		}
		log.GetLogger().WithFields(map[string]interface{}{
			"error":   err.Error(),
			"request": req,
		}).Error("[UserService][GetStudentByNIM] fail to get student from database")
		return nil, response.ErrInternalServerError
	}
	return student, nil
}

func (s *userService) CreateBemMember(ctx context.Context, req *user.CreateBemMemberRequest) error {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][CreateBemMember] validation error")
		return response.ErrValidation.WithDetail(valErr)
	}

	if err := utils.RequireKemenbiroID(ctx, req.KemenbiroID); err != nil {
		return err
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
			} else if pgErr.Code == "23505" {
				return user.ErrAlreadyBemMember
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

func (s *userService) GetBemMemberByNIM(ctx context.Context, req *user.GetUserRequest) (*entity.BemMember, error) {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][GetBemMemberByNIM] validation error")
		return nil, response.ErrValidation.WithDetail(valErr)
	}

	bemMember, err := s.r.GetBemMemberByNIM(ctx, req.ID)
	if err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   err.Error(),
			"request": req,
		}).Error("[UserService][GetBemMemberByNIM] fail to get BEM member from database")
		return nil, response.ErrInternalServerError
	}
	return bemMember, nil
}

func (s *userService) GetRole(ctx context.Context, req *user.GetUserRequest) (entity.UserRole, error) {
	if valErr := validator.GetValidator().ValidateStruct(req); valErr != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   valErr.Error(),
			"request": req,
		}).Error("[UserService][GetRole] validation error")
		return entity.RoleUnregistered, response.ErrValidation.WithDetail(valErr)
	}

	role, err := s.r.GetRole(ctx, req.ID)
	if err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error":   err.Error(),
			"request": req,
		}).Error("[UserService][GetRole] fail to get role from database")
		return role, response.ErrInternalServerError
	}
	return role, nil
}
