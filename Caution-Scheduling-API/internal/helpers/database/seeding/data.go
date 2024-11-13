package seeding

import (
	"database/sql"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/meeting"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/topic"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"

)

var database *sql.DB

func SeedData(db *sql.DB) {
	database = db
	seedTimeSlots()
	seedUsers()
	seedUserHours()
	seedTopics()
	seedMeetings()
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
	var Tutor2 user.LocalUser = user.LocalUser{
		UserName:  "Tutor",
		FirstName: "School",
		LastName:  "Tutor",
		FullName:  "School Tutor",
		Email:     "tutor@localhost.com",
		Password:  "P@33word123!",
		IsAdmin:   false,
		Role:      2,
	}

	err = Tutor2.SignUp(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", Tutor2.UserName+" user is Created")
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}
}

func seedTimeSlots() {
	timeSlots := []hour.Hour{
		{
			TimeCode:  7000715,
			StartTime: "7:00 AM",
			EndTime:   "7:15 AM",
			Active:    false,
		},
		{
			TimeCode:  7150730,
			StartTime: "7:15 AM",
			EndTime:   "7:30 AM",
			Active:    false,
		},
		{
			TimeCode:  7300745,
			StartTime: "7:30 AM",
			EndTime:   "7:45 AM",
			Active:    false,
		},
		{
			TimeCode:  7450800,
			StartTime: "7:45 AM",
			EndTime:   "8:00 AM",
			Active:    false,
		},
		{
			TimeCode:  8000815,
			StartTime: "8:00 AM",
			EndTime:   "8:15 AM",
			Active:    false,
		},
		{
			TimeCode:  8150830,
			StartTime: "8:15 AM",
			EndTime:   "8:30 AM",
			Active:    false,
		},
		{
			TimeCode:  8300845,
			StartTime: "8:30 AM",
			EndTime:   "8:45 AM",
			Active:    false,
		},
		{
			TimeCode:  8450900,
			StartTime: "8:45 AM",
			EndTime:   "9:00 AM",
			Active:    false,
		}, {
			TimeCode:  9000915,
			StartTime: "9:00 AM",
			EndTime:   "9:15 AM",
			Active:    true,
		},
		{
			TimeCode:  9150930,
			StartTime: "9:15 AM",
			EndTime:   "9:30 AM",
			Active:    true,
		},
		{
			TimeCode:  9300945,
			StartTime: "9:30 AM",
			EndTime:   "9:45 AM",
			Active:    true,
		},
		{
			TimeCode:  9451000,
			StartTime: "9:45 AM",
			EndTime:   "10:00 AM",
			Active:    true,
		},
		{
			TimeCode:  10001015,
			StartTime: "10:00 AM",
			EndTime:   "10:15 AM",
			Active:    true,
		},
		{
			TimeCode:  10151030,
			StartTime: "10:15 AM",
			EndTime:   "10:30 AM",
			Active:    true,
		},
		{
			TimeCode:  10301045,
			StartTime: "10:30 AM",
			EndTime:   "10:45 AM",
			Active:    true,
		},
		{
			TimeCode:  10451100,
			StartTime: "10:45 AM",
			EndTime:   "11:00 AM",
			Active:    true,
		},
		{
			TimeCode:  11001115,
			StartTime: "11:00 AM",
			EndTime:   "11:15 AM",
			Active:    true,
		},
		{
			TimeCode:  11151130,
			StartTime: "11:15 AM",
			EndTime:   "11:30 AM",
			Active:    true,
		},
		{
			TimeCode:  11301145,
			StartTime: "11:30 AM",
			EndTime:   "11:45 AM",
			Active:    true,
		},
		{
			TimeCode:  11451200,
			StartTime: "11:45 AM",
			EndTime:   "12:00 PM",
			Active:    true,
		},
		{
			TimeCode:  12001215,
			StartTime: "12:00 PM",
			EndTime:   "12:15 PM",
			Active:    true,
		},
		{
			TimeCode:  12151230,
			StartTime: "12:15 PM",
			EndTime:   "12:30 PM",
			Active:    true,
		},
		{
			TimeCode:  12301245,
			StartTime: "12:30 PM",
			EndTime:   "12:45 PM",
			Active:    true,
		},
		{
			TimeCode:  12451300,
			StartTime: "12:45 PM",
			EndTime:   "1:00 PM",
			Active:    true,
		},
		{
			TimeCode:  13001315,
			StartTime: "1:00 PM",
			EndTime:   "1:15 PM",
			Active:    true,
		},
		{
			TimeCode:  13151330,
			StartTime: "1:15 PM",
			EndTime:   "1:30 PM",
			Active:    true,
		},
		{
			TimeCode:  13301345,
			StartTime: "1:30 PM",
			EndTime:   "1:45 PM",
			Active:    true,
		},
		{
			TimeCode:  13451400,
			StartTime: "1:45 PM",
			EndTime:   "2:00 PM",
			Active:    true,
		},
		{
			TimeCode:  14001415,
			StartTime: "2:00 PM",
			EndTime:   "2:15 PM",
			Active:    true,
		},
		{
			TimeCode:  14151430,
			StartTime: "2:15 PM",
			EndTime:   "2:30 PM",
			Active:    true,
		},
		{
			TimeCode:  14301445,
			StartTime: "2:30 PM",
			EndTime:   "2:45 PM",
			Active:    true,
		},
		{
			TimeCode:  14451500,
			StartTime: "2:45 PM",
			EndTime:   "3:00 PM",
			Active:    true,
		},
		{
			TimeCode:  15001515,
			StartTime: "3:00 PM",
			EndTime:   "3:15 PM",
			Active:    true,
		},
		{
			TimeCode:  15151530,
			StartTime: "3:15 PM",
			EndTime:   "3:30 PM",
			Active:    true,
		},
		{
			TimeCode:  15301545,
			StartTime: "3:30 PM",
			EndTime:   "3:45 PM",
			Active:    true,
		},
		{
			TimeCode:  15451600,
			StartTime: "3:45 PM",
			EndTime:   "4:00 PM",
			Active:    true,
		},
		{
			TimeCode:  16001615,
			StartTime: "4:00 PM",
			EndTime:   "4:15 PM",
			Active:    true,
		},
		{
			TimeCode:  16151630,
			StartTime: "4:15 PM",
			EndTime:   "4:30 PM",
			Active:    true,
		},
		{
			TimeCode:  16301645,
			StartTime: "4:30 PM",
			EndTime:   "4:45 PM",
			Active:    true,
		},
		{
			TimeCode:  16451700,
			StartTime: "4:45 PM",
			EndTime:   "5:00 PM",
			Active:    true,
		},
		{
			TimeCode:  17001715,
			StartTime: "5:00 PM",
			EndTime:   "5:15 PM",
			Active:    true,
		},
		{
			TimeCode:  17151730,
			StartTime: "5:15 PM",
			EndTime:   "5:30 PM",
			Active:    true,
		},
		{
			TimeCode:  17301745,
			StartTime: "5:30 PM",
			EndTime:   "5:45 PM",
			Active:    true,
		},
		{
			TimeCode:  17451800,
			StartTime: "5:45 PM",
			EndTime:   "6:00 PM",
			Active:    true,
		}, {
			TimeCode:  18001815,
			StartTime: "6:00 PM",
			EndTime:   "6:15 PM",
			Active:    false,
		},
		{
			TimeCode:  18151830,
			StartTime: "6:15 PM",
			EndTime:   "6:30 PM",
			Active:    false,
		},
		{
			TimeCode:  18301845,
			StartTime: "6:30 PM",
			EndTime:   "6:45 PM",
			Active:    false,
		},
		{
			TimeCode:  18451900,
			StartTime: "6:45 PM",
			EndTime:   "7:00 PM",
			Active:    false,
		},
		{
			TimeCode:  19001915,
			StartTime: "7:00 PM",
			EndTime:   "7:15 PM",
			Active:    false,
		},
		{
			TimeCode:  19151930,
			StartTime: "7:15 PM",
			EndTime:   "7:30 PM",
			Active:    false,
		},
		{
			TimeCode:  19301945,
			StartTime: "7:30 PM",
			EndTime:   "7:45 PM",
			Active:    false,
		},
		{
			TimeCode:  19452000,
			StartTime: "7:45 PM",
			EndTime:   "8:00 PM",
			Active:    false,
		},
		{
			TimeCode:  20002015,
			StartTime: "8:00 PM",
			EndTime:   "8:15 PM",
			Active:    false,
		},
		{
			TimeCode:  20152030,
			StartTime: "8:15 PM",
			EndTime:   "8:30 PM",
			Active:    false,
		},
		{
			TimeCode:  20302045,
			StartTime: "8:30 PM",
			EndTime:   "8:45 PM",
			Active:    false,
		},
		{
			TimeCode:  20452100,
			StartTime: "8:45 PM",
			EndTime:   "9:00 PM",
			Active:    false,
		},
		{
			TimeCode:  21002115,
			StartTime: "9:00 PM",
			EndTime:   "9:15 PM",
			Active:    false,
		},
		{
			TimeCode:  21152130,
			StartTime: "9:15 PM",
			EndTime:   "9:30 PM",
			Active:    false,
		},
		{
			TimeCode:  21302145,
			StartTime: "9:30 PM",
			EndTime:   "9:45 PM",
			Active:    false,
		},
		{
			TimeCode:  21452200,
			StartTime: "9:45 PM",
			EndTime:   "10:00 PM",
			Active:    false,
		},
	}

	for dayOfWeek := 1; dayOfWeek <= 5; dayOfWeek++ {
		timeslotsWithDays := []hour.SQLHour{}

		for _, timeSlot := range timeSlots {
			newTimeslot := hour.SQLHour{
				StartTime: timeSlot.StartTime,
				EndTime:   timeSlot.EndTime,
			}

			newTimeslot.TimeCode = strconv.Itoa(timeSlot.TimeCode)
			newTimeslot.DayOfWeek = strconv.Itoa(dayOfWeek)
			if timeSlot.Active {
				newTimeslot.Active = "1"
			} else {
				newTimeslot.Active = "0"
			}

			timeslotsWithDays = append(timeslotsWithDays, newTimeslot)
		}

		err := hour.MassCreateHour(database, timeslotsWithDays)

		if err == nil {
			logger.Log(2, "database", "Seeding Data", "System", "Created all timeslots for day# "+strconv.Itoa(dayOfWeek))
		} else {
			logger.Log(3, "database", "Seeding Data", "System", err.Error())
		}
	}
}

func seedUserHours() {

	daysOfWeek := []string{"1", "3"}
	timeCodes := []string{"9150930", "9300945", "9451000", "14001415", "14151430", "14301445", "14451500"}

	for _, dayOfWeek := range daysOfWeek {
		for _, timeCode := range timeCodes {
			userHour := userHour.UserHour{}

			hour, _ := hour.GetHourByTimeCodeAndDay(database, hour.FilterHour{
				DayOfWeek: dayOfWeek,
				TimeCode:  timeCode,
			})

			users, _ := user.GetUsersByFilter(database, user.AdminViewUserInformation{UserName: "Tutor"})

			userHour.HourId = hour.Id
			userHour.TutorId, _ = strconv.Atoi(users[0].UserId)

			err := userHour.CreateUserHour(database)

			if err == nil {
				logger.Log(2, "database", "Seeding Data", "System", "Tutor: "+users[0].UserName+" User Hour is Created for "+users[0].UserName+" at "+strconv.Itoa(hour.TimeCode)+" on day#"+strconv.Itoa(hour.DayOfWeek))
			} else {
				logger.Log(3, "database", "Seeding Data", "System", err.Error())
			}
		}
	}

}

func seedMeetings() {
	var meeting meeting.Meeting
	tutor, _ := user.GetUsersByFilter(database, user.AdminViewUserInformation{UserName: "Tutor"})
	tutorHour, _ := userHour.GetUserTimeslotByFilter(database, userHour.TutorsAndHours{TutorId: tutor[0].UserId, HourId: "9300945", DayOfWeek: "1"})
	student, _ := user.GetUsersByFilter(database, user.AdminViewUserInformation{UserName: "Student"})

	meeting.StudentId, _ = strconv.Atoi(student[0].UserId)
	meeting.UserHourId, _ = strconv.Atoi(tutorHour[0].Id)
	meeting.Date = 12152024

	err := meeting.CreateMeeting(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", "Meeting Created for "+student[0].UserName+" with "+tutor[0].UserName+" on "+strconv.Itoa(meeting.Date)+" at "+tutorHour[0].HourId)
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	err = nil
	tutorHour, _ = userHour.GetUserTimeslotByFilter(database, userHour.TutorsAndHours{TutorId: tutor[0].UserId, HourId: "9300945", DayOfWeek: "3"})
	meeting.UserHourId, _ = strconv.Atoi(tutorHour[0].Id)

	meeting.Date = 12172024

	err = meeting.CreateMeeting(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", "Meeting Created for "+student[0].UserName+" with "+tutor[0].UserName+" on "+strconv.Itoa(meeting.Date)+" at "+tutorHour[0].HourId)
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

	err = nil
	tutorHour, _ = userHour.GetUserTimeslotByFilter(database, userHour.TutorsAndHours{TutorId: tutor[0].UserId, HourId: "14151430", DayOfWeek: "1"})
	meeting.UserHourId, _ = strconv.Atoi(tutorHour[0].Id)

	meeting.Date = 12152024

	err = meeting.CreateMeeting(database)

	if err == nil {
		logger.Log(2, "database", "Seeding Data", "System", "Meeting Created for "+student[0].UserName+" with "+tutor[0].UserName+" on "+strconv.Itoa(meeting.Date)+" at "+tutorHour[0].HourId)
	} else {
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}

}

func seedTopics() {
	topics := []string{"Computer Science","Algorithm Design", "Discrete Structures", "Web Development", "Computer Networking", "Computer Architecture", "Data Structures ", "Operating Systems", "System Administration"}

	for _, description := range topics {
		topic := topic.Topic{
			Description: description,
		}


		err := topic.AddTopic(database)

		if err == nil {
			logger.Log(2, "database", "Seeding Data", "System", "Topic: " +topic.Description)
		} else {
			logger.Log(3, "database", "Seeding Data", "System", err.Error())
		}

	}
}
