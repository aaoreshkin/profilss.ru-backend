package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/chat/entity"
)

type ChatUsecase struct {
	repository entity.ChatRepository
}

func NewChatUsecase(repository entity.ChatRepository) *ChatUsecase {
	return &ChatUsecase{
		repository: repository,
	}
}

func (usecase *ChatUsecase) Create(entity *entity.Chat) (*entity.Chat, error) {
	return usecase.repository.Create(entity)
}

func (usecase *ChatUsecase) Find() ([]entity.Chat, error) {
	return usecase.repository.Find()
}

func (usecase *ChatUsecase) First(id string) ([]entity.Chat, error) {
	return usecase.repository.First(id)
}

func (usecase *ChatUsecase) Update(entity *entity.Chat) (*entity.Chat, error) {
	return usecase.repository.Update(entity)
}

func (usecase *ChatUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
