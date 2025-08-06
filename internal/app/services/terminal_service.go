package services

import (
	"auth/internal/app/models"
	"auth/internal/app/repositories"
)

type TerminalServiceInterface interface {
	CreateTerminal(terminalName string) (*models.Terminal, error)
}

type TerminalService struct {
	Repo repositories.TerminalRepositoryInterface
}

func NewTerminalService(repo repositories.TerminalRepositoryInterface) TerminalServiceInterface {
	return &TerminalService{Repo: repo}
}

func (s *TerminalService) CreateTerminal(terminalName string) (*models.Terminal, error) {
	return s.Repo.CreateTerminal(terminalName)
}
