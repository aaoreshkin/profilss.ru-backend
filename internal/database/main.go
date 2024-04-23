package database

import (
	"github.com/oreshkindev/profilss.ru-backend/config"
	"github.com/oreshkindev/profilss.ru-backend/internal/database/postgres"
)

type Database struct {
	*postgres.Postgres
}

func NewDatabase(config *config.Config) (*Database, error) {
	database, err := postgres.Dial(&config.Database)
	if err != nil {
		panic(err)
	}

	return &Database{database}, nil
}
