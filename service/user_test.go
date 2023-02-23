package service

import (
	"errors"
	"pokedex/entity"
	mocks "pokedex/mocks/mocksRepository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var entityUser = &entity.User{
	ID:        11,
	Name:      "ammar",
	Email:     "ammar@gmail.com",
	Password:  "12345",
	Role:      "admin",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestLoginService(t *testing.T) {
	mockBorrowerRepo := &mocks.UserRepository{}
	h := NewUserService(mockBorrowerRepo)
	defer mockBorrowerRepo.AssertExpectations(t)

	//
	t.Run("should return nil error", func(t *testing.T) {

		mockBorrowerRepo.On("FindByEmail", entityUser.Email).
			Return(entityUser, nil).Once()

		_, err := h.userRepository.FindByEmail(entityUser.Email)

		assert.NoError(t, err)
	})

	//
	t.Run("should return error", func(t *testing.T) {
		mockBorrowerRepo.On("FindByEmail", entityUser.Email).
			Return(nil, errors.New("user is not found")).Once()

		_, err := h.userRepository.FindByEmail(entityUser.Email)

		assert.Error(t, err)
	})

	mockBorrowerRepo.AssertExpectations(t)
}

func TestRegisterService(t *testing.T) {
	mockBorrowerRepo := &mocks.UserRepository{}
	h := NewUserService(mockBorrowerRepo)
	defer mockBorrowerRepo.AssertExpectations(t)

	//
	t.Run("should return nil error", func(t *testing.T) {

		mockBorrowerRepo.On("Create", entityUser).
			Return(entityUser, nil).Once()

		_, err := h.userRepository.Create(entityUser)

		assert.NoError(t, err)
	})

	//
	t.Run("should return error", func(t *testing.T) {
		mockBorrowerRepo.On("Create", entityUser).
			Return(nil, errors.New("internal server error")).Once()

		_, err := h.userRepository.Create(entityUser)

		assert.Error(t, err)
	})

	mockBorrowerRepo.AssertExpectations(t)
}
