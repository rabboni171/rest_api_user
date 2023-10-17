package repository

import (
	"restApi/internal/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	GetUserByID(id string) (*models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
