package bid

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/bid/controller"
	"github.com/oreshkindev/profilss.ru-backend/internal/bid/repository"
	"github.com/oreshkindev/profilss.ru-backend/internal/bid/usecase"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
)

type Manager struct {
	BidRepository repository.BidRepository
	BidUsecase    usecase.BidUsecase
	BidController controller.BidController
}

func NewManager(database *database.Database) *Manager {
	bidRepository := repository.NewBidRepository(database)
	bidUsecase := usecase.NewBidUsecase(bidRepository)
	bidController := controller.NewBidController(bidUsecase)

	return &Manager{
		BidRepository: *bidRepository,
		BidUsecase:    *bidUsecase,
		BidController: *bidController,
	}
}
