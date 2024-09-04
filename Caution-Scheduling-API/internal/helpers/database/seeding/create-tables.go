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
	createMeetingsTable()
	createLabHoursTable()
	createUserHoursTable()
	createSessionCookieTable()
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

func createMeetingsTable() {
	query := "CREATE TABLE IF NOT EXISTS meetings(" +
		"`tutorHourId` int DEFAULT NULL," +
		"`labId` int DEFAULT NULL," +
		"`studentName` varchar(255) NOT NULL" +
		"`studentEmail` varchar(255) NOT NULL," +
		"`date` BIGINT NOT NULL );"

	logger.Log(2, "database", "Creating meetings table")

	logger.Log(1, "database", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", err.Error())
	} else {
		logger.Log(2, "database", "meetings table either created or already existed")
	}

}

func createLabHoursTable() {
	query := "CREATE TABLE `MYSQL_DATABASE`.`labHours` (" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`labId` int DEFAULT NULL," +
		"`hourId` INT NOT NULL , " +
		"`tutorId` INT NOT NULL) ;"

	logger.Log(2, "database", "Creating labHours table")

	logger.Log(1, "database", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", err.Error())
	} else {
		logger.Log(2, "database", "labHours table either created or already existed")
	}

}

func createUserHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS userHours(" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`username` VARCHAR(255) NOT NULL," +
		"`hoursId` INT NOT NULL , " +
		"`available` tinyint(1) NOT NULL DEFAULT '1' ) ;"

	logger.Log(2, "database", "Creating userHours table")

	logger.Log(1, "database", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", err.Error())
	} else {
		logger.Log(2, "database", "userHours table either created or already existed")
	}

}

func createSessionCookieTable() {
	query := "CREATE TABLE IF NOT EXISTS `sessionCookie`(" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`cookie` VARCHAR(255) NOT NULL UNIQUE KEY," +
		"`username` VARCHAR(255) NOT NULL );"

	logger.Log(2, "database", "Creating sessionCookie table")

	logger.Log(1, "database", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", err.Error())
	} else {
		logger.Log(2, "database", "sessionCookie table either created or already existed")
	}

}
