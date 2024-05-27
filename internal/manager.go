package internal

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/bid"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/doc"
	"github.com/oreshkindev/profilss.ru-backend/internal/hr"
	"github.com/oreshkindev/profilss.ru-backend/internal/post"
	"github.com/oreshkindev/profilss.ru-backend/internal/product"
	"github.com/oreshkindev/profilss.ru-backend/internal/service"
	"github.com/oreshkindev/profilss.ru-backend/internal/setting"
	"github.com/oreshkindev/profilss.ru-backend/internal/support"
	"github.com/oreshkindev/profilss.ru-backend/internal/user"
)

type Manager struct {
	Support *support.Manager
	Doc     *doc.Manager
	Bid     *bid.Manager
	Post    *post.Manager
	Service *service.Manager
	Setting *setting.Manager
	User    *user.Manager
	Product *product.Manager
	Hr      *hr.Manager
}

func NewManager(database *database.Database) *Manager {
	return &Manager{
		Support: support.NewManager(database),
		Doc:     doc.NewManager(),
		Bid:     bid.NewManager(database),
		Post:    post.NewManager(database),
		Service: service.NewManager(database),
		Setting: setting.NewManager(database),
		User:    user.NewManager(database),
		Product: product.NewManager(database),
		Hr:      hr.NewManager(database),
	}
}
