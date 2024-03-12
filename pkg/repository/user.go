package repository

import (
	"test-grpc-project/pkg/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *model.User) error
	DeleteUser(id uint32) error
	GetUser(id uint32) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) DeleteUser(id uint32) error {
	return r.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *userRepository) GetUser(id uint32) (*model.User, error) {
	var user model.User
	err := r.db.Where("id =?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
