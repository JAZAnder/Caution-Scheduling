package user

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"strconv"
)

type SessionCookie struct {
	Id      int    `json:"id"`
	UserName string    `json:"userName"`
	Cookie  string `json:"cookie"`
}

func (c *SessionCookie) CheckSession(db *sql.DB) (LocalUser, error) {
	var suUser SQLLocalUser
	query :="SELECT `localusers`.`userName`, `localusers`.`firstName`, `localusers`.`lastName`, `localusers`.`email`, `localusers`.`isAdmin`, `localusers`.`role`, `localusers`.`Id`"+
			"FROM `localusers`"+ 
				"LEFT JOIN `sessionCookie` ON `sessionCookie`.`username` = `localusers`.`userName`" +
			"WHERE `sessionCookie`.`cookie` = '"+c.Cookie+"';"
	result := db.QueryRow(query)

	var isAdmin string
	err := result.Scan(&suUser.UserName, &suUser.FirstName, &suUser.LastName, &suUser.Email, &isAdmin, &suUser.Role, &suUser.UserId)

	if err != nil{
		return LocalUser{}, err
	}

	user, err := suUser.toLocalUser()
	if err != nil{
		return user, err
	}

	user.IsAdmin, _ = strconv.ParseBool(isAdmin) 
	user.Password="REDACTED"

	return user, nil
}


func (c *SessionCookie) CreateSession(db *sql.DB) (error) {
	c.Cookie = GenerateRandomString(10)
	query := "INSERT INTO `sessionCookie` (`cookie`, `username`) VALUES ('"+c.Cookie+"', '"+c.UserName+"');"
	err := db.QueryRow(query)

	if err != nil {
		return err.Err()
	}

	return nil 
}

func (c* SessionCookie) DeleteSession(db *sql.DB) error {
	query := "DELETE FROM `sessionCookie` WHERE `cookie` = '"+c.Cookie+"';"
	_, err := db.Exec(query)
	return err
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
	   panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}