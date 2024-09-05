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
	createLogsTable()
}

func createLocalUserTables() {
	query := "CREATE TABLE IF NOT EXISTS localusers(" +
		"`userName` varchar(255) PRIMARY KEY," +
		"`firstName` varchar(255) NOT NULL," +
		"`lastName` varchar(225) NOT NULL," +
		"`email` varchar(225) NOT NULL," +
		"`password` varchar(225) NOT NULL," +
		"`isAdmin` boolean NOT NULL);"

	logger.Log(2, "database", "Create Table", "System", "Creating localUsers table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "localUsers table either created or already existed")
	}
}

func createHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS hours(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`startTime` varchar(255) NOT NULL," +
		"`endTime` varchar(225) NOT NULL," +
		"`dayOfWeek` int DEFAULT NULL);"

	logger.Log(2, "database", "Create Table", "System", "Creating hours table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "hours table either created or already existed")
	}

}

func createLabsTable() {
	query := "CREATE TABLE IF NOT EXISTS labs(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`name` varchar(255) NOT NULL," +
		"`location` varchar(225) DEFAULT NULL);"

	logger.Log(2, "database", "Create Table", "System", "Creating labs table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "labs table either created or already existed")
	}

}

func createMeetingsTable() {
	query := "CREATE TABLE IF NOT EXISTS meetings(" +
		"`tutorHourId` int DEFAULT NULL," +
		"`labId` int DEFAULT NULL," +
		"`studentName` varchar(255) NOT NULL," +
		"`studentEmail` varchar(255) NOT NULL," +
		"`date` BIGINT NOT NULL );"

	logger.Log(2, "database", "Create Table", "System", "Creating meetings table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "meetings table either created or already existed")
	}

}

func createLabHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS `MYSQL_DATABASE`.`labHours` (" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`labId` int DEFAULT NULL," +
		"`hourId` INT NOT NULL , " +
		"`tutorId` INT NOT NULL) ;"

	logger.Log(2, "database", "Create Table", "System", "Creating labHours table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "labHours table either created or already existed")
	}

}

func createUserHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS userHours(" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`username` VARCHAR(255) NOT NULL," +
		"`hoursId` INT NOT NULL , " +
		"`available` tinyint(1) NOT NULL DEFAULT '1' ) ;"

	logger.Log(2, "database", "Create Table", "System", "Creating userHours table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "userHours table either created or already existed")
	}

}

func createSessionCookieTable() {
	query := "CREATE TABLE IF NOT EXISTS `sessionCookie`(" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`cookie` VARCHAR(255) NOT NULL UNIQUE KEY," +
		"`username` VARCHAR(255) NOT NULL );"

	logger.Log(2, "database", "Create Table", "System", "Creating sessionCookie table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "sessionCookie table either created or already existed")
	}

}

func createLogsTable() {
	query := "CREATE TABLE IF NOT EXISTS `logs`(" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`level` VARCHAR(10) NOT NULL," +
		"`category` VARCHAR(25) NOT NULL," +
		"`subCategory` VARCHAR(25) NOT NULL," +
		"`user` VARCHAR(50) NOT NULL," +
		"`message` VARCHAR(255) NOT NULL );"

	logger.Log(2, "database", "Create Table", "System", "Creating logs table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "logs table either created or already existed")
	}

}
