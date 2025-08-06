package services

import (
	"auth/internal/app/models"
	"auth/internal/app/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(user models.UserRegisterRequest) (*models.User, error)
	Login(user models.UserLoginRequest) (*models.User, error)
}

type AuthService struct {
	Repo repositories.AuthRepositoryInterface
}

func NewAuthService(repo repositories.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(user models.UserRegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return s.Repo.Register(user)
}

func (s *AuthService) Login(user models.UserLoginRequest) (*models.User, error) {
	dbUser, err := s.Repo.Login(user)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("Invalid credentials!")
	}

	return dbUser, nil
}
