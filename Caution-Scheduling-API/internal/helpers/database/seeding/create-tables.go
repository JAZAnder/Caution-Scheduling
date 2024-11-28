package seeding

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
)

func CreateTables(db *sql.DB) {
	database = db

	createLocalUserTables()
	createUserSettingsTable()
	createHoursTable()
	createLabsTable()
	createMeetingsTable()
	createLabHoursTable()
	createUserHoursTable()
	createSessionCookieTable()
	createLogsTable()
	createGlobalSettingsTable()
	createTopicsTable()
	createNotesTable()

}

func createLocalUserTables() {
	query := "CREATE TABLE IF NOT EXISTS localusers(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`userName` varchar(255) unique," +
		"`firstName` varchar(255) NOT NULL," +
		"`lastName` varchar(225) NOT NULL," +
		"`email` varchar(225) NOT NULL," +
		"`password` varchar(225) NOT NULL," +
		"`role` int NOT NULL," +
		"`fullName` varchar(225) NOT NULL," +
		"`googleId` varchar(225)," +
		"`strikes` int DEFAULT 0," +
		"`isAdmin` boolean NOT NULL" +
		");"

	logger.Log(2, "database", "Create Table", "System", "Creating localUsers table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "localUsers table either created or already existed")
	}
}

func createUserSettingsTable() {
	query := "CREATE TABLE IF NOT EXISTS userSettings(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`userName` varchar(255) unique," +
		"`ReceiveMeetingEmails` boolean NOT NULL" +
		");"

	logger.Log(2, "database", "Create Table", "System", "Creating localUsers table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "userSettings table either created or already existed")
	}
}

func createGlobalSettingsTable() {
	query := "CREATE TABLE IF NOT EXISTS userSettings(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`userName` varchar(255) unique," +
		"`ReceiveMeetingEmails` boolean NOT NULL" +
		");"

	logger.Log(2, "database", "Create Table", "System", "Creating localUsers table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "userSettings table either created or already existed")
	}
}

func createHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS hours(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`timeCode` INT NOT NULL," +
		"`startTime` varchar(255) NOT NULL," +
		"`endTime` varchar(225) NOT NULL," +
		"`dayOfWeek` int DEFAULT NULL," +
		"`active` boolean DEFAULT 0, " +
		"UNIQUE INDEX `timeCode_UNIQUE` (`timeCode`, `dayOfWeek`));"

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
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`tutorHourId` int DEFAULT NULL," +
		"`studentId` int DEFAULT NULL," +
		"`date` BIGINT NOT NULL ," +
		"`state` int DEFAULT 1," +
		"`topicId` int DEFAULT NULL);"
		//State 0 = Deleted
		//State 1 = Active
		//State 2 = Archived

	logger.Log(2, "database", "Create Table", "System", "Creating meetings table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "meetings table either created or already existed")
	}

}

func createTopicsTable() {
	query := "CREATE TABLE IF NOT EXISTS topic(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`description` VARCHAR(255) DEFAULT NULL);"

	logger.Log(2, "database", "Create Table", "System", "Creating topic table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "topic table either created or already existed")
	}

}

func createNotesTable() {
	query := "CREATE TABLE IF NOT EXISTS note(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`userId` INT NOT NULL ," +
		"`meetingId` INT NOT NULL ," +
		"`note` VARCHAR(255) DEFAULT NULL);"

	logger.Log(2, "database", "Create Table", "System", "Creating note table")

	logger.Log(1, "database", "Create Table", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Table", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Table", "System", "note table either created or already existed")
	}

}

func createLabHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS `labHours` (" +
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
		"`userId` INT NOT NULL ," +
		"`hourId` INT NOT NULL , " +
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
