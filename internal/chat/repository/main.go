package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/chat/entity"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
)

type ChatRepository struct {
	database *database.Database
}

func NewChatRepository(database *database.Database) *ChatRepository {
	return &ChatRepository{database: database}
}

func (repository *ChatRepository) Create(entity *entity.Chat) (*entity.Chat, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *ChatRepository) Find() ([]entity.Chat, error) {
	entity := []entity.Chat{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *ChatRepository) First(id string) ([]entity.Chat, error) {
	entry := []entity.Chat{}

	return entry, repository.database.Where("session_id = ?", id).Find(&entry).Error
}

func (repository *ChatRepository) Update(entity *entity.Chat) (*entity.Chat, error) {
	return entity, repository.database.Save(&entity).Error
}

func (repository *ChatRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Chat{}).Error
}
