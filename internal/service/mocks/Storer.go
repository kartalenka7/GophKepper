// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "keeper/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: ctx, login, password
func (_m *Storer) AddUser(ctx context.Context, login string, password [32]byte) error {
	ret := _m.Called(ctx, login, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, [32]byte) error); ok {
		r0 = rf(ctx, login, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeData provides a mock function with given fields: ctx, data
func (_m *Storer) ChangeData(ctx context.Context, data model.DataBlock) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DataBlock) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckUserAuth provides a mock function with given fields: ctx, login, password
func (_m *Storer) CheckUserAuth(ctx context.Context, login string, password string) error {
	ret := _m.Called(ctx, login, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, login, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteData provides a mock function with given fields: ctx, login, dataKeyWord
func (_m *Storer) DeleteData(ctx context.Context, login string, dataKeyWord string) error {
	ret := _m.Called(ctx, login, dataKeyWord)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, login, dataKeyWord)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetData provides a mock function with given fields: ctx, login, dataKeyWord
func (_m *Storer) GetData(ctx context.Context, login string, dataKeyWord string) ([]model.DataBlock, error) {
	ret := _m.Called(ctx, login, dataKeyWord)

	var r0 []model.DataBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]model.DataBlock, error)); ok {
		return rf(ctx, login, dataKeyWord)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []model.DataBlock); ok {
		r0 = rf(ctx, login, dataKeyWord)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DataBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, dataKeyWord)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertData provides a mock function with given fields: ctx, data
func (_m *Storer) InsertData(ctx context.Context, data model.DataBlock) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DataBlock) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStorer creates a new instance of Storer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storer {
	mock := &Storer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
