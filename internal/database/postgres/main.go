package postgres

import (
	"sync"

	"github.com/oreshkindev/profilss.ru-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	*gorm.DB
}

var (
	database *Postgres
	once     sync.Once
)

func Dial(config *config.Database) (*Postgres, error) {
	once.Do(func() {
		conn, err := gorm.Open(postgres.Open(config.URL), &gorm.Config{})
		if err != nil {
			return
		}

		database = &Postgres{conn}
	})

	if database != nil {
		if err := Migrate(database); err != nil {
			return nil, err
		}
	}

	return database, nil
}
