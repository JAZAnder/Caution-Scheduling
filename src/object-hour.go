package main
import(
	"database/sql"
	"strconv"
)

type hour struct{
	Id int `json:"id"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	DayOfWeek int `json:"dayOfWeek"`
}

func (h *hour) getHour(db *sql.DB) error{
	var tempDayOfWeek string
	query := "SELECT startTime, endTime, dayOfWeek  FROM hours WHERE id=" + strconv.Itoa(h.Id)
	err := db.QueryRow(query).Scan(&h.StartTime, &h.EndTime, &tempDayOfWeek)

	h.DayOfWeek, err = strconv.Atoi(tempDayOfWeek)
	return err
}

func (h *hour) updateHour(db *sql.DB) error{
	query := "UPDATE `hours` SET `startTime` = '"+h.StartTime+"', `endTime` = '"+h.EndTime+"' WHERE `hours`.`id` = "+strconv.Itoa(h.Id)+""
	_, err := db.Exec(query)
	return err
}

func (h *hour) createHour(db *sql.DB) error{
	query := "INSERT INTO `hours` (`startTime`, `endTime`, `dayOfWeek`) VALUES ('"+h.StartTime+"','"+h.EndTime+"','"+strconv.Itoa(h.DayOfWeek)+"')"
	err := db.QueryRow(query)

	if err != nil{
		return err.Err()
	}
	return nil
}

func getHours(db *sql.DB) ([]hour, error){
	rows, err := db.Query("SELECT id, startTime, endTime, dayOfWeek FROM hours")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	hours := []hour{}

	for rows.Next(){
		var h hour
		err := rows.Scan(&h.Id, &h.StartTime, &h.EndTime, &h.DayOfWeek);
		if err!= nil{
			return nil,err
		}
		hours = append(hours, h)
	}
	return hours, nil
}

func getHoursByDay(db *sql.DB, dayOfWeek int) ([]hour, error){
	rows, err := db.Query("SELECT id, startTime, endTime, dayOfWeek FROM hours WHERE dayOfWeek = " + strconv.Itoa(dayOfWeek))

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	hours := []hour{}

	for rows.Next(){
		var h hour
		err := rows.Scan(&h.Id, &h.StartTime, &h.EndTime, &h.DayOfWeek);
		if err!= nil{
			return nil,err
		}
		hours = append(hours, h)
	}
	return hours, nil
}

func (h *hour) deleteHour(db *sql.DB) error{
	query := "DELETE FROM `hours` WHERE `hours`.`Id`="+strconv.Itoa(h.Id)+""
	_, err := db.Exec(query)
	return err
}