package helpers

import "strings"

func illegalString(test string) bool {
	if len(test) <= 0 {
		return true
	}
	if strings.ContainsAny(test, "/\\;<>'\"\b\n\r\t%") {
		return true
	}
	return false
}