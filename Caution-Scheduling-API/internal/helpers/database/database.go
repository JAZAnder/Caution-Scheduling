package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/joho/godotenv"
)

var once sync.Once

type database struct {
	DB *sql.DB
}

var (
	db database
)

func GetDatabase() *sql.DB {
	once.Do(func() {
		createDatabase()
	})

	return db.DB
}

func createDatabase() {
	user := "APP_DB_USERNAME"
	password := "APP_DB_PASSWORD"
	dbServer := "APP_DB"
	dbName := "APP_DB_NAME"

	err := godotenv.Load()
	if err != nil {
		log.Println("Error Loading .env file in database")
	} else {
		user = os.Getenv("APP_DB_USERNAME")
		password = os.Getenv("APP_DB_PASSWORD")
		dbServer = os.Getenv("APP_DB")
		dbName = os.Getenv("APP_DB_NAME")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, dbServer, dbName)

	// Create MySQL Connection
	dbConn, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Check if the connection is actually established
	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	db.DB = dbConn

	db = database{
		DB: dbConn, // Assign the connection here
	}

	dbLogger := db.DB // Test if the reference works
	logger.LogSetUpDb(dbLogger)

}
