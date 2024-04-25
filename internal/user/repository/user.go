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

func (repository *UserRepository) First(id string) (*entity.User, error) {
	entity := &entity.User{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *UserRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.User{}).Error
}
