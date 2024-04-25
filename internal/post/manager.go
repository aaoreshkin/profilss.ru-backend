package post

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/post/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/post/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/post/usecase"
)

type Manager struct {
	PostRepository repository.PostRepository
	PostUsecase    usecase.PostUsecase
	PostController controller.PostController
}

func NewManager(database *database.Database) *Manager {
	postRepository := repository.NewPostRepository(database)
	postUsecase := usecase.NewPostUsecase(postRepository)
	postController := controller.NewPostController(postUsecase)

	return &Manager{
		PostRepository: *postRepository,
		PostUsecase:    *postUsecase,
		PostController: *postController,
	}
}
