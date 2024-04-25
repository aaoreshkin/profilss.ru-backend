package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type MeasureRepository struct {
	database *database.Database
}

func NewMeasureRepository(database *database.Database) *MeasureRepository {
	return &MeasureRepository{database: database}
}

func (repository *MeasureRepository) Create(entity *entity.Measure) (*entity.Measure, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *MeasureRepository) Find() ([]entity.Measure, error) {
	entity := []entity.Measure{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *MeasureRepository) First(id string) (*entity.Measure, error) {
	entity := &entity.Measure{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *MeasureRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Measure{}).Error
}
