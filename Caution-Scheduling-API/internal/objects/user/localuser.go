package user

import (
	"database/sql"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type LocalUser struct{
	UserName string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsAdmin bool `json:"isAdmin"`
}

func (u *LocalUser) Login(db *sql.DB) (error) {
	query := "SELECT `password` FROM `localusers` WHERE `userName` = '"+u.UserName+"';"
	result := db.QueryRow(query)

	var storedCreds LocalUser

	err := result.Scan(&storedCreds.Password)
	if err == sql.ErrNoRows {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(u.Password)); 
	err != nil {
		return err
	}

	query = "SELECT `firstName`, `lastName`, `email`, `isAdmin` FROM `localusers` WHERE `userName` = '"+u.UserName+"';"
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
		return err; 
	}

	isAdmin := "0"
	if(u.IsAdmin){
		isAdmin = "1"
	}

	query := "INSERT INTO `localusers` (`userName`, `firstName`, `lastName`, `email`, `password`, `isAdmin`) VALUES ('"+u.UserName+"', '"+u.FirstName+"', '"+u.LastName+"', '"+u.Email+"', '"+string(hashedPassword)+"', '"+isAdmin+"');"
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
		if (!isAdmin) {
			u.IsAdmin = false
		}
		lusers = append(lusers, u)
	}
	return lusers, nil
}

func (u *LocalUser) ChangePassword(db *sql.DB)error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err; 
	}
	query := "UPDATE `localusers` SET `password` = '"+string(hashedPassword)+"'WHERE `userName` = '"+u.UserName+"';"
	sqlerr := db.QueryRow(query) 

	if sqlerr != nil {
		return sqlerr.Err()
	}

	return nil
}

func (u *LocalUser) GetUser(db *sql.DB) error{
	query := "SELECT `firstName`, `lastName` FROM `localusers` WHERE `userName` = '" + u.UserName + "'"
	err := db.QueryRow(query).Scan(&u.FirstName, &u.LastName)
	u.Email = "REDACTED"
	u.UserName = "REDACTED"
	u.Password = "REDACTED"
	u.IsAdmin = false
	return err
}