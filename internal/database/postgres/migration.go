package postgres

import (
	"github.com/oreshkindev/profilss.ru-backend/common"
	bid "github.com/oreshkindev/profilss.ru-backend/internal/bid/entity"
	post "github.com/oreshkindev/profilss.ru-backend/internal/post/entity"
	product "github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
	user "github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

func Migrate(database *Postgres) error {
	tables := []interface{}{
		&user.User{},
		&user.Permission{},
		&post.Post{},
		&product.Measure{},
		&product.Characteristic{},
		&product.Product{},
		&product.ProductsCharacteristics{},
		&bid.Bid{},
	}

	// Use it for development only
	if err := DropTables(database); err != nil {
		return err
	}

	if err := database.AutoMigrate(tables...); err != nil {
		return err
	}

	if err := SeedPermission(database); err != nil {
		return err
	}

	if err := SeedUser(database); err != nil {
		return err
	}

	if err := SeedMeasure(database); err != nil {
		return err
	}

	if err := SeedCharacteristic(database); err != nil {
		return err
	}

	return nil
}

func DropTables(database *Postgres) error {
	return database.Exec("DO $$DECLARE r RECORD; BEGIN FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE'; END LOOP; END;$$").Error
}

func SeedPermission(database *Postgres) error {
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

	return database.Create(&roles).Error
}

func SeedUser(database *Postgres) error {
	const email = "oreshkin.dev@outlook.com"
	const password = "password"

	// Hash raw password
	hashedPassword, err := common.HashPassword(password)
	if err != nil {
		return err
	}

	var permissionID string

	// Get permissionID by title
	if err := database.Model(&user.Permission{}).Select("id").Where("rule = ?", "Superuser").First(&permissionID).Error; err != nil {
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

	return database.Create(&user).Error
}

func SeedMeasure(database *Postgres) error {
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

	for i := range measure {
		err := database.Model(&product.Measure{}).Create(&measure[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func SeedCharacteristic(database *Postgres) error {
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

	for i := range characteristic {
		err := database.Model(&product.Characteristic{}).Create(&characteristic[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}
