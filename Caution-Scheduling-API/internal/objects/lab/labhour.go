package lab

import (
	"database/sql"
	"strconv"
	"fmt"
)

type LabHour struct{
	Id int `json:"id"`
	LabId int `json:"labId"`
	HourId int `json:"hourId"`
	UserHourId int `json:"userHourId"`
}

func GetLabHours(db *sql.DB) ([]LabHour, error){
	var tempLabId string
	var tempHourId string
	var tempUserHourId string
	var tempId string

	rows, err := db.Query("SELECT Id, LabId, HoursId, TutorId FROM `labHours`")
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	labHours := []LabHour{}

	for rows.Next(){
		var lh LabHour
		if err := rows.Scan(&tempId,&tempLabId, &tempHourId, &tempUserHourId); err != nil {
			return nil, err
		}
		lh.Id, err = strconv.Atoi(tempId)
		if err != nil { return nil, err }
		lh.LabId, err = strconv.Atoi(tempLabId)
		if err != nil { return nil, err }
		lh.HourId, err = strconv.Atoi(tempHourId)
		if err != nil { return nil, err }
		lh.UserHourId, err = strconv.Atoi(tempUserHourId)
		if err != nil { return nil, err }

		labHours = append(labHours, lh)
	}
	return labHours, nil
}

func (lh *LabHour) GetLabTimeslot(db *sql.DB) error{
	var tempLabId string
	var tempHourId string
	var tempUserHourId string
	query := "SELECT Id, LabId, HoursId, TutorId FROM `labHours` WHERE `Id` = "+ strconv.Itoa(lh.Id) +";	"
	err := db.QueryRow(query).Scan(&tempLabId, &tempHourId, &tempUserHourId)
	if err != nil { return err }
	lh.LabId, err = strconv.Atoi(tempLabId)
	if err != nil { return err }
	lh.HourId, err = strconv.Atoi(tempHourId)
	if err != nil { return err }
	lh.UserHourId, err = strconv.Atoi(tempUserHourId)
	fmt.Println(query)
	return err
}

func (lh *LabHour) CreateLabTimeSlot(db *sql.DB) error{
	query := "INSERT INTO `labHours` (`LabId`, `HoursId`, `TutorId`) VALUES ('"+strconv.Itoa(lh.LabId)+"', '"+strconv.Itoa(lh.HourId)+"', '"+strconv.Itoa(lh.UserHourId)+"');"
	err := db.QueryRow(query)
	fmt.Println(query)
	if err != nil{
		return err.Err()
	}
	return nil
}

func (lh *LabHour) DeleteLabTimeSlot(db *sql.DB) error{
	query := "DELETE FROM `labHours` WHERE `Id` = '"+strconv.Itoa(lh.LabId)+"'"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}

func (lh *LabHour) ChangeTutor(db *sql.DB) error{
	query := "UPDATE `labHours` SET `TutorId` = '"+strconv.Itoa(lh.UserHourId) +"'WHERE `Id` = '"+strconv.Itoa(lh.LabId)+"'"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}

