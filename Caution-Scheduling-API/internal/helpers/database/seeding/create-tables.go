package seeding

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
)

var database *sql.DB

func CreateTables(db *sql.DB) {
	database = db

	createLocalUserTables()
	createHoursTable()
	createLabsTable()
}

func createLocalUserTables() {
	query := "CREATE TABLE IF NOT EXISTS localusers(" +
		"`userName` varchar(255) PRIMARY KEY," +
		"`firstName` varchar(255) NOT NULL," +
		"`lastName` varchar(225) NOT NULL," +
		"`email` varchar(225) NOT NULL," +
		"`password` varchar(225) NOT NULL," +
		"`isAdmin` boolean NOT NULL);"

	logger.Log(2, "database", "Creating localUsers table")

	logger.Log(1, "database", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", err.Error())
	} else {
		logger.Log(2, "database", "localUsers table either created or already existed")
	}
}

func createHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS hours(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`startTime` varchar(255) NOT NULL," +
		"`endTime` varchar(225) NOT NULL," +
		"`dayOfWeek` int DEFAULT NULL);"

	logger.Log(2, "database", "Creating hours table")

	logger.Log(1, "database", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", err.Error())
	} else {
		logger.Log(2, "database", "hours table either created or already existed")
	}

}

func createLabsTable() {
	query := "CREATE TABLE IF NOT EXISTS labs(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`name` varchar(255) NOT NULL," +
		"`location` varchar(225) DEFAULT NULL);"

	logger.Log(2, "database", "Creating labs table")

	logger.Log(1, "database", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", err.Error())
	} else {
		logger.Log(2, "database", "labs table either created or already existed")
	}

}
