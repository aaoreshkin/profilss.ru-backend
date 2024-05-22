package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/hr/entity"
)

type HrUsecase struct {
	repository entity.HrRepository
}

func NewHrUsecase(repository entity.HrRepository) *HrUsecase {
	return &HrUsecase{
		repository: repository,
	}
}

func (usecase *HrUsecase) Create(entity *entity.Hr) (*entity.Hr, error) {
	return usecase.repository.Create(entity)
}

func (usecase *HrUsecase) Find() ([]entity.Hr, error) {
	return usecase.repository.Find()
}

func (usecase *HrUsecase) First(id string) (*entity.Hr, error) {
	return usecase.repository.First(id)
}

func (usecase *HrUsecase) Update(entity *entity.Hr, id string) (*entity.Hr, error) {
	return usecase.repository.Update(entity, id)
}

func (usecase *HrUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
