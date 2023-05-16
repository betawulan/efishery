package repository

import (
	"context"

	"github.com/betawulan/efishery/model"
)

type AuthRepository interface {
	Register(ctx context.Context, register model.Register) error
	GetUser(ctx context.Context, filter model.RegisterFilter) (model.Register, error)
	Login(ctx context.Context, phone string, password string) (model.Register, error)
}
