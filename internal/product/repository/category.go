package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type CategoryRepository struct {
	database *database.Database
}

func NewCategoryRepository(database *database.Database) *CategoryRepository {
	return &CategoryRepository{database: database}
}

func (repository *CategoryRepository) Create(entity *entity.Category) (*entity.Category, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *CategoryRepository) Find() ([]entity.Category, error) {
	entity := []entity.Category{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *CategoryRepository) First(id string) (*entity.Category, error) {
	entity := &entity.Category{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *CategoryRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Category{}).Error
}
