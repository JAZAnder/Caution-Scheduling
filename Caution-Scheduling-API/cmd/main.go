package main

import (
	"log"
	"os"

	. "github.com/JAZAnder/Caution-Scheduling/internal/app"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/joho/godotenv"
)

// Entry Point
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error Loading .env file")
	}

	// Initialize database
	db := database.GetDatabase()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	defer db.Close()

	// Set up logger and create logs table
	logger.LogSetUpDb(db)

	a := App{}
	a.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	a.Run(port)
}
