package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type IsoRepository struct {
	database *database.Database
}

func NewIsoRepository(database *database.Database) *IsoRepository {
	return &IsoRepository{database: database}
}

func (repository *IsoRepository) Create(entity *entity.Iso) (*entity.Iso, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *IsoRepository) Find() ([]entity.Iso, error) {
	entity := []entity.Iso{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *IsoRepository) First(id string) (*entity.Iso, error) {
	entity := &entity.Iso{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *IsoRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Iso{}).Error
}
