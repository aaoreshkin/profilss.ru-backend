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

	MeasureRepository repository.MeasureRepository
	MeasureUsecase    usecase.MeasureUsecase
	MeasureController controller.MeasureController

	CharacteristicRepository repository.CharacteristicRepository
	CharacteristicUsecase    usecase.CharacteristicUsecase
	CharacteristicController controller.CharacteristicController
}

func NewManager(database *database.Database) *Manager {
	productRepository := repository.NewProductRepository(database)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productController := controller.NewProductController(productUsecase)

	measureRepository := repository.NewMeasureRepository(database)
	measureUsecase := usecase.NewMeasureUsecase(measureRepository)
	measureController := controller.NewMeasureController(measureUsecase)

	characteristicRepository := repository.NewCharacteristicRepository(database)
	characteristicUsecase := usecase.NewCharacteristicUsecase(characteristicRepository)
	characteristicController := controller.NewCharacteristicController(characteristicUsecase)

	return &Manager{
		ProductRepository: *productRepository,
		ProductUsecase:    *productUsecase,
		ProductController: *productController,

		MeasureRepository: *measureRepository,
		MeasureUsecase:    *measureUsecase,
		MeasureController: *measureController,

		CharacteristicRepository: *characteristicRepository,
		CharacteristicUsecase:    *characteristicUsecase,
		CharacteristicController: *characteristicController,
	}
}
