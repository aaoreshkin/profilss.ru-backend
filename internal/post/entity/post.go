package entity

import "time"

type (
	Post struct {
		ID          uint64    `json:"id"`
		CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
		Description string    `json:"description"`
		Promo       bool      `json:"promo" gorm:"default:false"`
		Published   bool      `json:"published" gorm:"default:false"`
		Title       string    `json:"title"`
	}

	PostUsecase interface {
		Create(*Post) (*Post, error)
		Find() ([]Post, error)
		First(string) (*Post, error)
		Delete(string) error
	}

	PostRepository interface {
		Create(*Post) (*Post, error)
		Find() ([]Post, error)
		First(string) (*Post, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
func (response *Post) NewResponse() *Post {
	return &Post{
		ID:          response.ID,
		CreatedAt:   response.CreatedAt,
		Description: response.Description,
		Promo:       response.Promo,
		Published:   response.Published,
		Title:       response.Title,
	}
}
