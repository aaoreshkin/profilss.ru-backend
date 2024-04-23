package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(database *database.Database) *UserRepository {
	return &UserRepository{database: database}
}

func (repository *UserRepository) Get() ([]entity.User, error) {
	entity := []entity.User{}

	if r := repository.database.Preload(clause.Associations).Find(&entity); r.Error != nil {
		return nil, r.Error
	}

	return entity, nil
}
