package entity

// Rules: F - forbidden, R - read, W - write

type (
	Permission struct {
		ID   string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
		Rule string `json:"rule" gorm:"unique"`
	}

	PermissionUsecase interface {
		Get(string) (*Permission, error)
	}

	PermissionRepository interface {
		Get(string) (*Permission, error)
	}
)

// fields of struct that will be returned
func (response *Permission) NewResponse() *Permission {
	return &Permission{
		ID:   response.ID,
		Rule: response.Rule,
	}
}
