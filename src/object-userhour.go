package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

type userHour struct {
	Id        int    `json:"id"`
	HourId    int    `json:"hourId"`
	Tutor     string `json:"tutor"`
	Available bool   `json:"available"`
}

func (uh *userHour) getUserHour(db *sql.DB) error{
	var tempHourId string
	var tempAvailable string
	query := "SELECT `Id`, `hourId`, `username`, `available` FROM `userHours` WHERE `Id` = '"+strconv.Itoa(uh.Id) +"'"
	err := db.QueryRow(query).Scan(&uh.Id, &tempHourId, &uh.Tutor, &tempAvailable)
	uh.HourId, err = strconv.Atoi(tempHourId)
	uh.Available, err = strconv.ParseBool(tempAvailable)
	fmt.Println(query)
	return err
}

