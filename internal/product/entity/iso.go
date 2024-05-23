package entity

type (
	Iso struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	IsoUsecase interface {
		Create(*Iso) (*Iso, error)
		Find() ([]Iso, error)
		First(string) (*Iso, error)
		Delete(string) error
	}

	IsoRepository interface {
		Create(*Iso) (*Iso, error)
		Find() ([]Iso, error)
		First(string) (*Iso, error)
		Delete(string) error
	}
)
