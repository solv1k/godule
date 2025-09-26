package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/solv1k/croco-api/cmd/seed/seeders"
	"github.com/solv1k/croco-api/database"
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
		log.Fatal("Database connection failed: " + err.Error())
	}

	// Checking if the count argument is provided
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run cmd/seed/main.go [key] [count]")
	}

	// Parsing the key
	key := os.Args[1]

	// Parsing the count
	count := os.Args[2]
	countInt, err := strconv.Atoi(count)
	if err != nil {
		log.Fatal("Invalid seeding count: " + err.Error())
	}

	// Running the seeders
	if err := seeders.Run(db, key, countInt); err != nil {
		log.Fatal("Seeding database failed: " + err.Error())
	}
}
