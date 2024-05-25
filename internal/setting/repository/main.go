package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/setting/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SettingRepository struct {
	database *database.Database
}

func NewSettingRepository(database *database.Database) *SettingRepository {
	return &SettingRepository{database: database}
}

func (repository *SettingRepository) Create(entity *entity.Setting) (*entity.Setting, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *SettingRepository) Find() ([]entity.Setting, error) {
	entity := []entity.Setting{}

	return entity, repository.database.Preload(clause.Associations).Find(&entity).Error
}

func (repository *SettingRepository) First(id string) (*entity.Setting, error) {
	entity := &entity.Setting{}

	return entity, repository.database.Where("id = ?", id).Preload(clause.Associations).Preload("Employees." + clause.Associations).First(&entity).Error
}

func (repository *SettingRepository) Update(entry *entity.Setting, id string) (*entity.Setting, error) {

	if err := repository.database.Model(&entity.Setting{ID: entry.ID}).Association("Contacts").Clear(); err != nil {
		return nil, err
	}

	if err := repository.database.Model(&entity.Employee{ID: entry.ID}).Association("Contact").Clear(); err != nil {
		return nil, err
	}

	if err := repository.database.Model(&entity.Setting{ID: entry.ID}).Association("Media").Clear(); err != nil {
		return nil, err
	}

	return entry, repository.database.Session(&gorm.Session{FullSaveAssociations: true}).Model(&entry).Where("id = ?", id).Updates(&entry).Error
}

func (repository *SettingRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Setting{}).Error
}
