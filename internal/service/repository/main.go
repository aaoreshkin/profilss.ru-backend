package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/service/entity"
)

type ServiceRepository struct {
	database *database.Database
}

func NewServiceRepository(database *database.Database) *ServiceRepository {
	return &ServiceRepository{database: database}
}

func (repository *ServiceRepository) Create(entity *entity.Service) (*entity.Service, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *ServiceRepository) Find() ([]entity.Service, error) {
	entity := []entity.Service{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *ServiceRepository) First(id string) (*entity.Service, error) {
	entity := &entity.Service{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *ServiceRepository) Update(entity *entity.Service, id string) (*entity.Service, error) {
	return entity, repository.database.Save(&entity).Error
}

func (repository *ServiceRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Service{}).Error
}
