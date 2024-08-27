package lab

import (
	"database/sql"
	"strconv"
	"fmt"
)

type labHour struct{
	Id int `json:"id"`
	LabId int `json:"labId"`
	HourId int `json:"hourId"`
	UserHourId int `json:"userHourId"`
}

func getLabHous(db *sql.DB) ([]labHour, error){
	var tempLabId string
	var tempHourId string
	var tempUserHourId string
	var tempId string

	rows, err := db.Query("SELECT Id, LabId, HoursId, TutorId FROM `labHours`")
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	labHours := []labHour{}

	for rows.Next(){
		var lh labHour
		if err := rows.Scan(&tempId,&tempLabId, &tempHourId, &tempUserHourId); err != nil {
			return nil, err
		}
		lh.Id, err = strconv.Atoi(tempId)
		lh.LabId, err = strconv.Atoi(tempLabId)
		lh.HourId, err = strconv.Atoi(tempHourId)
		lh.UserHourId, err = strconv.Atoi(tempUserHourId)

		labHours = append(labHours, lh)
	}
	return labHours, nil
}

func (lh *labHour) getLabTimeslot(db *sql.DB) error{
	var tempLabId string
	var tempHourId string
	var tempUserHourId string
	query := "SELECT Id, LabId, HoursId, TutorId FROM `labHours` WHERE `Id` = "+ strconv.Itoa(lh.Id) +";	"
	err := db.QueryRow(query).Scan(&tempLabId, &tempHourId, &tempUserHourId)
	lh.LabId, err = strconv.Atoi(tempLabId)
	lh.HourId, err = strconv.Atoi(tempHourId)
	lh.UserHourId, err = strconv.Atoi(tempUserHourId)
	fmt.Println(query)
	return err
}

func (lh *labHour) createLabTimeSlot(db *sql.DB) error{
	query := "INSERT INTO `labHours` (`LabId`, `HoursId`, `TutorId`) VALUES ('"+strconv.Itoa(lh.LabId)+"', '"+strconv.Itoa(lh.HourId)+"', '"+strconv.Itoa(lh.UserHourId)+"');"
	err := db.QueryRow(query)
	fmt.Println(query)
	if err != nil{
		return err.Err()
	}
	return nil
}

func (lh *labHour) deleteLabTimeSlot(db *sql.DB) error{
	query := "DELETE FROM `labHours` WHERE `Id` = '"+strconv.Itoa(lh.LabId)+"'"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}

func (lh *labHour) changeTutor(db *sql.DB) error{
	query := "UPDATE `labHours` SET `TutorId` = '"+strconv.Itoa(lh.UserHourId) +"'WHERE `Id` = '"+strconv.Itoa(lh.LabId)+"'"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}

