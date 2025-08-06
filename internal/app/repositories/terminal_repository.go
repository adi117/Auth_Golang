package repositories

import (
	"auth/internal/app/models"
	"time"

	"gorm.io/gorm"
)

type TerminalRepository struct {
	DB *gorm.DB
}

func NewTerminalRepository(db *gorm.DB) TerminalRepositoryInterface {
	return &TerminalRepository{DB: db}
}

func (r *TerminalRepository) CreateTerminal(terminalName string) (*models.Terminal, error) {
	newTerminal := models.Terminal{
		Name:      terminalName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := r.DB.Create(&newTerminal).Error

	if err != nil {
		return nil, err
	}

	return &newTerminal, nil
}
