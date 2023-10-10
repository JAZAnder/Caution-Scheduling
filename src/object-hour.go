package main
import(
	"database/sql"
	"strconv"
)

type hour struct{
	Id int `json:"id"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
}

func (h *hour) getHour(db *sql.DB) error{
	query := "SELECT startTime, endTime FROM hours WHERE id=" + strconv.Itoa(h.Id)
	return db.QueryRow(query).Scan(&h.startTime, &h.endTime)
}

func (h *hour) updateHour(db *sql.DB) error{
	query := "UPDATE `hours` SET `startTime` = '"+h.startTime+"', `endTime` = '"+endTime+"' WHERE `hours`.`id` = "+strconv.Itoa(h.Id)+""
	_, err := db.Exec(query)
	return err
}

func (h *hour) createHour(db *sql.DB) error{
	query := "INSERT INTO `hours` (`startTime`, `endTime`) VALUES ('"+h.startTime+"','"+h.endTime+"')"
	err := db.QueryRow(query)

	if err != nil{
		return err.Err()
	}
	return nil
}

func getHours(db *sql.DB) ([]hour, error){
	rows, err := db.Query("SELECT id, startTime, endTime FROM hours")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	hours := []hour{}

	for rows.Next(){
		var h hour
		if err:= rows.Scan(&h.Id, &h.startTime, &h.endTime); err!= nil{
			return nil,err
		}
		hours = append(hours, h)
	}
	return hours, nil
}