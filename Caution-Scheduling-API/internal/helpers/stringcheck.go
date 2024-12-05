package helpers

import (
	"errors"
	"strings"
)

func IllegalString(test string) bool {
	if len(test) <= 0 {
		return true
	}
	if strings.ContainsAny(test, "/\\;<>'\"\b\n\r\t%") {
		return true
	}
	return false
}

func PasswordTest(password string) (bool, error) {
	if len(password) <= 0 {
		return false, errors.New("passwords Cannot be Blank")
	}
	if !(strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")) {
		return false, errors.New("passwords Must Contain Capital Letter")
	}
	if !(strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")) {
		return false, errors.New("passwords Must Contain Lowercase Letter")
	}
	if !(strings.ContainsAny(password, "0123456789")) {
		return false, errors.New("passwords Must Contain Number")
	}
	if !(strings.ContainsAny(password, "!@#$%^&*(),.?\":{}|<>")) {
		return false, errors.New("passwords Must Contain Special Character")
	}
	return true, nil
}