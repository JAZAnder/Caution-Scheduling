package main

import (
	"database/sql"
	"strconv"
)

type labHour struct{
	Id int `json:"id"`
	LabId int `json:"labId"`
	HourId int `json:"hoursId"`
	UserHourId int `json:"userHourId"`
}

func (lh *labHour) getLabTimeslot (db *sql.DB) error{
	var tempLabId string
	var tempHourId string
	var tempUserHourId string
	query := "SELECT Id, LabId, HoursId, TutorId FROM `labHours` WHERE `Id` = "+ strconv.Itoa(lh.Id) +";	"
	err := db.QueryRow(query).Scan(&tempLabId, &tempHourId, &tempUserHourId)
	lh.LabId, err = strconv.Atoi(tempLabId)
	lh.HourId, err = strconv.Atoi(tempHourId)
	lh.UserHourId, err = strconv.Atoi(tempUserHourId)
	return err
}

func (lh *labHour) createLabTimeSlot (db *sql.DB) error{
	query := "INSERT INTO `labHours` (`LabId`, `HoursId`, `TutorId`) VALUES ('"+strconv.Itoa(lh.LabId)+"', '"+strconv.Itoa(lh.HourId)+"', '"+strconv.Itoa(lh.UserHourId)+"');"
	err := db.QueryRow(query)
	if err != nil{
		return err.Err()
	}
	return nil
}

// func (lh *labHour) deleteLabTimeSlot (db *sql.DB) error{

// }

// func (lh *labHour) changeTutor (db *sql.DB) error{

// }

