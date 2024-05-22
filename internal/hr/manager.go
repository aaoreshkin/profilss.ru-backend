package hr

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/hr/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/hr/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/hr/usecase"
)

type Manager struct {
	HrRepository repository.HrRepository
	HrUsecase    usecase.HrUsecase
	HrController controller.HrController
}

func NewManager(database *database.Database) *Manager {
	hrRepository := repository.NewHrRepository(database)
	hrUsecase := usecase.NewHrUsecase(hrRepository)
	hrController := controller.NewHrController(hrUsecase)

	return &Manager{
		HrRepository: *hrRepository,
		HrUsecase:    *hrUsecase,
		HrController: *hrController,
	}
}
