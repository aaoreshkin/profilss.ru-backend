package postgres

import (
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

func Migrate(database *Postgres) error {
	tables := []interface{}{
		&entity.User{},
		&entity.AccessLevel{},
	}

	// Use it for development only
	if err := DropTables(database); err != nil {
		return err
	}

	if err := database.AutoMigrate(tables...); err != nil {
		return err
	}

	if err := SeedAccessLevel(database); err != nil {
		return err
	}

	if err := SeedUser(database); err != nil {
		return err
	}

	return nil
}

func DropTables(database *Postgres) error {
	return database.Exec("DO $$DECLARE r RECORD; BEGIN FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE'; END LOOP; END;$$").Error
}

func SeedAccessLevel(database *Postgres) error {
	roles := []entity.AccessLevel{
		{
			Title: "Администратор",
		},
		{
			Title: "Менеджер",
		},
	}

	return database.Create(&roles).Error
}

func SeedUser(database *Postgres) error {
	const email = "oreshkin@ya.ru"
	const password = "password"

	// Hash raw password
	hashedPassword, err := common.HashPassword(password)
	if err != nil {
		return err
	}

	var accessLevelID string

	// Get access level id by title
	if err := database.Model(&entity.AccessLevel{}).Select("id").Where("title = ?", "Администратор").First(&accessLevelID).Error; err != nil {
		return err
	}

	// Hash access token
	accessToken, err := common.GenerateJWT(email, accessLevelID)
	if err != nil {
		return err
	}

	user := entity.User{
		Email:         email,
		Password:      hashedPassword,
		AccessLevelID: accessLevelID,
		AccessToken:   accessToken,
	}

	return database.Create(&user).Error
}
