package product

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/usecase"
)

type Manager struct {
	ProductRepository repository.ProductRepository
	ProductUsecase    usecase.ProductUsecase
	ProductController controller.ProductController

	IsoRepository repository.IsoRepository
	IsoUsecase    usecase.IsoUsecase
	IsoController controller.IsoController

	CharacteristicRepository repository.CharacteristicRepository
	CharacteristicUsecase    usecase.CharacteristicUsecase
	CharacteristicController controller.CharacteristicController

	CategoryRepository repository.CategoryRepository
	CategoryUsecase    usecase.CategoryUsecase
	CategoryController controller.CategoryController
}

func NewManager(database *database.Database) *Manager {
	productRepository := repository.NewProductRepository(database)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productController := controller.NewProductController(productUsecase)

	isoRepository := repository.NewIsoRepository(database)
	isoUsecase := usecase.NewIsoUsecase(isoRepository)
	isoController := controller.NewIsoController(isoUsecase)

	characteristicRepository := repository.NewCharacteristicRepository(database)
	characteristicUsecase := usecase.NewCharacteristicUsecase(characteristicRepository)
	characteristicController := controller.NewCharacteristicController(characteristicUsecase)

	categoryRepository := repository.NewCategoryRepository(database)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUsecase)

	return &Manager{
		ProductRepository: *productRepository,
		ProductUsecase:    *productUsecase,
		ProductController: *productController,

		IsoRepository: *isoRepository,
		IsoUsecase:    *isoUsecase,
		IsoController: *isoController,

		CharacteristicRepository: *characteristicRepository,
		CharacteristicUsecase:    *characteristicUsecase,
		CharacteristicController: *characteristicController,

		CategoryRepository: *categoryRepository,
		CategoryUsecase:    *categoryUsecase,
		CategoryController: *categoryController,
	}
}
