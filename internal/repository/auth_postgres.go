package repository

import (
	"restApi/internal/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (repo *AuthPostgres) CreateUser(user models.User) error {
	return repo.db.Create(user).Error
}

func (repo *AuthPostgres) UpdateUser(user models.User) error {
	return repo.db.Save(user).Error
}

func (repo *AuthPostgres) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
