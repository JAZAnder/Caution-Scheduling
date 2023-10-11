package main

import(
	"database/sql"
	"strconv"
)

type meeting struct{
	Id int `json:"id"`
	TutorHourId int `json:"tutorHourId"`
	LabId int `json:"labId"`
	StudentName string `json:"studentName"`
	StudentEmail string `json:"studentEmail"`
}

func (m *meeting) getMeeting(db *sql.DB) error{
	query := "SELECT tutorHourId, labId, studentName, studentEmail FROM meetings WHERE id=" + strconv.Itoa(m.Id)
	return db.QueryRow(query).Scan(&m.tutorHourId, &m.labId, &m.studentName, &m.studentEmail)
}

func (m *meeting) updateMeeting(db *sql.DB) error{
	query := "Update `meetings` SET `tutorHourId` = '"+m.tutorHourId+"', `labId` = '"+m.labId+"', `studentName` = '"+m.studentName+"', `studentEmail` = '"+studentEmail+"' WHERE `meetings`.`id` ="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}

func (m *meeting) createMeeting(db *sql.DB) error{
	query := "INSERT INTO `meetings` (`tutorHourId`, `labId`, `studentName`, `studentEmail`) VALUES ('"+m.tutorHourId+"', '"+m.labId+"', '"+m.studentName+"', '"+m.studentEmail"')"
	err := db.QueryRow(query)

	if err != nil{
		return err.Err()
	}
	return nil
}

func getMeetings(db *sql.DB) ([]meeting, error){
	rows, err := db.Query("SELECT id, tutorHourId, labId, studentName, studentEmail FROM meetings")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	meetings := []meeting{}

	for rows.Next(){
		var m meeting
		if err := rows.Scan(&m.Id, &m.TutorHourId, &m.LabId, &m.StudentName, &m.StudentEmail); err != nil{
			return nil, err
		}
		meetings = append(meetings, m)
	}
	return meetings, nil
}

func (m *meeting) deleteMeeting(db *sql.DB) error{
	query := "DELETE FROM `meetings` WHERE `meetings`.`Id`="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}