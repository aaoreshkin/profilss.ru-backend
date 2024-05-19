package doc

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/doc/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/doc/usecase"
)

type Manager struct {
	DocUsecase    usecase.DocUsecase
	DocController controller.DocController
}

func NewManager() *Manager {
	docUsecase := usecase.NewDocUsecase()
	docController := controller.NewDocController(docUsecase)

	return &Manager{
		DocUsecase:    *docUsecase,
		DocController: *docController,
	}
}
