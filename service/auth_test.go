package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/betawulan/efishery/error_message"
	mocks "github.com/betawulan/efishery/mock"
	"github.com/betawulan/efishery/model"
)

func Test_AuthService_Register(t *testing.T) {
	ctx := context.Background()

	type getUser struct {
		ctx      context.Context
		filter   model.UserFilter
		response model.User
		err      error
	}

	type create struct {
		ctx  context.Context
		user model.User
		err  error
	}

	tests := []struct {
		name        string
		argCtx      context.Context
		argUser     model.User
		user        getUser
		register    create
		expResponse model.UserResponse
		expErr      error
	}{
		{
			name:   "error while GetUser",
			argCtx: ctx,
			argUser: model.User{
				ID:    1,
				Phone: "087658456886",
				Name:  "chiara",
				Role:  "guest",
			},
			user: getUser{
				ctx:    ctx,
				filter: model.UserFilter{Phone: "087658456886"},
				response: model.User{
					ID:    1,
					Phone: "087658456886",
					Name:  "chiara",
					Role:  "guest",
				},
				err: errors.New("sql: no rows in result set"),
			},
			expResponse: model.UserResponse{},
			expErr:      errors.New("sql: no rows in result set"),
		},
		{
			name:   "the phone already exist",
			argCtx: ctx,
			argUser: model.User{
				ID:    1,
				Phone: "087658456886",
				Name:  "chiara",
				Role:  "guest",
			},
			user: getUser{
				ctx:    ctx,
				filter: model.UserFilter{Phone: "087658456886"},
				response: model.User{
					ID:    1,
					Phone: "087658456886",
					Name:  "chiara",
					Role:  "guest",
				},
				err: nil,
			},
			expResponse: model.UserResponse{},
			expErr:      error_message.Duplicate{Message: "the phone already exist"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := new(mocks.AuthRepository)

			userRepoMock.On("GetUser", test.user.ctx, test.user.filter).
				Return(test.user.response, test.user.err).
				Once()
			userRepoMock.On("Register", test.register.ctx, test.register.user).
				Return(test.register.err).
				Once()

			authService := NewAuthService(userRepoMock, []byte("sidcfkjghkscoedjfmcfklm"))
			response, err := authService.Register(test.argCtx, test.argUser)
			if err != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.expResponse, response)
		})
	}
}
