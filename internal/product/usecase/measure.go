package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type MeasureUsecase struct {
	repository entity.MeasureRepository
}

func NewMeasureUsecase(repository entity.MeasureRepository) *MeasureUsecase {
	return &MeasureUsecase{
		repository: repository,
	}
}

func (usecase *MeasureUsecase) Create(entity *entity.Measure) (*entity.Measure, error) {
	return usecase.repository.Create(entity)
}

func (usecase *MeasureUsecase) Find() ([]entity.Measure, error) {
	return usecase.repository.Find()
}

func (usecase *MeasureUsecase) First(id string) (*entity.Measure, error) {
	return usecase.repository.First(id)
}

func (usecase *MeasureUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
