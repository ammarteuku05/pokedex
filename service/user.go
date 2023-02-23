package service

import (
	"errors"
	"pokedex/entity"
	"pokedex/repository"
	"pokedex/shared/dto"
	"pokedex/shared/mapper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name UserServices --case snake --output  ../mocks/mocksServices --disable-version-string
type UserServices interface {
	RegisterUser(req *dto.UserInputRegister) (*dto.UserResponse, error)
	LoginUser(req *dto.UserInputLogin) (*dto.UserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repositoryUser repository.UserRepository) *userService {
	return &userService{repositoryUser}
}

func (s *userService) RegisterUser(req *dto.UserInputRegister) (*dto.UserResponse, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	newUser := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(genPassword),
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	response, err := s.userRepository.Create(&newUser)

	if err != nil {
		return nil, err
	}

	return mapper.MapResponseUser(response), nil
}

func (s *userService) LoginUser(req *dto.UserInputLogin) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := "user is not found"
		return nil, errors.New(newError)
	}

	// pengecekan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("password invalid")
	}

	return mapper.MapResponseUser(user), nil
}
