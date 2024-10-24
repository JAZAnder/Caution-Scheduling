package seeding

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"
)

var database *sql.DB

func SeedData(db *sql.DB) {
	database = db
	seedUsers()

}

func seedUsers() {
	//Seeding Administrator User
	var Admin user.LocalUser = user.LocalUser{
		UserName:  "Admin",
		FirstName: "System",
		LastName:  "Administrator",
		FullName:  "System Administrator",
		Email:     "admin@localhost.com",
		Password:  "P@33word123!",
		IsAdmin:   true,
		Role:      4,
	}

	err := Admin.SignUp(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", Admin.UserName+" user is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	//Seeding Student User
	err = nil
	var student user.LocalUser = user.LocalUser{
		UserName:  "Student",
		FirstName: "Normal",
		LastName:  "Student",
		FullName:  "Normal Student",
		Email:     "student@localhost.com",
		Password:  "P@33word123!",
		IsAdmin:   false,
		Role:      1,
	}

	err = student.SignUp(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", student.UserName+" user is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	//Supervisor  User
	err = nil
	var supervisor user.LocalUser = user.LocalUser{
		UserName:  "Supervisor",
		FirstName: "Tutor",
		LastName:  "Supervisor",
		FullName:  "Tutor Supervisor",
		Email:     "supervisor@localhost.com",
		Password:  "P@33word123!",
		IsAdmin:   false,
		Role:      3,
	}

	err = supervisor.SignUp(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", supervisor.UserName+" user is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	//Tutor  User
	err = nil
	var Tutor user.LocalUser = user.LocalUser{
		UserName:  "Tutor",
		FirstName: "School",
		LastName:  "Tutor",
		FullName:  "School Tutor",
		Email:     "tutor@localhost.com",
		Password:  "P@33word123!",
		IsAdmin:   false,
		Role:      2,
	}

	err = Tutor.SignUp(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", Tutor.UserName+" user is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	err = nil
	var anotherTutor user.LocalUser = user.LocalUser{
		UserName:  "Tutor2",
		FirstName: "Jane",
		LastName:  "Doe",
		FullName:  "Jane Doe",
		Email:     "jane.doe@localhost.com",
		Password:  "P@33word123!",
		IsAdmin:   false,
		Role:      2,
	}

	err = anotherTutor.SignUp(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", anotherTutor.UserName+" user is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	//Timeslot1
	err = nil
	timeSlot1 := hour.Hour{
		StartTime: "9:00 AM",
		EndTime:   "10:15 AM",
		DayOfWeek: 1,
	}

	err = timeSlot1.CreateHour(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", timeSlot1.StartTime+" - "+timeSlot1.EndTime+" Timeslot is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	//Assign Tutor to Timeslot1
	err = nil

	userHour1 := userHour.UserHour{}

	userHour1.Tutor = "School Tutor"
	hours, _ := hour.GetHours(database)
	userHour1.HourId = hours[0].Id

	userHour1.CreateUserHour(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", "Tutor: "+userHour1.Tutor+" User Hour is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

}
