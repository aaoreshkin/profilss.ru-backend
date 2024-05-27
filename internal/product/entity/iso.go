package entity

type (
	Iso struct {
		ID   uint64 `json:"id"`
		Name string `json:"name" gorm:"unique"`
	}

	IsoUsecase interface {
		Create(*Iso) (*Iso, error)
		Find() ([]Iso, error)
		First(string) (*Iso, error)
		Update(*Iso) (*Iso, error)
		Delete(string) error
	}

	IsoRepository interface {
		Create(*Iso) (*Iso, error)
		Find() ([]Iso, error)
		First(string) (*Iso, error)
		Update(*Iso) (*Iso, error)
		Delete(string) error
	}
)
