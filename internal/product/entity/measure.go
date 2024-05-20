package entity

type (
	Measure struct {
		ID   uint64 `json:"id"`
		Code string `json:"code" gorm:"unique"`
		Name string `json:"name"`
	}

	MeasureUsecase interface {
		Create(*Measure) (*Measure, error)
		Find() ([]Measure, error)
		First(string) (*Measure, error)
		Delete(string) error
	}

	MeasureRepository interface {
		Create(*Measure) (*Measure, error)
		Find() ([]Measure, error)
		First(string) (*Measure, error)
		Delete(string) error
	}
)
