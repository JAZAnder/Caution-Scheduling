package seeding

import (
	"database/sql"
	"fmt"

)

var database *sql.DB

func CreateTables(db *sql.DB) {
	database = db

	createLocalUserTables()
	createHoursTable()
	createLabsTable()
}

func createLocalUserTables() {
	query := "CREATE TABLE IF NOT EXISTS localusers(" +
		"`userName` varchar(255) PRIMARY KEY," +
		"`firstName` varchar(255) NOT NULL," +
		"`lastName` varchar(225) NOT NULL," +
		"`email` varchar(225) NOT NULL," +
		"`password` varchar(225) NOT NULL," +
		"`isAdmin` boolean NOT NULL);"
	fmt.Println("\nCreating localUsers table :")
	fmt.Println(query)
	_, err := database.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("localUsers table either created or already existed")
	}
}

func createHoursTable() {
	query := "CREATE TABLE IF NOT EXISTS hours(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`startTime` varchar(255) NOT NULL," +
		"`endTime` varchar(225) NOT NULL," +
		"`dayOfWeek` int DEFAULT NULL);"

	fmt.Println("\nCreating hours table :")
	fmt.Println(query)

	_, err := database.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("hours table either created or already existed")
	}
}

func createLabsTable() {
	query := "CREATE TABLE IF NOT EXISTS labs(" +
		"`Id` int AUTO_INCREMENT PRIMARY KEY," +
		"`name` varchar(255) NOT NULL," +
		"`location` varchar(225) DEFAULT NULL);"

	fmt.Println("\nCreating labs table :")
	fmt.Println(query)

	_, err := database.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("labs table either created or already existed")
	}
}
