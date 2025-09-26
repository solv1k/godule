package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/solv1k/croco-api/cmd/seed/seeders"
)

func main() {
	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	runner := seeders.NewRunner()
	if err := runner.Run(key, countInt); err != nil {
		log.Fatal("Seeding database failed: " + err.Error())
	}
}
