package internal

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/bid"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/post"
	"github.com/oreshkindev/profilss.ru-backend/internal/user"
)

type Manager struct {
	Bid  *bid.Manager
	Post *post.Manager
	User *user.Manager
}

func NewManager(database *database.Database) *Manager {
	return &Manager{
		Bid:  bid.NewManager(database),
		Post: post.NewManager(database),
		User: user.NewManager(database),
	}
}
