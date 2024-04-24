package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

type PermissionUsecase struct {
	repository entity.PermissionRepository
}

func NewPermissionUsecase(repository entity.PermissionRepository) *PermissionUsecase {
	return &PermissionUsecase{
		repository: repository,
	}
}

func (usecase *PermissionUsecase) Get(id string) (*entity.Permission, error) {
	result, err := usecase.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
