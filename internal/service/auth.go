package service

import (
	"restApi/internal/models"
	"restApi/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user models.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) UpdateUser(user models.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.UpdateUser(user)
}

func (s *AuthService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func generatePasswordHash(password string) string {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	strHash := string(byteHash)
	return strHash
}
