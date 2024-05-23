package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type IsoUsecase struct {
	repository entity.IsoRepository
}

func NewIsoUsecase(repository entity.IsoRepository) *IsoUsecase {
	return &IsoUsecase{
		repository: repository,
	}
}

func (usecase *IsoUsecase) Create(entity *entity.Iso) (*entity.Iso, error) {
	return usecase.repository.Create(entity)
}

func (usecase *IsoUsecase) Find() ([]entity.Iso, error) {
	return usecase.repository.Find()
}

func (usecase *IsoUsecase) First(id string) (*entity.Iso, error) {
	return usecase.repository.First(id)
}

func (usecase *IsoUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
