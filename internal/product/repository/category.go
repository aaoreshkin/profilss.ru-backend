package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository struct {
	database *database.Database
}

func NewCategoryRepository(database *database.Database) *CategoryRepository {
	return &CategoryRepository{database: database}
}

func (repository *CategoryRepository) Create(entry *entity.Category) (*entity.Category, error) {
	return entry, repository.database.Session(&gorm.Session{FullSaveAssociations: true}).Create(&entry).Error
}

func (repository *CategoryRepository) Find() ([]entity.Category, error) {
	entry := []entity.Category{}

	return entry, repository.database.Preload(clause.Associations).Find(&entry).Error
}

func (repository *CategoryRepository) First(id string) (*entity.Category, error) {
	entry := &entity.Category{}

	return entry, repository.database.Preload(clause.Associations).Where("id = ?", id).First(&entry).Error
}

func (repository *CategoryRepository) Update(entry *entity.Category) (*entity.Category, error) {
	if r := repository.database.Model(&entity.Category{ID: entry.ID}).Association("Iso").Clear(); r != nil {
		return nil, r
	}

	if r := repository.database.Model(&entity.Category{ID: entry.ID}).Association("File").Clear(); r != nil {
		return nil, r
	}

	return entry, repository.database.Session(&gorm.Session{FullSaveAssociations: true}).Save(&entry).Error
}

func (repository *CategoryRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Category{}).Error
}
