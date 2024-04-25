package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type CharacteristicRepository struct {
	database *database.Database
}

func NewCharacteristicRepository(database *database.Database) *CharacteristicRepository {
	return &CharacteristicRepository{database: database}
}

func (repository *CharacteristicRepository) Create(entity *entity.Characteristic) (*entity.Characteristic, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *CharacteristicRepository) Find() ([]entity.Characteristic, error) {
	entity := []entity.Characteristic{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *CharacteristicRepository) First(id string) (*entity.Characteristic, error) {
	entity := &entity.Characteristic{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *CharacteristicRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Characteristic{}).Error
}
