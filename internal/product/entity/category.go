package entity

import "time"

type (
	Category struct {
		ID            uint64      `json:"id"`
		File          File        `json:"file" gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		FileID        uint64      `json:"-"`
		SubCategory   SubCategory `json:"sub_category" gorm:"foreignKey:SubCategoryID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		SubCategoryID uint64      `json:"-"`
		Iso           []Iso       `json:"iso" gorm:"many2many:categories_isos;constraint:OnDelete:CASCADE;"`
		Name          string      `json:"name"`
		Description   string      `json:"description"`
		Published     bool        `json:"published" gorm:"default:false"`
		CreatedAt     time.Time   `json:"created_at" gorm:"default:now()"`
	}

	File struct {
		ID    uint64 `json:"id"`
		Image string `json:"image"`
		Video string `json:"video"`
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
