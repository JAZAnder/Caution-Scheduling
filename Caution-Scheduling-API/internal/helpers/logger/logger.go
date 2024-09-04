package logger

import (
	"errors"
	"fmt"
	"time"

)

// Level 1 Debug
// Level 2 Info
// Level 3 Warning
// Level 4 Error
// Level 5 Critical

// Type

var (
	minLogLevelConsole int = 1
	minLogLevelDb      int = 1
)

func LogSetUpDb(miLogLevel int) error {
	if miLogLevel > 5 || miLogLevel < 1 {
		err := errors.New("level is not within range")
		return err
	}

	minLogLevelDb = miLogLevel
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

func Log(level int, category string, subCategory string, user string ,message string) error {
	var err error
	var levelStr string
	if level > 5 || level < 1 {
		err = errors.New("level is not within range")
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

func printMessageToConsole(time string, level string, category string, subCategory string, user string ,  message string) error {

	fmt.Println("\n " + level + " || " + time + " || Category: " + category +" - " + subCategory + " || "+ user)
	fmt.Println("\t" + message + "\n")

	return nil
}

func logMessageInDatabase(time string, level string, category string, subCategory string, user string,  message string) error {

	return nil
}
