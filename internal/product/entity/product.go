package entity

import "time"

type (
	Product struct {
		ID               uint64         `json:"id"`
		Category         Category       `json:"category" gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		CategoryID       uint64         `json:"-"`
		Isos             []Iso          `json:"isos" gorm:"many2many:products_isos;constraint:OnDelete:CASCADE;"`
		Characteristic   Characteristic `json:"characteristic" gorm:"foreignKey:CharacteristicID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
		CharacteristicID uint64         `json:"-"`
		CreatedAt        time.Time      `json:"created_at" gorm:"default:now()"`
		Published        bool           `json:"published" gorm:"default:false"`
	}

	ProductUsecase interface {
		Create(*Product) (*Product, error)
		Find() ([]Product, error)
		First(string) (*Product, error)
		Update(*Product, string) (*Product, error)
		Delete(string) error
	}

	ProductRepository interface {
		Create(*Product) (*Product, error)
		Find() ([]Product, error)
		First(string) (*Product, error)
		Update(*Product, string) (*Product, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
