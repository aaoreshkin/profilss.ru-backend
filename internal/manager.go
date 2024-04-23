package internal

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/user"
)

type Manager struct {
	User *user.Manager
}

func NewManager(database *database.Database) *Manager {
	return &Manager{
		User: user.NewManager(database),
	}
}
