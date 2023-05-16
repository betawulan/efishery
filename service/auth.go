package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/betawulan/efishery/packages/auth"
	"github.com/betawulan/efishery/packages/error_message"
	"github.com/betawulan/efishery/repository"
)

type authService struct {
	authRepo repository.AuthRepository
	jwt      auth.Auth
}

func (a authService) Login(ctx context.Context, phone string, password string) (string, error) {
	user, err := a.authRepo.Login(ctx, phone, password)
	if err != nil {
		switch _err := err.(type) {
		case error_message.NotFound:
			return "", error_message.Unauthorized{
				Message: _err.Error(),
			}
		default:
			return "", err
		}
	}

	claim := auth.Claim{
		Name:           user.Name,
		Phone:          user.Phone,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix()},
	}

	token, err := a.jwt.Encode(claim)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewAuthService(authRepo repository.AuthRepository, jwt auth.Auth) AuthService {
	return authService{
		authRepo: authRepo,
		jwt:      jwt,
	}
}
