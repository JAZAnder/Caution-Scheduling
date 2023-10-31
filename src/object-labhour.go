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
	query := "SELECT "
	err := db.QueryRow(query).Scan(&tempLabId, &tempHourId, &tempUserHourId)
	lh.LabId, err = strconv.Atoi(tempLabId)
	lh.HourId, err = strconv.Atoi(tempHourId)
	lh.UserHourId, err = strconv.Atoi(tempUserHourId)
	return err
}