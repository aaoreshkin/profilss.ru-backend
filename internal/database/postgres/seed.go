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

func (s *Seed) Measure() error {
	measure := []product.Measure{
		{
			Code: "г",
			Name: "Грамм",
		},
		{
			Code: "кг",
			Name: "Килограмм",
		},
		{
			Code: "шт",
			Name: "Штук",
		},
		{
			Code: "тн",
			Name: "Тонна",
		},
		{
			Code: "м3",
			Name: "Метр кубический",
		},
		{
			Code: "м2",
			Name: "Метр квадратный",
		},
		{
			Code: "м",
			Name: "Метр",
		},
		{
			Code: "см",
			Name: "Сантиметр",
		},
	}

	return s.Create(&measure).Error
}

func (s *Seed) Characteristic() error {
	characteristic := []product.Characteristic{
		{
			Name:        "Вес",
			Description: "Вес",
		},
		{
			Name:        "Длина",
			Description: "Длина",
		},
		{
			Name:        "Ширина",
			Description: "Ширина",
		},
		{
			Name:        "Высота",
			Description: "Высота",
		},
		{
			Name:        "Объем",
			Description: "Объем",
		},
		{
			Name:        "Площадь",
			Description: "Площадь",
		},
		{
			Name:        "Толщина",
			Description: "Толщина",
		},
		{
			Name:        "Диаметр",
			Description: "Диаметр",
		},
		{
			Name:        "Количество",
			Description: "Количество",
		},
		{
			Name:        "Размер",
			Description: "Размер",
		},
		{
			Name:        "Глубина",
			Description: "Глубина",
		},
		{
			Name:        "Вместимость",
			Description: "Вместимость",
		},
	}

	return s.Create(&characteristic).Error
}

func (s *Seed) Category() error {
	characteristic := []product.Category{
		{
			Name:        "Труба плоскоовальная",
			Description: "-",
		},
		{
			Name:        "Труба квадратная",
			Description: "-",
		},
		{
			Name:        "Труба круглая",
			Description: "-",
		},
		{
			Name:        "Труба овальная",
			Description: "-",
		},
		{
			Name:        "Труба прямоугольная",
			Description: "-",
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
