package jwt

import "C"
import (
	"encoding/json"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/entity"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/env"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	IsSuperAdmin          bool            `json:"is_super_admin"`
	Role                  entity.UserRole `json:"role"`
	KemenbiroID           uuid.UUID       `json:"kemenbiro_id,omitempty"`
	KemenbiroAbbreviation string          `json:"kemenbiro_abbreviation,omitempty"`
}

func (c Claims) MarshalJSON() ([]byte, error) {
	type Alias Claims
	aux := &struct {
		KemenbiroID *uuid.UUID `json:"kemenbiro_id,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&c),
	}

	if c.KemenbiroID != uuid.Nil {
		aux.KemenbiroID = &c.KemenbiroID
	}

	return json.Marshal(aux)
}

type CreateRequest struct {
	Subject               string
	IsSuperAdmin          bool
	Role                  entity.UserRole
	KemenbiroID           uuid.UUID
	KemenbiroAbbreviation string
}

func CreateAccessToken(req *CreateRequest) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   req.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(env.GetEnv().JwtAccessExpireTime)),
		},
		Role:                  req.Role,
		KemenbiroID:           req.KemenbiroID,
		KemenbiroAbbreviation: req.KemenbiroAbbreviation,
		IsSuperAdmin:          req.KemenbiroAbbreviation == "PIT",
	}

	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJWT, err := unsignedJWT.SignedString(env.GetEnv().JwtAccessSecretKey)
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func DecodeAccessToken(tokenString string, claims *Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return env.GetEnv().JwtAccessSecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	return nil
}
