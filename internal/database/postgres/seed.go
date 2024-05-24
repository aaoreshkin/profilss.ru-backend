package postgres

import (
	"github.com/oreshkindev/profilss.ru-backend/common"
	product "github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
	user "github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

type Seed struct {
	*Postgres
}

// new seed
func NewSeed(database *Postgres) (*Seed, error) {
	return &Seed{database}, nil
}

func (s *Seed) Seed() error {

	if err := s.Iso(); err != nil {
		return err
	}

	if err := s.Category(); err != nil {
		return err
	}

	if err := s.Permission(); err != nil {
		return err
	}

	if err := s.User(); err != nil {
		return err
	}

	return nil
}

func (s *Seed) Iso() error {
	entries := []product.Iso{
		{
			Name: "ГОСТ 3262-75",
		},
		{
			Name: "ГОСТ 10704-91",
		},
		{
			Name: "ГОСТ 10705-80",
		},
		{
			Name: "ГОСТ 8639-82",
		},
		{
			Name: "ГОСТ 8645-68",
		},
		{
			Name: "ГОСТ 10707-80",
		},
		{
			Name: "ГОСТ 8644-68 (тип А и тип В)",
		},
		{
			Name: "ГОСТ 8642-68",
		},
		{
			Name: "ГОСТ 34028-16",
		},
		{
			Name: "ГОСТ 2590-2006",
		},
		{
			Name: "ГОСТ 30136-95",
		},
		{
			Name: "ГОСТ 8509-93",
		},
		{
			Name: "ГОСТ 103-2006 МТ",
		},
		{
			Name: "ГОСТ 2591-2006",
		},
		{
			Name: "ГОСТ 16523-97",
		},
	}

	return s.Create(&entries).Error
}

func (s *Seed) Category() error {
	characteristic := []product.Category{
		{
			Name: "Трубы стальные эл/сварные прямошовные",
			File: product.File{
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Трубы стальные квадратные эл/сварные прямошовные",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Трубы стальные круглые эл/сварные прямошовные",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Трубы стальные прямоугольные эл/сварные прямошовные",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Трубы стальные плоскоовальные эл/сварные прямошовные",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Трубы стальные овальные эл/сварные прямошовные",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Трубы стальные арочные эл/сварные прямошовные",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Арматура A III",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Арматура A I круг",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Уголок",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Полоса",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Квадрат",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
		{
			Name: "Лист г/к",
			File: product.File{
				ID:      1,
				Preview: "preview.png",
				Video:   "video.mp4",
			},
		},
	}

	return s.Create(&characteristic).Error
}

func (s *Seed) Permission() error {
	roles := []user.Permission{
		{
			Rule: "Superuser",
		},
		{
			Rule: "Manager",
		},
		{
			Rule: "Guest",
		},
	}

	return s.Create(&roles).Error
}

func (s *Seed) User() error {
	users := []user.User{
		{
			Email:    "oreshkin.dev@outlook.com",
			Password: "pAss1word*",
		},
		{
			Email:    "ioreshkin@outlook.com",
			Password: "pAss1word*",
		},
	}

	var permissionID string

	// Get permissionID by name
	if err := s.Model(&user.Permission{}).Select("id").Where("rule = ?", "Superuser").First(&permissionID).Error; err != nil {
		return err
	}

	for _, user := range users {

		// Hash raw password
		hashedPassword, err := common.HashPassword(user.Password)
		if err != nil {
			return err
		}

		// Hash access token
		accessToken, err := common.HashToken(user.Email, permissionID)
		if err != nil {
			return err
		}

		user.AccessToken = accessToken
		user.Password = hashedPassword
		user.PermissionID = permissionID

		if err := s.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
