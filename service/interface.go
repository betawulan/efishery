package service

import (
	"context"
	"github.com/betawulan/efishery/model"
)

type AuthService interface {
	Register(ctx context.Context, register model.User) (model.UserResponse, error)
	Login(ctx context.Context, phone string, password string) (string, error)
}
