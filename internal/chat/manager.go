package chat

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/chat/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/chat/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/chat/usecase"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
)

type Manager struct {
	ChatRepository repository.ChatRepository
	ChatUsecase    usecase.ChatUsecase
	ChatController controller.ChatController
}

func NewManager(database *database.Database) *Manager {
	chatRepository := repository.NewChatRepository(database)
	chatUsecase := usecase.NewChatUsecase(chatRepository)
	chatController := controller.NewChatController(chatUsecase)

	return &Manager{
		ChatRepository: *chatRepository,
		ChatUsecase:    *chatUsecase,
		ChatController: *chatController,
	}
}
