package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type localUser struct{
	UserName string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsAdmin bool `json:"isAdmin"`
}

func (u *localUser) login(db *sql.DB) (error) {
	query := "SELECT `password` FROM `localusers` WHERE `userName` = '"+u.UserName+"';"
	result := db.QueryRow(query)
	var storedCreds localUser

	err := result.Scan(&storedCreds.Password)
	if err == sql.ErrNoRows {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(u.Password)); 
	err != nil {
		return err
	}

	//Get rest of information

	return nil
}

func (u *localUser) signUp(db *sql.DB) error {
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