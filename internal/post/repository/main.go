package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/post/entity"
)

type PostRepository struct {
	database *database.Database
}

func NewPostRepository(database *database.Database) *PostRepository {
	return &PostRepository{database: database}
}

func (repository *PostRepository) Create(entity *entity.Post) (*entity.Post, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *PostRepository) Find() ([]entity.Post, error) {
	entity := []entity.Post{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *PostRepository) First(id string) (*entity.Post, error) {
	entity := &entity.Post{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *PostRepository) Update(entity *entity.Post, id string) (*entity.Post, error) {
	return entity, repository.database.Save(&entity).Error
}

func (repository *PostRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Post{}).Error
}
