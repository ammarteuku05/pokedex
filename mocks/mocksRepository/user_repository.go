// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	entity "pokedex/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserRepository) Create(user *entity.User) (*entity.User, error) {
	ret := _m.Called(user)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: email
func (_m *UserRepository) FindByEmail(email string) (*entity.User, error) {
	ret := _m.Called(email)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ID
func (_m *UserRepository) FindByID(ID int) (*entity.User, error) {
	ret := _m.Called(ID)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(int) *entity.User); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ID, dataUpdate
func (_m *UserRepository) UpdateByID(ID int, dataUpdate map[string]interface{}) (*entity.User, error) {
	ret := _m.Called(ID, dataUpdate)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) *entity.User); ok {
		r0 = rf(ID, dataUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, map[string]interface{}) error); ok {
		r1 = rf(ID, dataUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
