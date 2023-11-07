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

func (uh *userHour) makeUnavailabe(db *sql.DB)error{
	query := "UPDATE `userHours` SET `available` = `0` WHERE `Id` = '"+strconv.Itoa(uh.Id)+"'"
	_, err := db.Exec(query)
	fmt.Println(query)
	uh.Available = false
	return err
} 

func (uh *userHour) createUserHour(db *sql.DB)error{
	query := "INSERT INTO `userhours` (`hourId`,`username`) VALUES ('"+strconv.Itoa(uh.HourId)+"','"+uh.Tutor+"')"
	err := db.QueryRow(query)
	fmt.Println(query)
	if err != nil{
		return err.Err()
	}
	return nil
}

func (uh *userHour) deleteUserHour(db *sql.DB) error{
	query := "DELETE FROM `userhours` WHERE `Id` = '"+strconv.Itoa(uh.Id)+"'"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}