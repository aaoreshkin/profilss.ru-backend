package entity

type (
	Support struct {
		ManagerID uint64 `json:"manager_id" gorm:"default:null"`
		Message   string `json:"message"`
		SessionID string `json:"session_id" gorm:"type:uuid;"`
		Status    bool   `json:"status" gorm:"default:false"`
	}

	SupportUsecase interface {
		Create(*Support) (*Support, error)
		Find() ([]Support, error)
		First(string) ([]Support, error)
		Update(string) error
		Delete(string) error
	}

	SupportRepository interface {
		Create(*Support) (*Support, error)
		Find() ([]Support, error)
		First(string) ([]Support, error)
		Update(string) error
		Delete(string) error
	}
)
