package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	database *database.Database
}

func NewProductRepository(database *database.Database) *ProductRepository {
	return &ProductRepository{database: database}
}

func (repository *ProductRepository) Create(entity *entity.Product) (*entity.Product, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *ProductRepository) Find() ([]entity.Product, error) {
	entity := []entity.Product{}

	return entity, repository.database.Debug().Preload(clause.Associations).Preload("Characteristics." + clause.Associations).Find(&entity).Error
}

func (repository *ProductRepository) First(id string) (*entity.Product, error) {
	entity := &entity.Product{}

	return entity, repository.database.Where("id = ?", id).Preload(clause.Associations).Preload("Characteristics." + clause.Associations).First(&entity).Error
}

func (repository *ProductRepository) Update(entity *entity.Product, id string) (*entity.Product, error) {
	return entity, repository.database.Session(&gorm.Session{FullSaveAssociations: true}).Model(&entity).Where("id = ?", id).Updates(&entity).Error
}

func (repository *ProductRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Product{}).Error
}
