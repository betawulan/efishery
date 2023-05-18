package service

import (
	"context"
	"github.com/betawulan/efishery/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.User) (model.UserResponse, error)
	Login(ctx context.Context, phone string, password string) (string, error)
	Validate(token string) (model.User, error)
}

type FishService interface {
	GetDataStorages(token string) ([]model.Fish, error)
	Summary(ctx context.Context, token string) ([]model.Summary, error)
}
