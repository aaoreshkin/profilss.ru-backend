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

func (controller *PermissionController) Get(rule string) (string, error) {
	result, err := controller.usecase.Get(rule)
	if err != nil {
		return "", err
	}

	return result.Rule, nil
}
