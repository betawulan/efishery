package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	_uuid "github.com/satori/go.uuid"

	"github.com/betawulan/efishery/error_message"
	"github.com/betawulan/efishery/model"
	"github.com/betawulan/efishery/repository"
)

type authService struct {
	authRepo  repository.AuthRepository
	SecretKey []byte
}

type claims struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	jwt.StandardClaims
}

func (a authService) Register(ctx context.Context, register model.Register) (model.RegisterResponse, error) {
	getUser, err := a.authRepo.GetUser(ctx, model.RegisterFilter{Phone: register.Phone})
	if err != nil {
		if err != sql.ErrNoRows {
			return model.RegisterResponse{}, err
		}
	}

	if getUser.Phone != "" {
		return model.RegisterResponse{}, error_message.Duplicate{Message: "the phone already exist"}
	}

	uuidPassword := _uuid.NewV4().String()
	password := strings.Split(uuidPassword, "-")
	if len(password) < 1 {
		return model.RegisterResponse{}, error_message.Failed{Message: "failed generate password"}
	}

	register.Password = password[1]

	err = a.authRepo.Register(ctx, register)
	if err != nil {
		return model.RegisterResponse{}, err
	}

	return model.RegisterResponse{
		Password: register.Password,
	}, nil
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

	claim := claims{
		Name:           user.Name,
		Phone:          user.Phone,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	tokenString, err := token.SignedString(a.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewAuthService(authRepo repository.AuthRepository, secretKey []byte) AuthService {
	return authService{
		authRepo:  authRepo,
		SecretKey: secretKey,
	}
}
