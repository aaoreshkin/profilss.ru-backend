package entity

// Rules: S - superuser, M - manager, G - guest

type (
	Permission struct {
		ID   string `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
		Rule string `json:"rule" gorm:"unique"`
	}

	PermissionUsecase interface {
		First(string) (*Permission, error)
	}

	PermissionRepository interface {
		First(string) (*Permission, error)
	}
)

// fields of struct that will be returned
func (response *Permission) NewResponse() *Permission {
	return &Permission{
		ID:   response.ID,
		Rule: response.Rule,
	}
}
