package service

import (
	"context"
	"errors"
	ubauth "github.com/ahmdyaasiin/ub-auth-without-notification/v2"
	"github.com/bem-filkom/web-bem-backend/internal/api/auth"
	"github.com/bem-filkom/web-bem-backend/internal/api/user"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
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

	jwtCreateReq := &jwt.CreateRequest{
		Subject: studentDetails.NIM,
	}

	role, err := s.us.GetRole(ctx, &user.GetUserRequest{ID: studentDetails.NIM})
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":           err.Error(),
			"student_details": studentDetails,
		}).Error("[AuthService][LoginAuthUb] fail to get role")
		return "", response.ErrInternalServerError
	}
	jwtCreateReq.Role = role

	if role == entity.RoleBemMember {
		bemMember, err := s.us.GetBemMemberByNIM(ctx, &user.GetUserRequest{ID: studentDetails.NIM})
		if err != nil {
			log.GetLogger().WithFields(map[string]interface{}{
				"error":           err.Error(),
				"student_details": studentDetails,
			}).Error("[AuthService][LoginAuthUb] fail to get bem member")
			return "", response.ErrInternalServerError
		}
		jwtCreateReq.KemenbiroID = bemMember.KemenbiroID
		jwtCreateReq.KemenbiroAbbreviation = bemMember.Kemenbiro.Abbreviation
	}

	accessToken, err := jwt.CreateAccessToken(jwtCreateReq)
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":    err.Error(),
			"username": req.Username,
		}).Error("[AuthService][LoginAuthUb] fail to create jwt access token")
		return "", response.ErrInternalServerError
	}
	return accessToken, nil
}
