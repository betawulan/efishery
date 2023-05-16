package auth

import (
	"github.com/golang-jwt/jwt"

	"github.com/betawulan/efishery/packages/error_message"
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

func (a auth) Decode(tokenString string) (Claim, error) {
	c := Claim{}

	token, err := jwt.ParseWithClaims(tokenString, &c, func(token *jwt.Token) (interface{}, error) {
		return a.SecretKey, nil
	})
	if err != nil {
		return Claim{}, err
	}

	if !token.Valid {
		return Claim{}, error_message.Unauthorized{Message: "token invalid"}
	}

	return c, nil
}

func NewAuth(secretKey []byte) Auth {
	return auth{
		SecretKey: secretKey,
	}
}
