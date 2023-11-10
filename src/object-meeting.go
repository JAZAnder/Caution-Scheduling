package main

import (
	"database/sql"
	"errors"
	"strconv"
)

type meeting struct{
	Id int `json:"id"`
	UserHourId int `json:"userHourId"`
	LabId int `json:"labId"`
	StudentName string `json:"studentName"`
	StudentEmail string `json:"studentEmail"`
}

func (m *meeting) getMeeting(db *sql.DB) error{
	var tempTutorHourId string
	var tempLabId string
	query := "SELECT userHourId, labId, studentName, studentEmail FROM meetings WHERE id=" + strconv.Itoa(m.Id)
	err := db.QueryRow(query).Scan(&tempTutorHourId, &tempLabId, &m.StudentName, &m.StudentEmail)
	m.UserHourId, err = strconv.Atoi(tempTutorHourId)
	m.LabId, err = strconv.Atoi(tempLabId)
	return err
}



func (m *meeting) updateMeeting(db *sql.DB) error{
	query := "Update `meetings` SET `userHourId` = '"+strconv.Itoa(m.UserHourId)+"', `labId` = '"+strconv.Itoa(m.LabId)+"', `studentName` = '"+m.StudentName+"', `studentEmail` = '"+m.StudentEmail+"' WHERE `meetings`.`id` ="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}

func (m *meeting) createMeeting(db *sql.DB) error{
	var available bool
	query := "SELECT `Available` FROM `UserHours` WHERE `userHourId` = " + strconv.Itoa(m.UserHourId)
	err := db.QueryRow(query).Scan(&available)

	if err != nil{
		return err
	}
	
	if !available{
		return errors.New("This tutor is not available for that given time")
	}

	query = "INSERT INTO `meetings` (`userHourId`, `labId`, `studentName`, `studentEmail`) VALUES ('"+strconv.Itoa(m.UserHourId)+"', '"+strconv.Itoa(m.LabId)+"', '"+m.StudentName+"', '"+m.StudentEmail+"');"
	errsql := db.QueryRow(query)

	if errsql.Err() != nil{
		return errsql.Err()
	}

	query = "UPDATE `UserHours` SET `Available` = 'false' WHERE `userHourId` = " + strconv.Itoa(m.UserHourId)
	errsql = db.QueryRow(query)

	if errsql.Err() != nil{
		return errsql.Err()
	}

	return nil
}

func getMeetings(db *sql.DB) ([]meeting, error){
	rows, err := db.Query("SELECT id, userHourId, labId, studentName, studentEmail FROM meetings")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	meetings := []meeting{}

	for rows.Next(){
		var m meeting
		if err := rows.Scan(&m.Id, &m.UserHourId, &m.LabId, &m.StudentName, &m.StudentEmail); err != nil{
			return nil, err
		}
		meetings = append(meetings, m)
	}
	return meetings, nil
}

// func getMyMeetings(db *sql.DB) ([]meeting, error){
// 	var c sessionCookie
//     cookie, err := r.Cookie("key")
//     if err != nil {
//         if errors.Is(err, http.ErrNoCookie) {
//             return
//         } else {
//             return
//         }
//     }

//     c.Cookie = cookie.Value

//     currentUser, err := c.checkSession(a.DB)
//     if err != nil {
//         if err == sql.ErrNoRows {
//             respondWithError(w, http.StatusUnauthorized, "Session Expired")
//             return
//         } else {
//             respondWithError(w, http.StatusInternalServerError, err.Error())
//             return
//         }
//     } 

// 	rows, err := db.Query("SELECT id, userHourId, labId, studentName, studentEmail FROM meetings m join userHours u on m.userHourId = u.id where u.userId = " + currentUser)

// 	if err != nil{
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	meetings := []meeting{}

// 	for rows.Next(){
// 		var m meeting
// 		if err := rows.Scan(&m.Id, &m.UserHourId, &m.LabId, &m.StudentName, &m.StudentEmail); err != nil{
// 			return nil, err
// 		}
// 		meetings = append(meetings, m)
// 	}
// 	return meetings, nil
// }

func (m *meeting) deleteMeeting(db *sql.DB) error{
	query := "DELETE FROM `meetings` WHERE `meetings`.`Id`="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}