package repository

import (
	"context"

	"github.com/betawulan/efishery/model"
)

type RegisterRepository interface {
	Register(ctx context.Context, register model.Register) error
	GetUser(ctx context.Context, filter model.RegisterFilter) (model.Register, error)
}