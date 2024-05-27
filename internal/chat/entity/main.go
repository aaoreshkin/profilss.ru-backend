package entity

type (
	Chat struct {
		ManagerID uint64 `json:"manager_id" gorm:"default:null"`
		Message   string `json:"message"`
		SessionID string `json:"session_id" gorm:"type:uuid;"`
		Status    bool   `json:"status" gorm:"default:false"`
	}

	ChatUsecase interface {
		Create(*Chat) (*Chat, error)
		Find() ([]Chat, error)
		First(string) ([]Chat, error)
		Update(*Chat) (*Chat, error)
		Delete(string) error
	}

	ChatRepository interface {
		Create(*Chat) (*Chat, error)
		Find() ([]Chat, error)
		First(string) ([]Chat, error)
		Update(*Chat) (*Chat, error)
		Delete(string) error
	}
)
