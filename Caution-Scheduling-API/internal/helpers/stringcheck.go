package helpers

import "strings"

func IllegalString(test string) bool {
	if len(test) <= 0 {
		return true
	}
	if strings.ContainsAny(test, "/\\;<>'\"\b\n\r\t%") {
		return true
	}
	return false
}