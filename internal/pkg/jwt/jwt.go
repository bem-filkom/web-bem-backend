package jwt

import "C"
import (
	"github.com/bem-filkom/web-bem-backend/internal/pkg/env"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	KemenbiroID string `json:"kemenbiro_id,omitempty"`
}

type CreateRequest struct {
	Subject     string
	KemenbiroID string
}

func Create(req *CreateRequest) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   req.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(env.GetEnv().JwtAccessExpireTime)),
		},
	}

	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJWT, err := unsignedJWT.SignedString(env.GetEnv().JwtAccessSecretKey)
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func Decode(tokenString string, claims *Claims) error {
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
