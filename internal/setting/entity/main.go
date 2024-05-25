package entity

type (
	Setting struct {
		ID             uint64     `json:"id"`
		Name           string     `json:"name"`
		Abbreviation   string     `json:"abbreviation"`
		Contacts       []Contact  `json:"contacts" gorm:"many2many:setting_contacts;constraint:OnDelete:CASCADE;"`
		INN            string     `json:"inn"`
		KPP            string     `json:"kpp"`
		OGRN           string     `json:"ogrn"`
		Employees      []Employee `json:"employees" gorm:"many2many:setting_employees;constraint:OnDelete:CASCADE;"`
		Bank           string     `json:"bank"`
		BIK            string     `json:"bik"`
		KOR            string     `json:"kor"`
		CurrentAccount string     `json:"current_account"`
		OKVD           string     `json:"okvd"`
		Tax            string     `json:"tax"`
		Media          []Media    `json:"media" gorm:"many2many:setting_media;constraint:OnDelete:CASCADE;"`
	}

	Employee struct {
		ID          uint64    `json:"id"`
		Contact     []Contact `json:"contacts,omitempty" gorm:"many2many:employees_contacts;constraint:OnDelete:CASCADE;"`
		Description string    `json:"description"`
		Name        string    `json:"name"`
	}

	Contact struct {
		ID      uint64 `json:"id"`
		Address string `json:"address"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	Media struct {
		ID   uint64 `json:"id"`
		Href string `json:"href"`
		Name string `json:"name"`
	}

	SettingUsecase interface {
		Create(*Setting) (*Setting, error)
		Find() ([]Setting, error)
		First(string) (*Setting, error)
		Update(*Setting, string) (*Setting, error)
		Delete(string) error
	}

	SettingRepository interface {
		Create(*Setting) (*Setting, error)
		Find() ([]Setting, error)
		First(string) (*Setting, error)
		Update(*Setting, string) (*Setting, error)
		Delete(string) error
	}
)

// fields of struct that will be returned
func (response *Setting) NewResponse() *Setting {
	return &Setting{
		ID:             response.ID,
		Name:           response.Name,
		Abbreviation:   response.Abbreviation,
		Contacts:       response.Contacts,
		INN:            response.INN,
		KPP:            response.KPP,
		OGRN:           response.OGRN,
		Employees:      response.Employees,
		Bank:           response.Bank,
		BIK:            response.BIK,
		KOR:            response.KOR,
		CurrentAccount: response.CurrentAccount,
		OKVD:           response.OKVD,
		Tax:            response.Tax,
		Media:          response.Media,
	}
}
