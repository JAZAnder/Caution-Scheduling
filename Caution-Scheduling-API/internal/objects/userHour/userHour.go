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
	err := db.QueryRow(query).Scan(&tempHourId, &uh.TutorId, &tempAvailable)
	if err != nil {
		return err
	}
	uh.HourId, err = strconv.Atoi(tempHourId)
	if err != nil {
		return err
	}
	uh.Available, err = strconv.ParseBool(tempAvailable)
	fmt.Println(query)
	return err
}

func (uh *UserHour) GetUserHourId(db *sql.DB) error {
	var tempId string
	query := "SELECT `Id` FROM `userHours` WHERE `username` = '" + strconv.Itoa(uh.TutorId) + "' AND `hourId` = '" + strconv.Itoa(uh.HourId) + "'"
	err := db.QueryRow(query).Scan(&tempId)
	if err != nil {
		return err
	}
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
	query := "INSERT INTO `userHours` (`hourId`,`userId`) VALUES ('" + strconv.Itoa(uh.HourId) + "','" + strconv.Itoa(uh.TutorId) + "')"
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

func GetUserTimeslotByFilter(db *sql.DB, filter TutorsAndHours) ([]TutorsAndHours, error) {
	rows, err := db.Query("SELECT uh.id, lu.Id As `userId`,lu.firstName, lu.lastName, h.Id As `hourId`, h.startTime, h.endTime, h.dayOfWeek  FROM userHours `uh` inner join localusers `lu` on uh.userId = lu.Id inner join hours `h` on uh.hourId = h.Id where lu.id = '" + filter.TutorId + "' OR h.id = '" + filter.HourId + "' OR h.dayOfWeek = '" + filter.DayOfWeek + "';")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filteredResults := []TutorsAndHours{}

	for rows.Next() {
		var result TutorsAndHours
		if err := rows.Scan(&result.Id, &result.TutorId, &result.FirstName, &result.LastName, &result.HourId, &result.StartTime, &result.EndTime, &result.DayOfWeek); err != nil {
			return nil, err
		}

		filteredResults = append(filteredResults, result)
	}

	return filteredResults, nil
}


func GetAvailableHoursByUserAndDay(db *sql.DB, userId int, date int)([]TutorsAndHours, error){
	query := "SELECT userHours.id, userHours.userId,localusers.firstName, localusers.lastName, userHours.hourId, hours.startTime, hours.endTime, hours.dayOfWeek " +
		" FROM userHours Join hours on userHours.hourId = hours.Id join localusers on userHours.userId = localusers.Id " +
		" Where userId = "+strconv.Itoa(userId)+" And userHours.id NOT IN ( " +
				" Select userHours.id" +
				" From meetings join userHours on meetings.tutorHourId = userHours.id " +
				" where userHours.userId = "+strconv.Itoa(userId)+" AND date = "+strconv.Itoa(date)+");"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filteredResults := []TutorsAndHours{}

	for rows.Next() {
		var result TutorsAndHours
		if err := rows.Scan(&result.Id, &result.TutorId, &result.FirstName, &result.LastName, &result.HourId, &result.StartTime, &result.EndTime, &result.DayOfWeek); err != nil {
			return nil, err
		}

		filteredResults = append(filteredResults, result)
	}

	return filteredResults, nil
}

func (uh *UserHour) GetHoursByUserId(db *sql.DB) ([]UserHour, error) {
	rows, err := db.Query("SELECT `Id`, `hourId`, `userId`, `available` FROM `userHours` WHERE `userId` = '" + strconv.Itoa(uh.TutorId) + "'")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []UserHour{}
	var tempAva int
	for rows.Next() {
		var uh UserHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.TutorId, &tempAva); err != nil {
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
	rows, err := db.Query("SELECT `Id`, `hourId`, `username` FROM `userHours` WHERE `username` = '" + strconv.Itoa(uh.TutorId) + "' AND `available` = 1;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userHours := []UserHour{}

	for rows.Next() {
		var uh UserHour
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.TutorId); err != nil {
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
		if err := rows.Scan(&uh.Id, &uh.HourId, &uh.TutorId, &uh.Available); err != nil {
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
		if err := rows.Scan(&uh.Id, &uh.HourId ,&uh.TutorId); err != nil {
			return nil, err
		}
		uh.Available = true
		userHours = append(userHours, uh)
	}
	return userHours, nil
}
