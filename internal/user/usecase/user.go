package usecase

import (
	"fmt"

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

func (usecase *UserUsecase) First(email string) (*entity.User, error) {
	return usecase.repository.First(email)
}

func (usecase *UserUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}

func (usecase *UserUsecase) Verify(entity *entity.User) (*entity.User, error) {

	// Check is user exist
	exist, err := usecase.repository.First(entity.Email)
	if err != nil {
		return nil, err
	}

	// Check is password equal to exist
	equal := common.CheckPasswordHash(entity.Password, exist.Password)
	if !equal {
		return nil, fmt.Errorf("invalid password")
	}

	// Hash access token
	hashedToken, err := common.HashToken(exist.Email, exist.PermissionID)
	if err != nil {
		return nil, err
	}

	// Set access token
	exist.AccessToken = hashedToken

	result, err := usecase.repository.Update(exist)
	if err != nil {
		return nil, err
	}

	return result, nil
}
