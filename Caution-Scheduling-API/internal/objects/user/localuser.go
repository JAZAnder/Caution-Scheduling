package user

import (
	"database/sql"
	"strconv"

	"golang.org/x/crypto/bcrypt"

)

type LocalUser struct {
	UserId int `json:"userId"`
	GoogleId   string `json:"googleId"`
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FullName string `json:"fullName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	IsAdmin    bool   `json:"isAdmin"`
	Settings   userSettings
}

func (u *LocalUser) Login(db *sql.DB) error {
	query := "SELECT `password` FROM `localusers` WHERE `userName` = '" + u.UserName + "';"
	result := db.QueryRow(query)

	var storedCreds LocalUser

	err := result.Scan(&storedCreds.Password)
	if err == sql.ErrNoRows {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(u.Password)); err != nil {
		return err
	}

	query = "SELECT `firstName`, `lastName`, `email`, `isAdmin` FROM `localusers` WHERE `userName` = '" + u.UserName + "';"
	result = db.QueryRow(query)

	var isAdmin string
	err = result.Scan(&u.FirstName, &u.LastName, &u.Email, &isAdmin)
	if err != nil {
		return err
	}
	u.IsAdmin, _ = strconv.ParseBool(isAdmin)

	return nil
}

func (u *LocalUser) SignUp(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err
	}

	isAdmin := "0"
	if u.IsAdmin {
		isAdmin = "1"
	}

	query := "INSERT INTO `userSettings` (`userName`, `ReceiveMeetingEmails`) VALUES ('" + u.UserName + "', '" + "1" + "');"
	db.QueryRow(query)

	query = "INSERT INTO `localusers` (`userName`, `firstName`, `lastName`, `email`, `password`, `isAdmin`, `role`, `fullName`, `googleId`) VALUES ('" + u.UserName + "', '" + u.FirstName + "', '" + u.LastName + "', '" + u.Email + "', '" + string(hashedPassword) + "', '" + isAdmin + "', '" + u.Role + "', '" + u.FullName + "', '" + u.GoogleId + "');"
	sqlerr := db.QueryRow(query)

	if sqlerr != nil {
		return sqlerr.Err()
	}

	return nil
}

func GetLusers(db *sql.DB, isAdmin bool) ([]LocalUser, error) {
	rows, err := db.Query("SELECT * FROM `localusers`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lusers := []LocalUser{}

	for rows.Next() {
		var u LocalUser
		var uAdmin string
		if err := rows.Scan(&u.UserName, &u.FirstName, &u.LastName, &u.Email, &u.Password, &uAdmin); err != nil {
			return nil, err
		}
		u.Password = "REDACTED"
		u.IsAdmin, _ = strconv.ParseBool(uAdmin)
		if !isAdmin {
			u.IsAdmin = false
		}
		lusers = append(lusers, u)
	}
	return lusers, nil
}

func (u *LocalUser) ChangePassword(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err
	}
	query := "UPDATE `localusers` SET `password` = '" + string(hashedPassword) + "'WHERE `userName` = '" + u.UserName + "';"
	sqlerr := db.QueryRow(query)

	if sqlerr != nil {
		return sqlerr.Err()
	}

	return nil
}

func (u *LocalUser) GetUser(db *sql.DB) error {
	query := "SELECT `firstName`, `lastName` FROM `localusers` WHERE `userName` = '" + u.UserName + "'"
	err := db.QueryRow(query).Scan(&u.FirstName, &u.LastName)
	u.Email = "REDACTED"
	u.UserName = "REDACTED"
	u.Password = "REDACTED"
	u.IsAdmin = false
	return err
}

func (u *LocalUser) GetFullName(db *sql.DB) error {
	query := "SELECT `firstName`, `lastName` FROM `localusers` WHERE `userName` = '" + u.UserName + "'"
	err := db.QueryRow(query).Scan(&u.FirstName, &u.LastName)
	u.Email = "REDACTED"
	u.UserName = "REDACTED"
	u.Password = "REDACTED"
	u.IsAdmin = false
	return err
}