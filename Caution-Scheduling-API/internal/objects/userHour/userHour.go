package userHour

import (
	"database/sql"
	"fmt"
	"strconv"
)


func (uh *UserHour) GetUserHour(db *sql.DB) error {
	var tempHourId string
	var tempAvailable string
	query := "SELECT `hourId`, `username`, `available` FROM `userHours` WHERE `Id` = '" + strconv.Itoa(uh.Id) + "'"
	err := db.QueryRow(query).Scan(&tempHourId, &uh.Tutor, &tempAvailable)
	if err != nil { return err}
	uh.HourId, err = strconv.Atoi(tempHourId)
	if err != nil { return err}
	uh.Available, err = strconv.ParseBool(tempAvailable)
	fmt.Println(query)
	return err
}

func (uh *UserHour) GetUserHourId(db *sql.DB) error {
	var tempId string
	query := "SELECT `Id` FROM `userHours` WHERE `username` = '" + uh.Tutor + "' AND `hourId` = '" + strconv.Itoa(uh.HourId) + "'"
	err := db.QueryRow(query).Scan(&tempId)
	if err != nil { return err}
	uh.Id, err = strconv.Atoi(tempId)
	return err
}

func (uh *UserHour) MakeUnavailabe(db *sql.DB) error {
	query := "UPDATE `userHours` SET `available` = `0` WHERE `Id` = '" + strconv.Itoa(uh.Id) + "'"
	_, err := db.Exec(query)
	fmt.Println(query)
	uh.Available = false
	return err
}

func (uh *UserHour) CreateUserHour(db *sql.DB) error {
	query := "INSERT INTO `userHours` (`hourId`,`username`) VALUES ('" + strconv.Itoa(uh.HourId) + "','" + uh.Tutor + "')"
	err := db.QueryRow(query)
	fmt.Println(query)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (uh *UserHour) DeleteUserHourById(db *sql.DB) error {
	query := "DELETE FROM `userHours` WHERE `Id` = '" + strconv.Itoa(uh.Id) + "';"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}

func (uh *UserHour) DeleteUserHour(db *sql.DB) error {
	query := "DELETE FROM `userHours` WHERE `hourId` = '" + strconv.Itoa(uh.HourId) + "' AND `username` = '" + uh.Tutor + "'"
	_, err := db.Exec(query)
	fmt.Println(query)
	return err
}

func (uh *UserHour) GetHours(db *sql.DB) ([]UserHour, error) {
	rows, err := db.Query("SELECT `Id`, `hourId`, `username`, `available` FROM `userHours` WHERE `username` = '" + uh.Tutor + "'")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []UserHour{}
	var tempAva int
	for rows.Next() {
		var uh UserHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.Tutor, &tempAva); err != nil {
			return nil, err
		}
		if tempAva == 1 {
			uh.Available = true
		} else {
			uh.Available = false
		}
		userHours = append(userHours, uh)
	}
	return userHours, nil
}

func (uh *UserHour) GetAvailableHours(db *sql.DB) ([]UserHour, error) {
	rows, err := db.Query("SELECT `Id`, `hourId`, `username` FROM `userHours` WHERE `username` = '" + uh.Tutor + "' AND `available` = 1;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []UserHour{}

	for rows.Next() {
		var uh UserHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.Tutor); err != nil {
			return nil, err
		}
		uh.Available = true
		userHours = append(userHours, uh)
	}
	return userHours, nil
}

func GetUserHours(db *sql.DB) ([]UserHour, error) {
	rows, err := db.Query("SELECT `Id`, `hourId`, `username`, `available` FROM `userHours`")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []UserHour{}

	for rows.Next() {
		var uh UserHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.Tutor, &uh.Available); err != nil {
			return nil, err
		}
		userHours = append(userHours, uh)
	}
	return userHours, nil
}

func GetUsersByHour(db *sql.DB, hourId int)([]UserHour, error){
	rows, err := db.Query("SELECT `Id`, `hourId`, `username` FROM `userHours` WHERE `hourId` = '" + strconv.Itoa(hourId)  + "' AND `available` = 1;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []UserHour{}

	for rows.Next() {
		var uh UserHour
		if err := rows.Scan(&uh.Id, &uh.HourId ,&uh.Tutor); err != nil {
			return nil, err
		}
		uh.Available = true
		userHours = append(userHours, uh)
	}
	return userHours, nil
}
