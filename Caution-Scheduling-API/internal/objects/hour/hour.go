package hour

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"

)

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

	h.DayOfWeek, err = strconv.Atoi(tempDayOfWeek)
	return err
}

func (h *Hour) UpdateHour(db *sql.DB) error {
	query := "UPDATE `hours` SET `startTime` = '" + h.StartTime + "', `endTime` = '" + h.EndTime + "' WHERE `hours`.`id` = " + strconv.Itoa(h.Id) + ""
	_, err := db.Exec(query)
	return err
}

func (h *Hour) CreateHour(db *sql.DB) error {
	query := "INSERT INTO `hours` (`timeCode`,`startTime`, `endTime`, `dayOfWeek`) VALUES ('" + strconv.Itoa(h.TimeCode) + "','" + h.StartTime + "','" + h.EndTime + "','" + strconv.Itoa(h.DayOfWeek) + "')"
	err := db.QueryRow(query)

	if err != nil {
		return err.Err()
	}
	return nil
}

func MassCreateHour(db *sql.DB, timeslots []Hour) error {

	query := "INSERT INTO `hours` (`timeCode`,`startTime`, `endTime`, `dayOfWeek`) VALUES "

	for _, timeslot := range timeslots {
		sqlString := "('" + strconv.Itoa(timeslot.TimeCode) + "','" + timeslot.StartTime + "','" + timeslot.EndTime + "','" + strconv.Itoa(timeslot.DayOfWeek) + "'),"
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

func GetHours(db *sql.DB) ([]Hour, error) {
	rows, err := db.Query("SELECT id, startTime, endTime, dayOfWeek FROM hours")

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
