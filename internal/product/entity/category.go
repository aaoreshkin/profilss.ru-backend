package entity

type (
	Category struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	CategoryUsecase interface {
		Create(*Category) (*Category, error)
		Find() ([]Category, error)
		First(string) (*Category, error)
		Delete(string) error
	}

	CategoryRepository interface {
		Create(*Category) (*Category, error)
		Find() ([]Category, error)
		First(string) (*Category, error)
		Delete(string) error
	}
)
