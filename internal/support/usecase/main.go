package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/support/entity"
)

type SupportUsecase struct {
	repository entity.SupportRepository
}

func NewSupportUsecase(repository entity.SupportRepository) *SupportUsecase {
	return &SupportUsecase{
		repository: repository,
	}
}

func (usecase *SupportUsecase) Create(entity *entity.Support) (*entity.Support, error) {
	return usecase.repository.Create(entity)
}

func (usecase *SupportUsecase) Find() ([]entity.Support, error) {
	return usecase.repository.Find()
}

func (usecase *SupportUsecase) First(id string) ([]entity.Support, error) {
	return usecase.repository.First(id)
}

func (usecase *SupportUsecase) Update(id string) error {
	return usecase.repository.Update(id)
}

func (usecase *SupportUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
