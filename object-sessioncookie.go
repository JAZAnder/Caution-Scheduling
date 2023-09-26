package main

import "database/sql"

type sessionCookie struct {
	Id      int    `json:"id"`
	WNumber int    `json:"wnumber"`
	Cookie  string `json:"cookie"`
	user    localUser
}

// func (c *sessionCookie) createSession (db *sql.DB) error {
//  return err
// }

func (c *sessionCookie) checkSession(db *sql.DB, cookie string) bool {
	return true
}
