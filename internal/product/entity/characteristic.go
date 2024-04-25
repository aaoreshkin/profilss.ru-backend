package entity

type (
	Characteristic struct {
		ID          uint64 `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	CharacteristicUsecase interface {
		Create(*Characteristic) (*Characteristic, error)
		Find() ([]Characteristic, error)
		First(string) (*Characteristic, error)
		Delete(string) error
	}

	CharacteristicRepository interface {
		Create(*Characteristic) (*Characteristic, error)
		Find() ([]Characteristic, error)
		First(string) (*Characteristic, error)
		Delete(string) error
	}
)
