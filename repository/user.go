package repository

import (
	"pokedex/entity"

	"gorm.io/gorm"
)

//go:generate mockery --name UserRepository --case snake --output ../mocks/mocksRepository --disable-version-string
type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByID(ID int) (*entity.User, error)
	UpdateByID(ID int, dataUpdate map[string]interface{}) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID int) (*entity.User, error) {
	var user *entity.User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateByID(ID int, dataUpdate map[string]interface{}) (*entity.User, error) {

	var user *entity.User

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
