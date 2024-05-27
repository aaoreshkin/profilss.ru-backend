package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubCategoryRepository struct {
	database *database.Database
}

func NewSubCategoryRepository(database *database.Database) *SubCategoryRepository {
	return &SubCategoryRepository{database: database}
}

func (repository *SubCategoryRepository) Create(entry *entity.SubCategory) (*entity.SubCategory, error) {
	return entry, repository.database.Create(&entry).Error
}

func (repository *SubCategoryRepository) Find() ([]entity.SubCategory, error) {
	entry := []entity.SubCategory{}

	return entry, repository.database.Preload(clause.Associations).Find(&entry).Error
}

func (repository *SubCategoryRepository) First(id string) (*entity.SubCategory, error) {
	entry := &entity.SubCategory{}

	return entry, repository.database.Preload(clause.Associations).Where("id = ?", id).First(&entry).Error
}

func (repository *SubCategoryRepository) Update(entry *entity.SubCategory) (*entity.SubCategory, error) {
	return entry, repository.database.Session(&gorm.Session{FullSaveAssociations: true}).Save(&entry).Error
}

func (repository *SubCategoryRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.SubCategory{}).Error
}
