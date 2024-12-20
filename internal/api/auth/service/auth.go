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

func (s *authService) LoginUB(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	if err := validator.GetValidator().ValidateStruct(req); err != nil {
		return nil, response.ErrValidation.WithDetail(err)
	}

	var res auth.LoginResponse

	studentDetails, err := s.ubAuth.AuthUB(req.Username, req.Password)
	if err != nil {
		var respErr *ubauth.ResponseDetails
		if !errors.As(err, &respErr) {
			log.GetLogger().WithFields(map[string]interface{}{
				"error":    err.Error(),
				"username": req.Username,
			}).Error("[AuthService][LoginUB] fail to authenticate user")
		}

		switch respErr.Message {
		case "Invalid username or password":
			return nil, auth.ErrInvalidCredentials
		default:
			log.GetLogger().WithFields(map[string]interface{}{
				"error":    respErr.Error(),
				"username": req.Username,
			}).Error("[AuthService][LoginUB] fail to authenticate user")
			return nil, response.ErrInternalServerError
		}
	}

	student := &entity.Student{
		NIM: studentDetails.NIM,
		User: &entity.User{
			Email:    studentDetails.Email,
			FullName: studentDetails.FullName,
		},
		ProgramStudi: studentDetails.ProgramStudi,
		Fakultas:     studentDetails.Fakultas,
	}

	if err := s.us.SaveStudent(ctx, &user.SaveStudentRequest{
		NIM:          studentDetails.NIM,
		Email:        studentDetails.Email,
		FullName:     studentDetails.FullName,
		Fakultas:     studentDetails.Fakultas,
		ProgramStudi: studentDetails.ProgramStudi,
	}); err != nil {
		return nil, err
	}

	jwtCreateReq := &jwt.CreateRequest{
		Subject: studentDetails.NIM,
	}

	role, err := s.us.GetRole(ctx, &user.GetUserRequest{ID: studentDetails.NIM})
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":           err.Error(),
			"student_details": studentDetails,
		}).Error("[AuthService][LoginUB] fail to get role")
		return nil, response.ErrInternalServerError
	}
	jwtCreateReq.Role = role
	res.Role = role

	if role == entity.RoleBemMember {
		bemMember, err := s.us.GetBemMemberByNIM(ctx, &user.GetUserRequest{ID: studentDetails.NIM})
		if err != nil {
			log.GetLogger().WithFields(map[string]interface{}{
				"error":           err.Error(),
				"student_details": studentDetails,
			}).Error("[AuthService][LoginUB] fail to get bem member")
			return nil, response.ErrInternalServerError
		}
		jwtCreateReq.KemenbiroID = bemMember.KemenbiroID
		jwtCreateReq.KemenbiroAbbreviation = bemMember.Kemenbiro.Abbreviation

		student.NIM = ""
		res.BemMember = &entity.BemMember{
			NIM:         studentDetails.NIM,
			Student:     student,
			KemenbiroID: bemMember.KemenbiroID,
			Kemenbiro: &entity.Kemenbiro{
				Name:         bemMember.Kemenbiro.Name,
				Abbreviation: bemMember.Kemenbiro.Abbreviation,
			},
			Position: bemMember.Position,
			Period:   bemMember.Period,
		}
	} else {
		res.Student = student
	}

	accessToken, err := jwt.CreateAccessToken(jwtCreateReq)
	if err != nil {
		log.GetLogger().WithFields(map[string]interface{}{
			"error":    err.Error(),
			"username": req.Username,
		}).Error("[AuthService][LoginUB] fail to create jwt access token")
		return nil, response.ErrInternalServerError
	}
	res.AccessToken = accessToken

	return &res, nil
}
