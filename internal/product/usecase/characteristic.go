package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type CharacteristicUsecase struct {
	repository entity.CharacteristicRepository
}

func NewCharacteristicUsecase(repository entity.CharacteristicRepository) *CharacteristicUsecase {
	return &CharacteristicUsecase{
		repository: repository,
	}
}

func (usecase *CharacteristicUsecase) Create(entity *entity.Characteristic) (*entity.Characteristic, error) {
	return usecase.repository.Create(entity)
}

func (usecase *CharacteristicUsecase) Find() ([]entity.Characteristic, error) {
	return usecase.repository.Find()
}

func (usecase *CharacteristicUsecase) First(id string) (*entity.Characteristic, error) {
	return usecase.repository.First(id)
}

func (usecase *CharacteristicUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
