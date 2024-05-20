package service

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/service/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/service/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/service/usecase"
)

type Manager struct {
	ServiceRepository repository.ServiceRepository
	ServiceUsecase    usecase.ServiceUsecase
	ServiceController controller.ServiceController
}

func NewManager(database *database.Database) *Manager {
	serviceRepository := repository.NewServiceRepository(database)
	serviceUsecase := usecase.NewServiceUsecase(serviceRepository)
	serviceController := controller.NewServiceController(serviceUsecase)

	return &Manager{
		ServiceRepository: *serviceRepository,
		ServiceUsecase:    *serviceUsecase,
		ServiceController: *serviceController,
	}
}
