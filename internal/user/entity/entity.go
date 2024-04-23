package entity

type User struct {
	ID            string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Email         string `json:"email" gorm:"unique"`
	Password      string `json:"password,omitempty"`
	AccessLevelID string `json:"access_level_id" gorm:"type:uuid"`
	AccessToken   string `json:"access_token" gorm:"default:null"`
}

type AccessLevel struct {
	ID    string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Title string `json:"title" gorm:"unique"`
}

type (
	UserUsecase interface {
		Get() ([]User, error)
	}

	UserRepository interface {
		Get() ([]User, error)
	}
)
