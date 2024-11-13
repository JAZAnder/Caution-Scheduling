package meeting

import (
	"database/sql"
	"fmt"
	//"errors"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"

)

// func (m *Meeting) GetMeeting(db *sql.DB) error{
// 	var tempTutorHourId string
// 	var tempLabId string
// 	query := "SELECT tutorHourId, labId, studentName, studentEmail FROM meetings WHERE id=" + strconv.Itoa(m.Id)
// 	err := db.QueryRow(query).Scan(&tempTutorHourId, &tempLabId, &m.StudentName, &m.StudentEmail)
// 	if err != nil { return err }
// 	m.UserHourId, err = strconv.Atoi(tempTutorHourId)
// 	if err != nil { return err }
// 	m.LabId, err = strconv.Atoi(tempLabId)
// 	return err
// }

// func (m *Meeting) UpdateMeeting(db *sql.DB) error {
// 	query := "Update `meetings` SET `userHourId` = '" + strconv.Itoa(m.UserHourId) + "', `labId` = '" + strconv.Itoa(m.LabId) + "', `studentName` = '" + m.StudentName + "', `studentEmail` = '" + m.StudentEmail + "' WHERE `meetings`.`id` =" + strconv.Itoa(m.Id) + ""
// 	_, err := db.Exec(query)
// 	return err
// }

func (m *Meeting) CreateMeeting(db *sql.DB) error {

	query := "INSERT INTO `meetings` (`topicId`,`tutorHourId`, `studentId`, `date`) VALUES ('" + strconv.Itoa(m.TopicId) + "','" + strconv.Itoa(m.UserHourId) + "', '" + strconv.Itoa(m.StudentId) + "', '" + strconv.Itoa(m.Date) + "');"
	fmt.Print(query)
	errSql := db.QueryRow(query)

	if errSql.Err() != nil {
		return errSql.Err()
	}

	return nil
}

// func GetMeetings(db *sql.DB) ([]Meeting, error) {
// 	rows, err := db.Query("SELECT `Id`, `tutorHourId`, `labId`, `studentName`, `studentEmail`, `date` FROM `meetings`")

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	meetings := []Meeting{}

// 	for rows.Next() {
// 		var tempId string
// 		var tempuserHourId string
// 		var labId string
// 		var date string

// 		var m Meeting
// 		if err := rows.Scan(&tempId, &tempuserHourId, &labId, &m.StudentName, &m.StudentEmail, &date); err != nil {
// 			return nil, err
// 		}
// 		m.Id, err = strconv.Atoi(tempId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		m.UserHourId, err = strconv.Atoi(tempuserHourId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		m.LabId, err = strconv.Atoi(labId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		m.Date, err = strconv.Atoi(date)
// 		if err != nil {
// 			return nil, err
// 		}

// 		meetings = append(meetings, m)
// 	}
// 	return meetings, nil
// }

func GetMyMeetings(db *sql.DB, userId int) ([]BasicMeetingDto, error) {
	query :=  "SELECT  m.Id, m.date, topic.Id, topic.description, h.Id, h.startTime, h.endTime, tutor.Id, tutor.firstName, tutor.lastName, tutor.email, student.Id, student.firstName, student.lastName, student.email " +
			"FROM meetings m " +
			"join userHours uh on m.tutorHourId = uh.id  " +
			"join localusers tutor on uh.userId = tutor.Id " +
			"join hours h on uh.hourId = h.Id " +
			"join localusers student on m.studentId = student.Id " +
			"left join topic on m.topicId = topic.Id " +
			"Where student.Id = " + strconv.Itoa(userId) + " or tutor.Id = " + strconv.Itoa(userId) +
			";"


			logger.Log(1, "Database", "GetMeetings", "System", query)

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	meetings := []BasicMeetingDto{}

	for rows.Next() {
		var m BasicMeetingDto
		if err := rows.Scan(&m.Id, &m.Date, &m.Topic.Id, &m.Topic.Description, &m.Hour.Id, &m.Hour.StartTime, &m.Hour.EndTime, &m.Tutor.Id, &m.Tutor.FirstName, &m.Tutor.LastName, &m.Tutor.Email, &m.Student.Id, &m.Student.FirstName, &m.Student.LastName, &m.Student.Email); err != nil {
			return nil, err
		}
		meetings = append(meetings, m)
	}
	return meetings, nil
}

// func (m *Meeting) DeleteMeeting(db *sql.DB) error {
// 	query := "DELETE FROM `meetings` WHERE `meetings`.`Id`=" + strconv.Itoa(m.Id) + ""
// 	_, err := db.Exec(query)
// 	return err
// }
