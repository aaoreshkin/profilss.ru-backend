package usecase

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/bid/entity"
)

type BidUsecase struct {
	repository entity.BidRepository
}

func NewBidUsecase(repository entity.BidRepository) *BidUsecase {
	return &BidUsecase{
		repository: repository,
	}
}

func (usecase *BidUsecase) Create(entity *entity.Bid) (*entity.Bid, error) {
	return usecase.repository.Create(entity)
}

func (usecase *BidUsecase) Find() ([]entity.Bid, error) {
	return usecase.repository.Find()
}

func (usecase *BidUsecase) First(id string) (*entity.Bid, error) {
	return usecase.repository.First(id)
}

func (usecase *BidUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
