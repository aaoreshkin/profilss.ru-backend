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
	if err := s.Measure(); err != nil {
		return err
	}

	if err := s.Characteristic(); err != nil {
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

func (s *Seed) Measure() error {
	measure := []product.Measure{
		{
			Code:  "г",
			Title: "Грамм",
		},
		{
			Code:  "кг",
			Title: "Килограмм",
		},
		{
			Code:  "шт",
			Title: "Штук",
		},
		{
			Code:  "тн",
			Title: "Тонна",
		},
		{
			Code:  "м3",
			Title: "Метр кубический",
		},
		{
			Code:  "м2",
			Title: "Метр квадратный",
		},
		{
			Code:  "м",
			Title: "Метр",
		},
		{
			Code:  "см",
			Title: "Сантиметр",
		},
	}

	return s.Create(&measure).Error
}

func (s *Seed) Characteristic() error {
	characteristic := []product.Characteristic{
		{
			Title:       "Вес",
			Description: "Вес",
		},
		{
			Title:       "Длина",
			Description: "Длина",
		},
		{
			Title:       "Ширина",
			Description: "Ширина",
		},
		{
			Title:       "Высота",
			Description: "Высота",
		},
		{
			Title:       "Объем",
			Description: "Объем",
		},
		{
			Title:       "Площадь",
			Description: "Площадь",
		},
		{
			Title:       "Толщина",
			Description: "Толщина",
		},
		{
			Title:       "Диаметр",
			Description: "Диаметр",
		},
		{
			Title:       "Количество",
			Description: "Количество",
		},
		{
			Title:       "Размер",
			Description: "Размер",
		},
		{
			Title:       "Глубина",
			Description: "Глубина",
		},
		{
			Title:       "Вместимость",
			Description: "Вместимость",
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
	const email = "oreshkin.dev@outlook.com"
	const password = "password"

	// Hash raw password
	hashedPassword, err := common.HashPassword(password)
	if err != nil {
		return err
	}

	var permissionID string

	// Get permissionID by title
	if err := s.Model(&user.Permission{}).Select("id").Where("rule = ?", "Superuser").First(&permissionID).Error; err != nil {
		return err
	}

	// Hash access token
	accessToken, err := common.HashToken(email, permissionID)
	if err != nil {
		return err
	}

	user := user.User{
		AccessToken:  accessToken,
		Email:        email,
		Password:     hashedPassword,
		PermissionID: permissionID,
	}

	return s.Create(&user).Error
}
