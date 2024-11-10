package service

import (
	"context"
	"errors"
	ubauth "github.com/ahmdyaasiin/ub-auth-without-notification/v2"
	"github.com/bem-filkom/web-bem-backend/internal/api/auth"
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/jwt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
)

func (s *authService) LoginAuthUb(ctx context.Context, req *auth.LoginRequest) (string, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return "", response.ErrValidation.WithDetail(err)
	}

	studentDetails, err := s.ubAuth.AuthUB(req.Username, req.Password)
	if err != nil {
		var respErr *ubauth.ResponseDetails
		if !errors.As(err, &respErr) {
			log.GetLogger().WithFields(map[string]interface{}{
				"error":    err.Error(),
				"username": req.Username,
			}).Error("[AuthService][LoginAuthUb] fail to authenticate user")
		}

		switch respErr.Message {
		case "Invalid username or password":
			return "", auth.ErrInvalidCredentials
		default:
			log.GetLogger().WithFields(map[string]interface{}{
				"error":    respErr.Error(),
				"username": req.Username,
			}).Error("[AuthService][LoginAuthUb] fail to authenticate user")
			return "", response.ErrInternalServerError
		}
	}

	if err := s.us.SaveStudent(ctx, &user.SaveStudentRequest{
		NIM:          studentDetails.NIM,
		Email:        studentDetails.Email,
		FullName:     studentDetails.FullName,
		Fakultas:     studentDetails.Fakultas,
		ProgramStudi: studentDetails.ProgramStudi,
	}); err != nil {
		return "", err
	}

	accessToken, err := jwt.Create(&jwt.CreateRequest{
		Subject: studentDetails.NIM,
	})
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":    err.Error(),
			"username": req.Username,
		}).Error("[AuthService][LoginAuthUb] fail to create jwt access token")
		return "", response.ErrInternalServerError
	}
	return accessToken, nil
}
