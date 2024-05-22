package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(database *database.Database) *UserRepository {
	return &UserRepository{database: database}
}

func (repository *UserRepository) Create(entity *entity.User) (*entity.User, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *UserRepository) Find() ([]entity.User, error) {
	entity := []entity.User{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *UserRepository) First(email string) (*entity.User, error) {
	entity := &entity.User{}

	return entity, repository.database.Where("email = ?", email).First(entity).Error
}

func (repository *UserRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.User{}).Error
}

func (repository *UserRepository) Update(entity *entity.User, id string) (*entity.User, error) {
	return entity, repository.database.Where("id = ?", id).Updates(&entity).Error
}

func (repository *UserRepository) FindManager() (*string, error) {
	var permissionID string

	// Get permissionID by name
	if err := repository.database.Model(&entity.Permission{}).Select("id").Where("rule = ?", "Manager").First(&permissionID).Error; err != nil {
		return nil, err
	}

	return &permissionID, nil
}
