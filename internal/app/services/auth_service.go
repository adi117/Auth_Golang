package services

import "auth/internal/app/models"

type AuthServiceInterface interface {
	Register(user models.UserRegisterRequest) (*models.User, error)
	Login(user models.UserLoginRequest) (*models.User, error)
}
