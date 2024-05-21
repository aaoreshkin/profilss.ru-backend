package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type CategoryUsecase struct {
	repository entity.CategoryRepository
}

func NewCategoryUsecase(repository entity.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		repository: repository,
	}
}

func (usecase *CategoryUsecase) Create(entity *entity.Category) (*entity.Category, error) {
	return usecase.repository.Create(entity)
}

func (usecase *CategoryUsecase) Find() ([]entity.Category, error) {
	return usecase.repository.Find()
}

func (usecase *CategoryUsecase) First(id string) (*entity.Category, error) {
	return usecase.repository.First(id)
}

func (usecase *CategoryUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
