package service

import (
	"context"

	"github.com/betawulan/efishery/model"
)

type RegisterService interface {
	Register(ctx context.Context, register model.Register) (model.RegisterResponse, error)
}
