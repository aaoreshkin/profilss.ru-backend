package repository

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/bid/entity"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
)

type BidRepository struct {
	database *database.Database
}

func NewBidRepository(database *database.Database) *BidRepository {
	return &BidRepository{database: database}
}

func (repository *BidRepository) Create(entity *entity.Bid) (*entity.Bid, error) {
	return entity, repository.database.Create(&entity).Error
}

func (repository *BidRepository) Find() ([]entity.Bid, error) {
	entity := []entity.Bid{}

	return entity, repository.database.Find(&entity).Error
}

func (repository *BidRepository) First(id string) (*entity.Bid, error) {
	entity := &entity.Bid{}

	return entity, repository.database.Where("id = ?", id).First(&entity).Error
}

func (repository *BidRepository) Delete(id string) error {
	return repository.database.Where("id = ?", id).Delete(&entity.Bid{}).Error
}
