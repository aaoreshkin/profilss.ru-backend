package postgres

import (
	bid "github.com/oreshkindev/profilss.ru-backend/internal/bid/entity"
	hr "github.com/oreshkindev/profilss.ru-backend/internal/hr/entity"
	post "github.com/oreshkindev/profilss.ru-backend/internal/post/entity"
	product "github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
	service "github.com/oreshkindev/profilss.ru-backend/internal/service/entity"
	setting "github.com/oreshkindev/profilss.ru-backend/internal/setting/entity"
	support "github.com/oreshkindev/profilss.ru-backend/internal/support/entity"
	user "github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

func Migrate(database *Postgres) error {
	tables := []interface{}{
		&user.User{},
		&user.Permission{},
		&post.Post{},
		&service.Service{},
		&product.Iso{},
		&product.Characteristic{},
		&product.Category{},
		&product.SubCategory{},
		&product.File{},
		&product.Product{},
		&bid.Bid{},
		&hr.Hr{},
		&setting.Setting{},
		&setting.Contact{},
		&setting.Media{},
		&setting.Employee{},
		&support.Support{},
	}

	// Use it for development only
	if err := DropTables(database); err != nil {
		return err
	}

	if err := database.AutoMigrate(tables...); err != nil {
		return err
	}

	seed, err := NewSeed(database)
	if err != nil {
		return err
	}

	if err := seed.Seed(); err != nil {
		return err
	}

	return nil
}

func DropTables(database *Postgres) error {
	return database.Exec("DO $$DECLARE r RECORD; BEGIN FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE'; END LOOP; END;$$").Error
}
