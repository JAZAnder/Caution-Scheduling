package lab

import (
	"database/sql"
	"fmt"
	"strconv"
	//"errors"
)

type Lab struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (l *Lab) GetLab(db *sql.DB) error {
	query := "SELECT name, location FROM labs WHERE id=" + strconv.Itoa(l.Id)
	fmt.Println(query)
	return db.QueryRow(query).Scan(&l.Name, &l.Location)
}

func (l *Lab) UpdateLab(db *sql.DB) error {
	query := "UPDATE `labs` SET `name` = '"+l.Name+"', `location`='"+l.Location+"' WHERE `labs`.`id`="+strconv.Itoa(l.Id)+""
	fmt.Println(query)
	_, err := db.Exec(query)
	return err
}

func (l *Lab) DeleteLab(db *sql.DB) error {
	query := "DELETE FROM `labs` WHERE `labs`.`Id`="+strconv.Itoa(l.Id)+""
	fmt.Println(query)
	_, err := db.Exec(query)
	return err
}

func (l *Lab) CreateLab(db *sql.DB) error {
	query := "INSERT INTO `labs` (`name`, `location`) VALUES ('" + l.Name + "','" + l.Location + "')"
	fmt.Println(query)
	err := db.QueryRow(query)

	if err != nil {
		return err.Err()
	}
	return nil
}

func GetLabs(db *sql.DB) ([]Lab, error) {
	rows, err := db.Query("Select id, name, location FROM labs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	labs := []Lab{}

	for rows.Next() {
		var l Lab
		if err := rows.Scan(&l.Id, &l.Name, &l.Location); err != nil {
			return nil, err
		}
		labs = append(labs, l)
	}
	return labs, nil
}