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

func (usecase *UserUsecase) Post(entity *entity.User) (*entity.User, error) {

	// Hash entity raw password
	hashedPassword, err := common.HashPassword(entity.Password)
	if err != nil {
		return nil, err
	}

	// Set hashed password
	entity.Password = hashedPassword

	result, err := usecase.repository.Post(entity)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (usecase *UserUsecase) Get() ([]entity.User, error) {
	result, err := usecase.repository.Get()
	if err != nil || result == nil {
		return nil, err
	}

	return result, nil
}
