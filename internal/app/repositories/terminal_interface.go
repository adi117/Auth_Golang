package repositories

import "auth/internal/app/models"

type TerminalRepositoryInterface interface {
	CreateTerminal(terminalName string) (*models.Terminal, error)
}
