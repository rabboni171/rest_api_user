package service

import (
	"restApi/internal/models"
	"restApi/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	GetUserByID(id string) (*models.User, error)
}

type Service struct {
	Authorization
}

func NewUserService(repos repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization)}
}
