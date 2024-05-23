package entity

import "time"

type (
	Category struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Content     string    `json:"content"`
		Adv         string    `json:"adv"`
		File        string    `json:"file"`
		Published   bool      `json:"published" gorm:"default:false"`
		CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
	}

	CategoryUsecase interface {
		Create(*Category) (*Category, error)
		Find() ([]Category, error)
		First(string) (*Category, error)
		Delete(string) error
	}

	CategoryRepository interface {
		Create(*Category) (*Category, error)
		Find() ([]Category, error)
		First(string) (*Category, error)
		Delete(string) error
	}
)
