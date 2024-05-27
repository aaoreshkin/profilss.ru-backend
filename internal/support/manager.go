package support

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/support/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/support/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/support/usecase"
)

type Manager struct {
	SupportRepository repository.SupportRepository
	SupportUsecase    usecase.SupportUsecase
	SupportController controller.SupportController
}

func NewManager(database *database.Database) *Manager {
	supportRepository := repository.NewSupportRepository(database)
	supportUsecase := usecase.NewSupportUsecase(supportRepository)
	supportController := controller.NewSupportController(supportUsecase)

	return &Manager{
		SupportRepository: *supportRepository,
		SupportUsecase:    *supportUsecase,
		SupportController: *supportController,
	}
}
