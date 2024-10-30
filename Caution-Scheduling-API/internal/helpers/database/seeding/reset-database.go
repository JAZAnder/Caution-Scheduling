package seeding

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"

)

func ResetDataTables(db *sql.DB, userName string) {
	database = db

	//Destroy Everything Except Logs and Global Settings
	dropLabHoursTable(userName)
	dropMeetingsTable(userName)
	dropLabsTable(userName)
	dropUserHoursTable(userName)
	dropUserSettingsTable(userName)
	dropHoursTable(userName)
	dropSessionCookieTable(userName)
	dropLocalUserTables(userName)
	

	//Rebuilding
	CreateTables(database)
	createConstraints(database)
}

func dropLocalUserTables(userName string) {
	query := "DROP TABLE `localusers`;"

	logger.Log(2, "database", "Drop Table", userName, "Dropping localUsers table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "localUsers table Dropped")
	}
}

func dropUserSettingsTable(userName string) {
	query := "DROP TABLE `userSettings`;"


	logger.Log(2, "database", "Drop Table", userName, "Dropping localUsers Settings table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "userSettings Dropped")
	}
}

func dropGlobalSettingsTable(userName string) {
	query := "CREATE TABLE IF NOT EXISTS userSettings(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`userName` varchar(255) unique," +
		"`ReceiveMeetingEmails` boolean NOT NULL" +
		");"

	logger.Log(2, "database", "Drop Table", userName, "Creating localUsers table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "userSettings table either created or already existed")
	}
}

func dropHoursTable(userName string) {
	query := "DROP TABLE `hours`"

	logger.Log(2, "database", "Drop Table", userName, "Drooping hours table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "hours table dropped")
	}

}

func dropLabsTable(userName string) {
	query := "DROP TABLE `labs`;"

	logger.Log(2, "database", "Drop Table", userName, "Dropping labs table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "labs table dropped")
	}

}

func dropMeetingsTable(userName string) {
	query := "DROP TABLE `meetings`;"

	logger.Log(2, "database", "Drop Table", userName, "Dropping meetings table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "meetings table dropped")
	}

}

func dropLabHoursTable(userName string) {
	query := "DROP TABLE `labHours`;"

	logger.Log(2, "database", "Drop Table", userName, "Dropping labHours table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "labHours table Dropped")
	}

}

func dropUserHoursTable(userName string) {
	query := "DROP TABLE `userHours`;"

	logger.Log(2, "database", "Drop Table", userName, "Dropping userHours table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "userHours table dropped")
	}

}

func dropSessionCookieTable(userName string) {
	query := "DROP TABLE `sessionCookie`;"

	logger.Log(2, "database", "Drop Table", userName, "Dropping sessionCookie table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "sessionCookie table dropped")
	}

}

func dropLogsTable(userName string) {
	query := "CREATE TABLE IF NOT EXISTS `logs`(" +
		"`id` INT AUTO_INCREMENT PRIMARY KEY, " +
		"`level` VARCHAR(10) NOT NULL," +
		"`category` VARCHAR(25) NOT NULL," +
		"`subCategory` VARCHAR(25) NOT NULL," +
		"`user` VARCHAR(50) NOT NULL," +
		"`message` VARCHAR(255) NOT NULL );"

	logger.Log(2, "database", "Drop Table", userName, "Creating logs table")

	logger.Log(1, "database", "Drop Table", userName, query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Drop Table", userName, err.Error())
	} else {
		logger.Log(2, "database", "Drop Table", userName, "logs table either created or already existed")
	}

}
