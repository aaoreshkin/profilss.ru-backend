package entity

import (
	"time"

	"github.com/lib/pq"
)

type (
	Hr struct {
		ID           uint64         `json:"id"`
		Name         string         `json:"name"`
		Period       string         `json:"period"`
		PeriodID     uint64         `json:"-"`
		Cost         string         `json:"cost"`
		Address      string         `json:"address"`
		Duties       pq.StringArray `json:"duties" gorm:"type:text[]"`
		Requirements pq.StringArray `json:"requirements" gorm:"type:text[]"`
		Conditions   pq.StringArray `json:"conditions" gorm:"type:text[]"`
		CreatedAt    time.Time      `json:"created_at" gorm:"default:now()"`
		Published    bool           `json:"published" gorm:"default:false"`
	}

	HrUsecase interface {
		Create(*Hr) (*Hr, error)
		Find() ([]Hr, error)
		First(string) (*Hr, error)
		Update(*Hr, string) (*Hr, error)
		Delete(string) error
	}

	HrRepository interface {
		Create(*Hr) (*Hr, error)
		Find() ([]Hr, error)
		First(string) (*Hr, error)
		Update(*Hr, string) (*Hr, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
func (response *Hr) NewResponse() *Hr {
	return &Hr{
		ID:           response.ID,
		Name:         response.Name,
		Period:       response.Period,
		Cost:         response.Cost,
		Address:      response.Address,
		Duties:       response.Duties,
		Requirements: response.Requirements,
		Conditions:   response.Conditions,
		CreatedAt:    response.CreatedAt,
		Published:    response.Published,
	}
}
