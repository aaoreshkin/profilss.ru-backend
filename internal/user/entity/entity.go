package entity

type (
	User struct {
		ID            string       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
		Email         string       `json:"email" gorm:"unique"`
		Password      string       `json:"password,omitempty"`
		AccessLevel   *AccessLevel `json:"access_level,omitempty"`
		AccessLevelID string       `json:"access_level_id" gorm:"type:uuid;default:null"`
		AccessToken   string       `json:"access_token,omitempty" gorm:"default:null"`
	}

	AccessLevel struct {
		ID    string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
		Title string `json:"title" gorm:"unique"`
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
		ID:            response.ID,
		Email:         response.Email,
		AccessLevelID: response.AccessLevelID,
		AccessToken:   response.AccessToken,
	}
}
