package entity

import "time"

type (
	User struct {
		ID           string      `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
		AccessToken  string      `json:"access_token,omitempty" gorm:"default:null"`
		Email        string      `json:"email" gorm:"unique"`
		Password     string      `json:"password,omitempty"`
		Permission   *Permission `json:"permission,omitempty"`
		PermissionID string      `json:"permission_id" gorm:"type:uuid;default:null"`
		Fullname     string      `json:"fullname,omitempty"`
		Phone        string      `json:"phone,omitempty"`
		UpdatedAt    time.Time   `json:"updated_at" gorm:"default:now()"`
	}

	UserUsecase interface {
		Create(*User) (*User, error)
		Find() ([]User, error)
		First(string) (*User, error)
		Update(*User, string) (*User, error)
		Delete(string) error
		Verify(*User) (*User, error)
	}

	UserRepository interface {
		Create(*User) (*User, error)
		Find() ([]User, error)
		First(string) (*User, error)
		Delete(string) error
		Update(*User, string) (*User, error)
		FindManager() (*string, error)
	}
)

// fields of struct that will be returned
func (response *User) NewResponse() *User {
	return &User{
		ID:           response.ID,
		AccessToken:  response.AccessToken,
		Email:        response.Email,
		Fullname:     response.Fullname,
		Phone:        response.Phone,
		PermissionID: response.PermissionID,
		UpdatedAt:    response.UpdatedAt,
	}
}
