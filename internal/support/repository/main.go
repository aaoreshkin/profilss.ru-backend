package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/support/entity"
)

type SupportRepository struct {
	database *database.Database
}

func NewSupportRepository(database *database.Database) *SupportRepository {
	return &SupportRepository{database: database}
}

func (repository *SupportRepository) Create(entity *entity.Support) (*entity.Support, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *SupportRepository) Find() ([]entity.Support, error) {
	entity := []entity.Support{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *SupportRepository) First(id string) ([]entity.Support, error) {
	entry := []entity.Support{}

	return entry, repository.database.Where("session_id = ?", id).Find(&entry).Error
}

func (repository *SupportRepository) Update(entity *entity.Support) (*entity.Support, error) {
	return entity, repository.database.Save(&entity).Error
}

func (repository *SupportRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Support{}).Error
}
