package repositories

import "auth/internal/app/models"

type AuthRepositoryInterface interface {
	Signup(user models.UserRegisterRequest) (models.User, error)
	Login(user models.UserLoginRequest) (models.User, error)
}
