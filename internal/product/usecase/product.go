package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type ProductUsecase struct {
	repository entity.ProductRepository
}

func NewProductUsecase(repository entity.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repository: repository,
	}
}

func (usecase *ProductUsecase) Create(entity *entity.Product) (*entity.Product, error) {
	return usecase.repository.Create(entity)
}

func (usecase *ProductUsecase) Find() ([]entity.Product, error) {
	return usecase.repository.Find()
}

func (usecase *ProductUsecase) First(id string) (*entity.Product, error) {
	return usecase.repository.First(id)
}

func (usecase *ProductUsecase) Update(entity *entity.Product, id string) (*entity.Product, error) {
	return usecase.repository.Update(entity, id)
}

func (usecase *ProductUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
