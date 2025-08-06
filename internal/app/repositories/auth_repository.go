package repositories

import (
	"auth/internal/app/models"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepositoryInterface {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) Register(user models.UserRegisterRequest) (*models.User, error) {
	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	err := r.DB.Create(&newUser).Error

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r *AuthRepository) Login(user models.UserLoginRequest) (*models.User, error) {
	var dbUser models.User

	err := r.DB.Where("email = ?", user.Email).First(&dbUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &dbUser, nil
}
