package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/service/entity"
)

type ServiceUsecase struct {
	repository entity.ServiceRepository
}

func NewServiceUsecase(repository entity.ServiceRepository) *ServiceUsecase {
	return &ServiceUsecase{
		repository: repository,
	}
}

func (usecase *ServiceUsecase) Create(entity *entity.Service) (*entity.Service, error) {
	return usecase.repository.Create(entity)
}

func (usecase *ServiceUsecase) Find() ([]entity.Service, error) {
	return usecase.repository.Find()
}

func (usecase *ServiceUsecase) First(id string) (*entity.Service, error) {
	return usecase.repository.First(id)
}

func (usecase *ServiceUsecase) Update(entity *entity.Service, id string) (*entity.Service, error) {
	return usecase.repository.Update(entity, id)
}

func (usecase *ServiceUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
