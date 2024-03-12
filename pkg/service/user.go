package service

import (
	"test-grpc-project/pkg/model"
	"test-grpc-project/pkg/repository"
	"test-grpc-project/pkg/utils"
	"time"
)

type IUserService interface {
	CreateUser(user *model.User) error
	DeleteUser(id uint32) error
	GetUser(id uint32) (*model.User, error)
}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *model.User) error {
	hashadPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashadPassword)
	user.CreatedAt = time.Now()
	return s.repo.CreateUser(user)
}

func (s *userService) DeleteUser(id uint32) error {
	return s.repo.DeleteUser(id)
}

func (s *userService) GetUser(id uint32) (*model.User, error) {
	return s.repo.GetUser(id)
}
