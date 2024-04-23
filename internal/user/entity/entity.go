package entity

type User struct {
	ID    uint64 `json:"id"`
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
