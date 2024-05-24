package entity

import "time"

type (
	Category struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Content     string    `json:"content"`
		Adv         string    `json:"adv"`
		File        File      `json:"file" gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		FileID      uint64    `json:"-"`
		Published   bool      `json:"published" gorm:"default:false"`
		CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
	}

	File struct {
		ID      uint64 `json:"id"`
		Preview string `json:"preview"`
		Video   string `json:"video"`
	}

	CategoryUsecase interface {
		Create(*Category) (*Category, error)
		Find() ([]Category, error)
		First(string) (*Category, error)
		Update(*Category, string) (*Category, error)
		Delete(string) error
	}

	CategoryRepository interface {
		Create(*Category) (*Category, error)
		Find() ([]Category, error)
		First(string) (*Category, error)
		Update(*Category, string) (*Category, error)
		Delete(string) error
	}
)
