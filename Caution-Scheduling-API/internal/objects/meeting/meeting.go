package meeting

import (
	"database/sql"
	"fmt"
	//"errors"
	"strconv"
)

type Meeting struct{
	Id int `json:"id"`
	UserHourId int `json:"userHourId"`
	LabId int `json:"labId"`
	StudentName string `json:"studentName"`
	StudentEmail string `json:"studentEmail"`
	Date int `json:"date"`
}

func (m *Meeting) GetMeeting(db *sql.DB) error{
	var tempTutorHourId string
	var tempLabId string
	query := "SELECT tutorHourId, labId, studentName, studentEmail FROM meetings WHERE id=" + strconv.Itoa(m.Id)
	err := db.QueryRow(query).Scan(&tempTutorHourId, &tempLabId, &m.StudentName, &m.StudentEmail)
	if err != nil { return err }
	m.UserHourId, err = strconv.Atoi(tempTutorHourId)
	if err != nil { return err }
	m.LabId, err = strconv.Atoi(tempLabId)
	return err
}



func (m *Meeting) UpdateMeeting(db *sql.DB) error{
	query := "Update `meetings` SET `userHourId` = '"+strconv.Itoa(m.UserHourId)+"', `labId` = '"+strconv.Itoa(m.LabId)+"', `studentName` = '"+m.StudentName+"', `studentEmail` = '"+m.StudentEmail+"' WHERE `meetings`.`id` ="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}

func (m *Meeting) CreateMeeting(db *sql.DB) error{
	// var available bool
	// query := "SELECT `Available` FROM `UserHours` WHERE `userHourId` = " + strconv.Itoa(m.UserHourId)
	// err := db.QueryRow(query).Scan(&available)

	// if err != nil{
	// 	return err
	// }
	
	// if !available{
	// 	return errors.New("This tutor is not available for that given time")
	// }

	query := "INSERT INTO `meetings` (`tutorHourId`, `labId`, `studentName`, `studentEmail`, `date`) VALUES ('"+strconv.Itoa(m.UserHourId)+"', '"+strconv.Itoa(m.LabId)+"', '"+m.StudentName+"', '"+m.StudentEmail+"', "+strconv.Itoa(m.Date)+");"
	fmt.Print(query)
	errsql := db.QueryRow(query)

	if errsql.Err() != nil{
		return errsql.Err()
	}

	// query = "UPDATE `UserHours` SET `Available` = 'false' WHERE `userHourId` = " + strconv.Itoa(m.UserHourId)
	// errsql = db.QueryRow(query)

	// if errsql.Err() != nil{
	// 	return errsql.Err()
	// }

	return nil
}

func GetMeetings(db *sql.DB) ([]Meeting, error){
	rows, err := db.Query("SELECT `Id`, `tutorHourId`, `labId`, `studentName`, `studentEmail`, `date` FROM `meetings`")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	meetings := []Meeting{}

	for rows.Next(){
		var tempId string
		var tempuserHourId string
		var labId string
		var date string

		var m Meeting
		if err := rows.Scan(&tempId, &tempuserHourId, &labId, &m.StudentName, &m.StudentEmail, &date); err != nil{
			return nil, err
		}
		m.Id, err = strconv.Atoi(tempId)
		if err != nil { return nil, err }
		m.UserHourId, err = strconv.Atoi(tempuserHourId)
		if err != nil { return nil, err }
		m.LabId, err = strconv.Atoi(labId)
		if err != nil { return nil, err }
		m.Date, err = strconv.Atoi(date)
		if err != nil { return nil, err }

		meetings = append(meetings, m)
	}
	return meetings, nil
}

func GetMyMeetings(db *sql.DB, userName string) ([]Meeting, error){
	rows, err := db.Query("SELECT m.Id, m.tutorHourId, m.labId, m.studentName, m.studentEmail, m.date FROM meetings m JOIN userHours u ON m.tutorHourId = u.Id WHERE u.username ='" + userName + "'")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	meetings := []Meeting{}

	for rows.Next(){
		var m Meeting
		if err := rows.Scan(&m.Id, &m.UserHourId, &m.LabId, &m.StudentName, &m.StudentEmail, &m.Date); err != nil{
			return nil, err
		}
		meetings = append(meetings, m)
	}
	return meetings, nil
}




func (m *Meeting) DeleteMeeting(db *sql.DB) error{
	query := "DELETE FROM `meetings` WHERE `meetings`.`Id`="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}