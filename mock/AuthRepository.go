// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/betawulan/efishery/model"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: ctx, filter
func (_m *AuthRepository) GetUser(ctx context.Context, filter model.UserFilter) (model.User, error) {
	ret := _m.Called(ctx, filter)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, model.UserFilter) model.User); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, phone, password
func (_m *AuthRepository) Login(ctx context.Context, phone string, password string) (model.User, error) {
	ret := _m.Called(ctx, phone, password)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) model.User); ok {
		r0 = rf(ctx, phone, password)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, phone, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, user
func (_m *AuthRepository) Register(ctx context.Context, user model.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAuthRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepository(t mockConstructorTestingTNewAuthRepository) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
