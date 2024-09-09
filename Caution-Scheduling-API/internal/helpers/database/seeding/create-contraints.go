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

func createConstraints(db *sql.DB){
	database = db
	createSessionCookieConstraints()
}

func removeConstraints(db *sql.DB) {

	removeSessionCookieConstraints()
}

func createSessionCookieConstraints() {
	query := "ALTER TABLE `sessionCookie` " +
		" ADD CONSTRAINT `username -> localuser.userName`" +
		" FOREIGN KEY(`username`)" +
		" REFERENCES `localusers`(`userName`)" +
		" ON DELETE CASCADE;"

	logger.Log(2, "database", "Create Constraint", "System", "Creating sessionCookie.username -> localuser.userName Constraint")

	logger.Log(1, "database", "Create Constraint", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Create Constraint", "System", err.Error())
	} else {
		logger.Log(2, "database", "Create Constraint", "System", "sessionCookie.username -> localuser.userName Constraint created")
	}
}

func removeSessionCookieConstraints() {
	query := "ALTER TABLE sessionCookie DROP FOREIGN KEY `username -> localuser.userName`;"

	logger.Log(2, "database", "Remove Old Constraint", "System", "Removing sessionCookie.username -> localuser.userName Constraint")

	logger.Log(1, "database", "Remove Old Constraint", "System", query)

	_, err := database.Exec(query)
	if err != nil {
		logger.Log(4, "database", "Remove Old Constraint", "System", err.Error())
	} else {
		logger.Log(2, "database", "Remove Old Constraint", "System", "sessionCookie.username -> localuser.userName Constraint Removed")
	}
}
