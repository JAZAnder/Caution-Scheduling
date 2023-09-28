package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"strconv"
)

type sessionCookie struct {
	Id      int    `json:"id"`
	UserName string    `json:"userName"`
	Cookie  string `json:"cookie"`
}

func (c *sessionCookie) checkSession(db *sql.DB) (localUser, error) {
	var user localUser
	query :="SELECT `localusers`.`userName`, `localusers`.`firstName`, `localusers`.`lastName`, `localusers`.`email`, `localusers`.`isAdmin`"+
			"FROM `localusers`"+ 
				"LEFT JOIN `sessionCookie` ON `sessionCookie`.`username` = `localusers`.`userName`" +
			"WHERE `sessionCookie`.`cookie` = '"+c.Cookie+"';"
	result := db.QueryRow(query)

	var isAdmin string
	err := result.Scan(&user.UserName, &user.FirstName, &user.LastName, &user.Email, &isAdmin)

	if err != nil{
		return user, err
	}

	user.IsAdmin, _ = strconv.ParseBool(isAdmin) 
	user.Password="REDACTED"

	return user, nil
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