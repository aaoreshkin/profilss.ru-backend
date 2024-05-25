package setting

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/setting/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/setting/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/setting/usecase"
)

type Manager struct {
	SettingRepository repository.SettingRepository
	SettingUsecase    usecase.SettingUsecase
	SettingController controller.SettingController
}

func NewManager(database *database.Database) *Manager {
	settingRepository := repository.NewSettingRepository(database)
	settingUsecase := usecase.NewSettingUsecase(settingRepository)
	settingController := controller.NewSettingController(settingUsecase)

	return &Manager{
		SettingRepository: *settingRepository,
		SettingUsecase:    *settingUsecase,
		SettingController: *settingController,
	}
}
