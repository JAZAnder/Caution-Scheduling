package logger

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

)

// Level 1 Debug
// Level 2 Info
// Level 3 Warning
// Level 4 Error
// Level 5 Critical
// Level 6 NoLogging

// Type

var (
	minLogLevelConsole int = 1
	minLogLevelDb      int = 6
)

var database *sql.DB

func LogSetUpDb(miLogLevel int, db *sql.DB) error {
	if miLogLevel > 5 || miLogLevel < 1 {
		err := errors.New("level is not within range")
		return err
	}

	minLogLevelDb = miLogLevel
	database = db
	return nil
}

func LogSetUpCon(miLogLevel int) error {
	if miLogLevel > 5 || miLogLevel < 1 {
		err := errors.New("level is not within range")
		return err
	}

	minLogLevelConsole = miLogLevel
	return nil
}

func Log(level int, category string, subCategory string, user string, message string) error {
	var err error
	var levelStr string
	if level > 5 || level < 1 {
		err = errors.New("level is not within range")
		return err
	}
	if len(category) > 20{
		err = errors.New("category is too long")
		return err
	}
	if len(subCategory) > 20{
		err = errors.New("subCategory is too long")
		return err
	}

	currentTime := time.Now()
	time := currentTime.Format("2006-01-02 15:04:05")

	if level == 1 {
		levelStr = "Debug"
	} else if level == 2 {
		levelStr = "Info"
	} else if level == 3 {
		levelStr = "Warning"
	} else if level == 4 {
		levelStr = "Error"
	} else if level == 5 {
		levelStr = "Critical"
	}

	if level >= minLogLevelConsole {
		err = printMessageToConsole(time, levelStr, category, subCategory, user, message)
		if err != nil {
			return err
		}
	}

	if level >= minLogLevelDb {
		err = logMessageInDatabase(time, levelStr, category, subCategory, user, message)

		if err != nil {
			return err
		}
	}

	return nil
}

func printMessageToConsole(time string, level string, category string, subCategory string, user string, message string) error {

	fmt.Println("\n " + level + " || " + time + " || Category: " + category + " - " + subCategory + " || " + user)
	fmt.Println("\t" + message + "\n")

	return nil
}

func logMessageInDatabase(time string, level string, category string, subCategory string, user string, message string) error {

	query := "Insert Into `logs` (`level`, `category`, `subCategory`, `user`, `message`) VALUES ('"+level+"','"+category+"','"+subCategory+"','"+user+"','"+message+"');"
	_, err := database.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
