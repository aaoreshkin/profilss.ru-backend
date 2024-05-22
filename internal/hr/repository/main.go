package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/hr/entity"
)

type HrRepository struct {
	database *database.Database
}

func NewHrRepository(database *database.Database) *HrRepository {
	return &HrRepository{database: database}
}

func (repository *HrRepository) Create(entity *entity.Hr) (*entity.Hr, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *HrRepository) Find() ([]entity.Hr, error) {
	entity := []entity.Hr{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *HrRepository) First(id string) (*entity.Hr, error) {
	entity := &entity.Hr{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *HrRepository) Update(entity *entity.Hr, id string) (*entity.Hr, error) {
	return entity, repository.database.Save(&entity).Error
}

func (repository *HrRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Hr{}).Error
}
