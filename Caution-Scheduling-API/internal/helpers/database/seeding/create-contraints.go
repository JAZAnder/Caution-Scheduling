package seeding

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"

)

func SetupConstraints(db *sql.DB) {
	database = db

	removeConstraints(db)
	createConstraints(db)

}

func createConstraints(db *sql.DB) {
	database = db
	createSessionCookieConstraints()
	createLabHourConstraints()
	createMeetingsConstraints()
}

func removeConstraints(db *sql.DB) {

	removeSessionCookieConstraints()
	removeLabHourConstraints()
	removeMeetingsConstraints()
}

func createSessionCookieConstraints() {
	query := "ALTER TABLE `sessionCookie` " +
		" ADD CONSTRAINT `username -> localuser.userName`" +
		" FOREIGN KEY(`username`)" +
		" REFERENCES `localusers`(`userName`)" +
		" ON DELETE CASCADE;"

	logger.Log(2, "database", "Create Constraint", "System", "sessionCookie Constraints")

	logger.Log(1, "database", "Create Constraint", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Constraint", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Constraint", "System", "sessionCookie Constraints created")
	}
}

func removeSessionCookieConstraints() {
	query := "ALTER TABLE sessionCookie DROP FOREIGN KEY `username -> localuser.userName`;"

	logger.Log(2, "database", "Remove Constraint", "System", "Removing sessionCookie Constraints")

	logger.Log(1, "database", "Remove Constraint", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Remove Constraint", "System", err.Error())
	} else {
		logger.Log(2, "database", "Remove Constraint", "System", "sessionCookie Constraints Removed")
	}
}

func createLabHourConstraints() {
	logger.Log(2, "database", "Create Constraint", "System", "Creating labHours Constraints")
	var masterErr error = nil
	query := "ALTER TABLE `labHours` " +
		" ADD CONSTRAINT `LabId -> Lab.Id`" +
		" FOREIGN KEY (`LabId`)" +
		" REFERENCES `labs`(`Id`)" +
		" ON DELETE RESTRICT; "

	logger.Log(1, "database", "Create Constraint", "System", query)
	_, err := database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Create Constraint", "System", err.Error())
	}

	query = "ALTER TABLE `labHours` " +
		" ADD CONSTRAINT `labHours.HourId -> hour.Id`" +
		"  FOREIGN KEY (`hourId`)" +
		" REFERENCES `hours`(`Id`)" +
		" ON DELETE RESTRICT; "

	logger.Log(1, "database", "Create Constraint", "System", query)
	_, err = database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Create Constraint", "System", err.Error())
	}

	query = "ALTER TABLE `labHours` " +
		" ADD CONSTRAINT `TutorId -> UserHours.Id`" +
		" FOREIGN KEY (`TutorId`)" +
		" REFERENCES `userHours`(`Id`)" +
		" ON DELETE RESTRICT; "

	logger.Log(1, "database", "Create Constraint", "System", query)
	_, err = database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Create Constraint", "System", err.Error())
	}

	if masterErr != nil {
		logger.Log(2, "database", "Create Constraint", "System", "labHours Constraints created with Errors")
	} else {
		logger.Log(2, "database", "Create Constraint", "System", "labHours Constraints created")
	}
}

func removeLabHourConstraints() {
	logger.Log(2, "database", "Remove Constraint", "System", "Removing labHours Constraints")
	var masterErr error = nil

	query := "ALTER TABLE labHours DROP FOREIGN KEY `LabId -> Lab.Id`;"
	logger.Log(1, "database", "Remove Constraint", "System", query)
	_, err := database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Remove Constraint", "System", err.Error())
	}

	query = "ALTER TABLE labHours DROP FOREIGN KEY `labHours.HourId -> hour.Id`;"
	logger.Log(1, "database", "Remove Old Constraint", "System", query)
	_, err = database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Remove Constraint", "System", err.Error())
	}

	query = "ALTER TABLE labHours DROP FOREIGN KEY `TutorId -> UserHours.Id`;"
	logger.Log(1, "database", "Remove Constraint", "System", query)
	_, err = database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Remove Constraint", "System", err.Error())
	}

	if masterErr != nil {
		logger.Log(2, "database", "Remove Constraint", "System", "labHours Constraints Removed with Errors")
	} else {
		logger.Log(2, "database", "Remove Constraint", "System", "labHours Constraints Removed")
	}
}

func createMeetingsConstraints() {
	logger.Log(2, "database", "Create Constraint", "System", "Creating Meeting Constraints")
	var masterErr error = nil
	query := "ALTER TABLE `meetings` " +
		" ADD CONSTRAINT `meetings.LabId-> Lab.Id`" +
		" FOREIGN KEY (`LabId`)" +
		" REFERENCES `labs`(`Id`)" +
		" ON DELETE RESTRICT; "

	logger.Log(1, "database", "Create Constraint", "System", query)
	_, err := database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Create Constraint", "System", err.Error())
	}

	if masterErr != nil {
		logger.Log(2, "database", "Create Constraint", "System", "Meeting Constraints created with Errors")
	} else {
		logger.Log(2, "database", "Create Constraint", "System", "Meeting Constraints created")
	}
}

func removeMeetingsConstraints() {
	logger.Log(2, "database", "Remove Constraint", "System", "Removing Meeting Constraints")
	var masterErr error = nil

	query := "ALTER TABLE meetings DROP FOREIGN KEY `meetings.LabId-> Lab.Id`;"
	logger.Log(1, "database", "Remove Constraint", "System", query)
	_, err := database.Exec(query)

	if err != nil {
		masterErr = err
		logger.Log(4, "database", "Remove Constraint", "System", err.Error())
	}

	if masterErr != nil {
		logger.Log(2, "database", "Remove Constraint", "System", "Meeting Constraints Removed with Errors")
	} else {
		logger.Log(2, "database", "Remove Constraint", "System", "Meeting Constraints Removed")
	}
}
