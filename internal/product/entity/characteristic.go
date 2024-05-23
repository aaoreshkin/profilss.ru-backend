package entity

type (
	Characteristic struct {
		ID        uint64 `json:"id"`
		MaxPrice  string `json:"max_price"`
		Price     string `json:"price"`
		Size      string `json:"size"`
		Thickness string `json:"thickness"`
		Weight    string `json:"weight"`
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
