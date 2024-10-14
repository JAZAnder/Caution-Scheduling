package seeding

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
)

type DatabaseSeeder struct {
	db     *sql.DB
	logger *logger.Logger
}

func CreateTables(db *sql.DB, l *logger.Logger) {
	seeder := &DatabaseSeeder{
		db:     db,
		logger: l,
	}

	seeder.createLocalUserTables()
	seeder.createHoursTable()
	seeder.createLabsTable()
	seeder.createMeetingsTable()
	seeder.createLabHoursTable()
	seeder.createUserHoursTable()
	seeder.createSessionCookieTable()
	seeder.createLogsTable()
}

func (s *DatabaseSeeder) createLocalUserTables() {
	query := `CREATE TABLE IF NOT EXISTS localusers(
		userName varchar(255) PRIMARY KEY,
		firstName varchar(255) NOT NULL,
		lastName varchar(225) NOT NULL,
		email varchar(225) NOT NULL,
		password varchar(225) NOT NULL,
		isAdmin boolean NOT NULL
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating localUsers table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "localUsers table either created or already existed")
	}
}

func (s *DatabaseSeeder) createHoursTable() {
	query := `CREATE TABLE IF NOT EXISTS hours(
		Id int AUTO_INCREMENT PRIMARY KEY,
		startTime varchar(255) NOT NULL,
		endTime varchar(225) NOT NULL,
		dayOfWeek int DEFAULT NULL
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating hours table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "hours table either created or already existed")
	}
}

func (s *DatabaseSeeder) createLabsTable() {
	query := `CREATE TABLE IF NOT EXISTS labs(
		Id int AUTO_INCREMENT PRIMARY KEY,
		name varchar(255) NOT NULL,
		location varchar(225) DEFAULT NULL
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating labs table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "labs table either created or already existed")
	}
}

func (s *DatabaseSeeder) createUserHoursTable() {
	query := `CREATE TABLE IF NOT EXISTS userHours(
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		hoursId INT NOT NULL,
		available tinyint(1) NOT NULL DEFAULT '1',
		FOREIGN KEY (username) REFERENCES localusers(userName),
		FOREIGN KEY (hoursId) REFERENCES hours(Id)
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating userHours table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "userHours table either created or already existed")
	}
}

func (s *DatabaseSeeder) createMeetingsTable() {
	query := `CREATE TABLE IF NOT EXISTS meetings(
		id INT AUTO_INCREMENT PRIMARY KEY,
		tutorHourId int NOT NULL,
		labId int NOT NULL,
		studentName varchar(255) NOT NULL,
		studentEmail varchar(255) NOT NULL,
		date BIGINT NOT NULL,
		FOREIGN KEY (tutorHourId) REFERENCES userHours(id),
		FOREIGN KEY (labId) REFERENCES labs(Id)
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating meetings table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "meetings table either created or already existed")
	}
}

func (s *DatabaseSeeder) createLabHoursTable() {
	query := `CREATE TABLE IF NOT EXISTS labHours(
		id INT AUTO_INCREMENT PRIMARY KEY,
		labId int NOT NULL,
		hourId INT NOT NULL,
		tutorId INT NOT NULL,
		FOREIGN KEY (labId) REFERENCES labs(Id),
		FOREIGN KEY (hourId) REFERENCES hours(Id)
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating labHours table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "labHours table either created or already existed")
	}
}

func (s *DatabaseSeeder) createSessionCookieTable() {
	query := `CREATE TABLE IF NOT EXISTS sessionCookie(
		id INT AUTO_INCREMENT PRIMARY KEY,
		cookie VARCHAR(255) NOT NULL UNIQUE,
		username VARCHAR(255) NOT NULL,
		FOREIGN KEY (username) REFERENCES localusers(userName)
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating sessionCookie table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "sessionCookie table either created or already existed")
	}
}

func (s *DatabaseSeeder) createLogsTable() {
	query := `CREATE TABLE IF NOT EXISTS logs(
		id INT AUTO_INCREMENT PRIMARY KEY,
		timestamp DATETIME NOT NULL,
		level VARCHAR(10) NOT NULL,
		category VARCHAR(25) NOT NULL,
		subcategory VARCHAR(25) NOT NULL,
		user VARCHAR(50) NOT NULL,
		message TEXT NOT NULL,
		INDEX idx_user_category (user, category)
	);`

	s.logger.Log(logger.Info, "database", "Create Table", "System", "Creating logs table")
	s.logger.Log(logger.Debug, "database", "Create Table", "System", query)

	_, err := s.db.Exec(query)
	if err != nil {
		s.logger.Log(logger.Error, "database", "Create Table", "System", err.Error())
	} else {
		s.logger.Log(logger.Info, "database", "Create Table", "System", "logs table either created or already existed")
	}
}
