package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/setting/entity"
)

type SettingUsecase struct {
	repository entity.SettingRepository
}

func NewSettingUsecase(repository entity.SettingRepository) *SettingUsecase {
	return &SettingUsecase{
		repository: repository,
	}
}

func (usecase *SettingUsecase) Create(entity *entity.Setting) (*entity.Setting, error) {
	return usecase.repository.Create(entity)
}

func (usecase *SettingUsecase) Find() ([]entity.Setting, error) {
	return usecase.repository.Find()
}

func (usecase *SettingUsecase) First(id string) (*entity.Setting, error) {
	return usecase.repository.First(id)
}

func (usecase *SettingUsecase) Update(entity *entity.Setting, id string) (*entity.Setting, error) {
	return usecase.repository.Update(entity, id)
}

func (usecase *SettingUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
