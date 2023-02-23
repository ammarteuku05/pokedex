package mapper

import (
	"pokedex/entity"
	"pokedex/shared/dto"
)

func MapResponseUser(user *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		Id:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
