package repository

import (
	"context"

	"github.com/betawulan/efishery/model"
)

type AuthRepository interface {
	Register(ctx context.Context, register model.User) error
	GetUser(ctx context.Context, filter model.UserFilter) (model.User, error)
	Login(ctx context.Context, phone string, password string) (model.User, error)
}
