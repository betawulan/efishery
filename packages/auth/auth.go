package auth

import (
	"github.com/golang-jwt/jwt"
)

type auth struct {
	SecretKey []byte
}

func (a auth) Encode(c Claim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &c)
	tokenString, err := token.SignedString(a.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
