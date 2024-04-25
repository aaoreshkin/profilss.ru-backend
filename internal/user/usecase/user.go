package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

type UserUsecase struct {
	repository entity.UserRepository
}

func NewUserUsecase(repository entity.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (usecase *UserUsecase) Create(entity *entity.User) (*entity.User, error) {

	// Создавать пользователей может только администратор
	// роль назначает тоже администратор

	// Hash entity raw password
	hashedPassword, err := common.HashPassword(entity.Password)
	if err != nil {
		return nil, err
	}

	// Set hashed password
	entity.Password = hashedPassword

	// Hash access token
	hashedToken, err := common.HashToken(entity.Email, entity.PermissionID)
	if err != nil {
		return nil, err
	}

	// Set access token
	entity.AccessToken = hashedToken

	result, err := usecase.repository.Create(entity)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (usecase *UserUsecase) Find() ([]entity.User, error) {
	return usecase.repository.Find()
}
