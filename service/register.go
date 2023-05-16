package service

import (
	"context"
	"database/sql"
	"strings"

	_uuid "github.com/satori/go.uuid"

	"github.com/betawulan/efishery/model"
	"github.com/betawulan/efishery/packages/error_message"
	"github.com/betawulan/efishery/repository"
)

type registerService struct {
	registerRepo repository.RegisterRepository
}

func (r registerService) Register(ctx context.Context, register model.Register) (model.RegisterResponse, error) {
	getUser, err := r.registerRepo.GetUser(ctx, model.RegisterFilter{Phone: register.Phone})
	if err != nil {
		if err != sql.ErrNoRows {
			return model.RegisterResponse{}, err
		}
	}

	if getUser.Phone != "" {
		return model.RegisterResponse{}, error_message.Unauthorized{Message: "the phone already exist"}
	}

	uuidPassword := _uuid.NewV4().String()
	password := strings.Split(uuidPassword, "-")
	if len(password) < 1 {
		return model.RegisterResponse{}, error_message.Failed{Message: "failed generate password"}
	}

	register.Password = password[1]

	err = r.registerRepo.Register(ctx, register)
	if err != nil {
		return model.RegisterResponse{}, err
	}

	return model.RegisterResponse{
		Password: register.Password,
	}, nil
}

func NewRegisterService(registerRepo repository.RegisterRepository) RegisterService {
	return registerService{
		registerRepo: registerRepo,
	}
}
