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


func (uh *userHour) deleteUserHourById(db *sql.DB) error{
	query := "DELETE FROM `userhours` WHERE `Id` = '"+strconv.Itoa(uh.Id)+"';"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}


func (uh *userHour) deleteUserHour(db *sql.DB) error{
	query := "DELETE FROM `userhours` WHERE `hourId` = '"+strconv.Itoa(uh.HourId)+"' AND `username` = '"+uh.Tutor+"'"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}

func (uh *userHour) getHours(db *sql.DB) ([]userHour, error){
	rows, err := db.Query("SELECT `Id`, `hourId`, `username`, `available` FROM `userHours` WHERE `username` = '"+uh.Tutor+"'")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []userHour{}
	var tempAva int
	for rows.Next(){
		var uh userHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.Tutor, &tempAva); err != nil{
			return nil, err
		}
		if(tempAva == 1){
			uh.Available = true
		}else{
			uh.Available = false
		}
		userHours = append(userHours, uh)
	}
	return userHours, nil
}

func (uh *userHour) getAvailableHours(db *sql.DB) ([]userHour, error){
	rows, err := db.Query("SELECT `Id`, `hourId`, `username` FROM `userHours` WHERE `username` = '"+uh.Tutor+"' AND `available` = 1;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []userHour{}

	for rows.Next(){
		var uh userHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.Tutor); err != nil{
			return nil, err
		}
		uh.Available = true
		userHours = append(userHours, uh)
	}
	return userHours, nil
}



func getUserHours(db *sql.DB) ([]userHour, error){
	rows, err := db.Query("SELECT `Id`, `hourId`, `username`, `available` FROM `userHours`")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []userHour{}

	for rows.Next(){
		var uh userHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.Tutor, &uh.Available); err != nil{
			return nil, err
		}
		userHours = append(userHours, uh)
	}
	return userHours, nil
}