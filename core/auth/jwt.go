package auth

import (
	customerrors "mid/core/errors"
	"mid/core/var/env"

	jwt "github.com/dgrijalva/jwt-go"
)

var key = env.JwtSecretKey

func DecodeJwtToken(tokenString string) (*User, error) {
	claims := &User{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, &customerrors.InvalidToken{}
	}

	return claims, nil
}
