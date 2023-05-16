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

func (a authService) Register(ctx context.Context, user model.User) (model.UserResponse, error) {
	getUser, err := a.authRepo.GetUser(ctx, model.UserFilter{Phone: user.Phone})
	if err != nil {
		if err != sql.ErrNoRows {
			return model.UserResponse{}, err
		}
	}

	if getUser.Phone != "" {
		return model.UserResponse{}, error_message.Duplicate{Message: "the phone already exist"}
	}

	uuidPassword := _uuid.NewV4().String()
	password := strings.Split(uuidPassword, "-")
	if len(password) < 1 {
		return model.UserResponse{}, error_message.Failed{Message: "failed generate password"}
	}

	user.Password = password[1]

	err = a.authRepo.Register(ctx, user)
	if err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{
		Password: user.Password,
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

func (a authService) Validate(tokenString string) (model.User, error) {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return a.SecretKey, nil
	})
	if err != nil {
		return model.User{}, err
	}

	if !token.Valid {
		return model.User{}, error_message.Unauthorized{Message: "token invalid"}
	}

	return model.User{
		Phone:     claim.Phone,
		Name:      claim.Name,
		Role:      claim.Role,
		CreatedAt: claim.CreatedAt,
	}, nil

}

func NewAuthService(authRepo repository.AuthRepository, secretKey []byte) AuthService {
	return authService{
		authRepo:  authRepo,
		SecretKey: secretKey,
	}
}
