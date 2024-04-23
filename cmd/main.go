package main

import (
	"net/http"

	"github.com/oreshkindev/profilss.ru-backend/config"
	"github.com/oreshkindev/profilss.ru-backend/internal"
	"github.com/oreshkindev/profilss.ru-backend/internal/database"
	"github.com/oreshkindev/profilss.ru-backend/internal/router"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	config, err := config.NewConfig()
	if err != nil {
		return err
	}

	database, err := database.NewDatabase(config)
	if err != nil {
		return err
	}

	manager := internal.NewManager(database)

	router, err := router.NewRouter(manager)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.Port, router)
}
