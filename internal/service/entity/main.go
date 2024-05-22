package entity

import "time"

type (
	Service struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Content     string    `json:"content"`
		Quote       string    `json:"quote"`
		File        string    `json:"file"`
		CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
		Published   bool      `json:"published" gorm:"default:false"`
	}

	ServiceUsecase interface {
		Create(*Service) (*Service, error)
		Find() ([]Service, error)
		First(string) (*Service, error)
		Update(*Service, string) (*Service, error)
		Delete(string) error
	}

	ServiceRepository interface {
		Create(*Service) (*Service, error)
		Find() ([]Service, error)
		First(string) (*Service, error)
		Update(*Service, string) (*Service, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
func (response *Service) NewResponse() *Service {
	return &Service{
		ID:          response.ID,
		Name:        response.Name,
		Description: response.Description,
		Content:     response.Content,
		Quote:       response.Quote,
		File:        response.File,
		CreatedAt:   response.CreatedAt,
		Published:   response.Published,
	}
}
