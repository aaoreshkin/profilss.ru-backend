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

func (usecase *PermissionUsecase) First(id string) (*entity.Permission, error) {
	return usecase.repository.First(id)
}
