package controller

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

type PermissionController struct {
	usecase entity.PermissionUsecase
}

func NewPermissionController(usecase entity.PermissionUsecase) *PermissionController {
	return &PermissionController{
		usecase: usecase,
	}
}

func (controller *PermissionController) First(id string) (*entity.Permission, error) {
	return controller.usecase.First(id)
}
