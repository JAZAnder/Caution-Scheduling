package main

import (
	"database/sql"
	"fmt"
	"strconv"
	//"errors"
)

type lab struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (l *lab) getLab(db *sql.DB) error {
	query := "SELECT name, location FROM labs WHERE id=" + strconv.Itoa(l.Id)
	fmt.Println(query)
	return db.QueryRow(query).Scan(&l.Name, &l.Location)
}

func (l *lab) updateLab(db *sql.DB) error {
	query := "UPDATE `labs` SET `name` = '"+l.Name+"', `location`='"+l.Location+"' WHERE `labs`.`id`="+strconv.Itoa(l.Id)+""
	fmt.Println(query)
	_, err := db.Exec(query)
	return err
}

func (l *lab) deleteLab(db *sql.DB) error {
	query := "DELETE FROM `labs` WHERE `labs`.`Id`="+strconv.Itoa(l.Id)+""
	fmt.Println(query)
	_, err := db.Exec(query)
	return err
}

func (l *lab) createLab(db *sql.DB) error {
	query := "INSERT INTO `labs` (`name`, `location`) VALUES ('" + l.Name + "','" + l.Location + "')"
	fmt.Println(query)
	err := db.QueryRow(query)

	if err != nil {
		return err.Err()
	}
	return nil
}

func getLabs(db *sql.DB) ([]lab, error) {
	rows, err := db.Query("Select id, name, location FROM labs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	labs := []lab{}

	for rows.Next() {
		var l lab
		if err := rows.Scan(&l.Id, &l.Name, &l.Location); err != nil {
			return nil, err
		}
		labs = append(labs, l)
	}
	return labs, nil
}