package entity

type (
	User struct {
		ID           string      `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
		Email        string      `json:"email" gorm:"unique"`
		Password     string      `json:"password,omitempty"`
		Permission   *Permission `json:"permission,omitempty"`
		PermissionID string      `json:"permission_id" gorm:"type:uuid;default:null"`
		AccessToken  string      `json:"access_token,omitempty" gorm:"default:null"`
	}

	UserUsecase interface {
		Post(*User) (*User, error)
		Get() ([]User, error)
	}

	UserRepository interface {
		Post(*User) (*User, error)
		Get() ([]User, error)
	}
)

// fields of struct that will be returned
func (response *User) NewResponse() *User {
	return &User{
		ID:           response.ID,
		Email:        response.Email,
		PermissionID: response.PermissionID,
		AccessToken:  response.AccessToken,
	}
}
