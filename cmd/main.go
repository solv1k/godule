package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/solv1k/croco-api/cmd/api"
	"github.com/solv1k/croco-api/database"
	"github.com/solv1k/croco-api/internal/modules/catalog"
	"github.com/solv1k/croco-api/internal/modules/media"
	"github.com/solv1k/croco-api/internal/modules/seller"
	"github.com/solv1k/croco-api/internal/modules/user"
	fiberAuth "github.com/solv1k/croco-api/pkg/auth/fiber"
)

func main() {
	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connecting to the database
	db, err := database.Default()
	if err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
	}

	// Running migrations for the authentication package
	fiberAuth.RunMigrations(db)

	// Creating the API server configuration
	config := api.Config{
		AppName:     "Modular API Example",
		BaseRoute:   "/api/v1",
		AutoMigrate: true,
		Modules: []api.Module{
			media.NewModule(db),
			user.NewModule(db),
			catalog.NewModule(db),
			seller.NewModule(db),
		},
	}

	// Creating the API server
	app := api.New(config)

	// Running the API server (listen on port 3000)
	log.Fatal(app.Run(":3000"))
}
