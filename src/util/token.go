package util

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/revandpratama/task-hub/config"
)

type JWTCustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (*int, *string, error) {
	secretKey := []byte(config.ENV.JWT_SECRET_KEY)

	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, nil, errors.New("invalid token signature")
		}

		return nil, nil, errors.New("token expired")
	}

	claims, ok := token.Claims.(*JWTCustomClaims)

	if !ok || !token.Valid {
		return nil, nil, errors.New("token expired")
	}

	return &claims.UserID, &claims.Role, nil
}
