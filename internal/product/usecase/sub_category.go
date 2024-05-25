package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type SubCategoryUsecase struct {
	repository entity.SubCategoryRepository
}

func NewSubCategoryUsecase(repository entity.SubCategoryRepository) *SubCategoryUsecase {
	return &SubCategoryUsecase{
		repository: repository,
	}
}

func (usecase *SubCategoryUsecase) Create(entity *entity.SubCategory) (*entity.SubCategory, error) {
	return usecase.repository.Create(entity)
}

func (usecase *SubCategoryUsecase) Find() ([]entity.SubCategory, error) {
	return usecase.repository.Find()
}

func (usecase *SubCategoryUsecase) First(id string) (*entity.SubCategory, error) {
	return usecase.repository.First(id)
}

func (usecase *SubCategoryUsecase) Update(entity *entity.SubCategory, id string) (*entity.SubCategory, error) {
	return usecase.repository.Update(entity, id)
}

func (usecase *SubCategoryUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
