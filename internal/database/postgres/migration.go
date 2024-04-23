package postgres

import (
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

func Migrate(database *Postgres) error {
	tables := []interface{}{
		&entity.User{},
	}

	// Use it for development only
	if err := database.Exec("DO $$DECLARE r RECORD; BEGIN FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE'; END LOOP; END;$$").Error; err != nil {
		return err
	}

	if err := database.AutoMigrate(tables...); err != nil {
		return err
	}

	// if err := author.SeedUser(database, config); err != nil {
	// 	return err
	// }

	return nil
}
