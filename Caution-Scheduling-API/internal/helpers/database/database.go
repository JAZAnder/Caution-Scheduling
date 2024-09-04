package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/database/seeding"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"

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

	//Creates MySQL Connection
	db.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	logger.LogSetUpDb(1, db.DB)

	seeding.CreateTables(db.DB)

}
