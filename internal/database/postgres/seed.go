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

func (s *Seed) Category() error {
	category := []product.Category{
		{
			Name: "Круглые",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   1,
				Name: "Трубы стальные эл/сварные прямошовные",
			},
			Iso: []product.Iso{
				{
					Name: "ГОСТ 10704-91",
				},
				{
					Name: "ГОСТ 10707-80",
				},
			},
		},
		{
			Name: "Квадратные",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   1,
				Name: "Трубы стальные эл/сварные прямошовные",
			},
			Iso: []product.Iso{
				{
					Name: "ГОСТ 8639-82",
				},
			},
		},
		{
			Name: "Прямоугольные",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   1,
				Name: "Трубы стальные эл/сварные прямошовные",
			},
			Iso: []product.Iso{
				{
					Name: "ГОСТ 8645-68",
				},
			},
		},
		{
			Name: "Плоскоовальные",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   1,
				Name: "Трубы стальные эл/сварные прямошовные",
			},
			Iso: []product.Iso{
				{
					Name: "ГОСТ 8644-68 (тип А и тип В)",
				},
				{
					Name: "ГОСТ 8642-68",
				},
			},
		},
		{
			Name: "Овальные",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   1,
				Name: "Трубы стальные эл/сварные прямошовные",
			},
		},
		{
			Name: "Арочные",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   1,
				Name: "Трубы стальные эл/сварные прямошовные",
			},
		},
		{
			Name: "Арматура A III",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   2,
				Name: "Металлопрокат",
			},
		},
		{
			Name: "Круг (Арматура A I)",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   2,
				Name: "Металлопрокат",
			},
		},
		{
			Name: "Уголок",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   2,
				Name: "Металлопрокат",
			},
		},
		{
			Name: "Полоса",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   2,
				Name: "Металлопрокат",
			},
		},
		{
			Name: "Квадрат",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   2,
				Name: "Металлопрокат",
			},
		},
		{
			Name: "Лист г/к",
			File: product.File{
				Image: "example.jpg",
				Video: "example.mp4",
			},
			SubCategory: product.SubCategory{
				ID:   2,
				Name: "Металлопрокат",
			},
		},
	}

	return s.Create(&category).Error
}

func (s *Seed) SubCategory() error {
	sub_category := []product.SubCategory{
		{
			Name: "Трубы стальные эл/сварные прямошовные",
		},
		{
			Name: "Металлопрокат",
		},
	}

	return s.Create(&sub_category).Error
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
