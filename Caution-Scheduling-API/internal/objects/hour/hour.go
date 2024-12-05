package hour

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"

)

func (h *Hour) toPrettyHour(db *sql.DB) PrettyHour {
	ph := PrettyHour{
		Id: h.Id,
		TimeCode: h.TimeCode,
		StartTime: h.StartTime,
		EndTime: h.EndTime,
		Active: h.Active,
	}

	if(h.DayOfWeek == 1){
		ph.DayOfWeek = "Monday"
	}else if(h.DayOfWeek == 2){
		ph.DayOfWeek = "Tuesday"
	}else if(h.DayOfWeek == 3){
		ph.DayOfWeek = "Wednesday"
	}else if(h.DayOfWeek == 4){
		ph.DayOfWeek = "Thursday"
	}else if(h.DayOfWeek == 5){
		ph.DayOfWeek = "Friday"
	}else if(h.DayOfWeek == 6){
		ph.DayOfWeek = "Saturday"
	}else if(h.DayOfWeek == 0){
		ph.DayOfWeek = "Sunday"
	}
	return ph
}

func (multiDay TimeslotsMultiDay) ToTimeslotArray() []Hour {
	timeslots := []Hour{}

	if multiDay.Monday {
		timeslots = append(timeslots, Hour{
			StartTime: multiDay.StartTime,
			EndTime:   multiDay.EndTime,
			DayOfWeek: 1,
		})
	}
	if multiDay.Tuesday {
		timeslots = append(timeslots, Hour{
			StartTime: multiDay.StartTime,
			EndTime:   multiDay.EndTime,
			DayOfWeek: 2,
		})

	}
	if multiDay.Wednesday {
		timeslots = append(timeslots, Hour{
			StartTime: multiDay.StartTime,
			EndTime:   multiDay.EndTime,
			DayOfWeek: 3,
		})

	}
	if multiDay.Thursday {
		timeslots = append(timeslots, Hour{
			StartTime: multiDay.StartTime,
			EndTime:   multiDay.EndTime,
			DayOfWeek: 4,
		})

	}
	if multiDay.Friday {
		timeslots = append(timeslots, Hour{
			StartTime: multiDay.StartTime,
			EndTime:   multiDay.EndTime,
			DayOfWeek: 5,
		})

	}
	if multiDay.Saturday {
		timeslots = append(timeslots, Hour{
			StartTime: multiDay.StartTime,
			EndTime:   multiDay.EndTime,
			DayOfWeek: 6,
		})

	}
	if multiDay.Sunday {
		timeslots = append(timeslots, Hour{
			StartTime: multiDay.StartTime,
			EndTime:   multiDay.EndTime,
			DayOfWeek: 0,
		})
	}

	return timeslots

}

func (h *Hour) GetHour(db *sql.DB) error {
	var tempDayOfWeek string
	query := "SELECT startTime, endTime, dayOfWeek  FROM hours WHERE id=" + strconv.Itoa(h.Id)
	err := db.QueryRow(query).Scan(&h.StartTime, &h.EndTime, &tempDayOfWeek)

	h.DayOfWeek, _ = strconv.Atoi(tempDayOfWeek)
	return err
}

func GetHourByTimeCodeAndDay(db *sql.DB, filter FilterHour) (Hour, error) {
	var result Hour
	var err error

	result.DayOfWeek, err = strconv.Atoi(filter.DayOfWeek)
	if err != nil { return Hour{}, err }

	result.TimeCode, err = strconv.Atoi(filter.TimeCode)
	if err != nil { return Hour{}, err }

	query := "SELECT Id, startTime, endTime FROM hours WHERE timeCode=" + filter.TimeCode + " AND dayOfWeek = " + filter.DayOfWeek + ";"
	err = db.QueryRow(query).Scan(&result.Id, &result.StartTime, &result.EndTime)
	if err != nil { return Hour{}, err }

	return result, nil
}

func (h *Hour) UpdateHour(db *sql.DB) error {
	query := "UPDATE `hours` SET `startTime` = '" + h.StartTime + "', `endTime` = '" + h.EndTime + "' WHERE `hours`.`id` = " + strconv.Itoa(h.Id) + ""
	_, err := db.Exec(query)
	return err
}

func (h *Hour) MakeHourActive(db *sql.DB) error {
	query := "UPDATE `hours` SET `active` = '1' WHERE `hours`.`id` = " + strconv.Itoa(h.Id) + ""
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return h.GetHour(db)
}

func (h *Hour) MakeHourDisable(db *sql.DB) error {
	query := "UPDATE `hours` SET `active` = '0' WHERE `hours`.`id` = " + strconv.Itoa(h.Id) + ""
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return h.GetHour(db)
}

func (h *Hour) CreateHour(db *sql.DB) error {
	query := "INSERT INTO `hours` (`timeCode`,`startTime`, `endTime`, `dayOfWeek`) VALUES ('" + strconv.Itoa(h.TimeCode) + "','" + h.StartTime + "','" + h.EndTime + "','" + strconv.Itoa(h.DayOfWeek) + "')"
	err := db.QueryRow(query)

	if err != nil {
		return err.Err()
	}
	return nil
}

func MassCreateHour(db *sql.DB, timeslots []SQLHour) error {

	query := "INSERT INTO `hours` (`timeCode`,`startTime`, `endTime`, `dayOfWeek`, `active`) VALUES "

	for _, timeslot := range timeslots {
		sqlString := "('" + timeslot.TimeCode + "','" + timeslot.StartTime + "','" + timeslot.EndTime + "','" + timeslot.DayOfWeek + "', '"+timeslot.Active+"'),"
		query = query + sqlString
	}
	query = strings.TrimRight(query, ",")
	query = query + ";"

	logger.Log(1, "database", "Adding Timeslots", "System", query)

	err := db.QueryRow(query)

	if err != nil {
		return err.Err()
	}
	return nil
}

func GetHours(db *sql.DB) ([]PrettyHour, error) {
	rows, err := db.Query("SELECT id, startTime, endTime, dayOfWeek, active FROM hours")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hours := []PrettyHour{}

	for rows.Next() {
		var h Hour
		err := rows.Scan(&h.Id, &h.StartTime, &h.EndTime, &h.DayOfWeek, &h.Active)
		if err != nil {
			return nil, err
		}

		ph := h.toPrettyHour(db)
		hours = append(hours, ph)
	}
	return hours, nil
}

func GetHoursByDay(db *sql.DB, dayOfWeek int) ([]Hour, error) {
	rows, err := db.Query("SELECT id, startTime, endTime, dayOfWeek FROM hours WHERE dayOfWeek = " + strconv.Itoa(dayOfWeek))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hours := []Hour{}

	for rows.Next() {
		var h Hour
		err := rows.Scan(&h.Id, &h.StartTime, &h.EndTime, &h.DayOfWeek)
		if err != nil {
			return nil, err
		}
		hours = append(hours, h)
	}
	return hours, nil
}

func (h *Hour) DeleteHour(db *sql.DB) error {
	query := "DELETE FROM `hours` WHERE `hours`.`Id`=" + strconv.Itoa(h.Id) + ""
	_, err := db.Exec(query)
	return err
}

func GetTimeCodes(db *sql.DB) ([]TimeCode, error) {
	rows, err := db.Query("SELECT DISTINCT timeCode, startTime, endTime FROM hours;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	hours := []TimeCode{}

	for rows.Next() {
		var h TimeCode
		err := rows.Scan(&h.TimeCode, &h.StartTime, &h.EndTime)
		if err != nil {
			return nil, err
		}


		hours = append(hours, h)
	}
	return hours, nil
}
