package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
)

type sessionCookie struct {
	Id      int    `json:"id"`
	UserName string    `json:"userName"`
	Cookie  string `json:"cookie"`
}

func (c *sessionCookie) checkSession(db *sql.DB) (bool, localUser) {
	var user localUser
	
	if len(c.Cookie) < 0 {
		return false, user
	}
	
	user.UserName = "defult"
	user.IsAdmin = true
	return true, user
}

func (c *sessionCookie) createSession(db *sql.DB) (error) {
	c.Cookie = generateRandomString(10)
	query := "INSERT INTO `sessionCookie` (`cookie`, `username`) VALUES ('"+c.Cookie+"', '"+c.UserName+"');"
	err := db.QueryRow(query)

	if err != nil {
		return err.Err()
	}

	return nil 
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
	   panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
 }