package entity

import "time"

type (
	SubCategory struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Content     string    `json:"content"`
		Published   bool      `json:"published" gorm:"default:false"`
		CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
	}

	SubCategoryUsecase interface {
		Create(*SubCategory) (*SubCategory, error)
		Find() ([]SubCategory, error)
		First(string) (*SubCategory, error)
		Update(*SubCategory) (*SubCategory, error)
		Delete(string) error
	}

	SubCategoryRepository interface {
		Create(*SubCategory) (*SubCategory, error)
		Find() ([]SubCategory, error)
		First(string) (*SubCategory, error)
		Update(*SubCategory) (*SubCategory, error)
		Delete(string) error
	}
)
