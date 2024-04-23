package usecase

import "github.com/oreshkindev/profilss.ru-backend/internal/user/entity"

type UserUsecase struct {
	repository entity.UserRepository
}

func NewUserUsecase(repository entity.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (usecase *UserUsecase) Get() ([]entity.User, error) {
	result, err := usecase.repository.Get()
	if err != nil || result == nil {
		return nil, err
	}

	return result, nil
}
