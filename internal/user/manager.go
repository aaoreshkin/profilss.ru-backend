package user

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/usecase"
)

type Manager struct {
	UserRepository repository.UserRepository
	UserUsecase    usecase.UserUsecase
	UserController controller.UserController

	PermissionRepository repository.PermissionRepository
	PermissionUsecase    usecase.PermissionUsecase
	PermissionController controller.PermissionController
}

func NewManager(database *database.Database) *Manager {
	userRepository := repository.NewUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	permissionRepository := repository.NewPermissionRepository(database)
	permissionUsecase := usecase.NewPermissionUsecase(permissionRepository)
	permissionController := controller.NewPermissionController(permissionUsecase)

	return &Manager{
		UserRepository: *userRepository,
		UserUsecase:    *userUsecase,
		UserController: *userController,

		PermissionRepository: *permissionRepository,
		PermissionUsecase:    *permissionUsecase,
		PermissionController: *permissionController,
	}
}
