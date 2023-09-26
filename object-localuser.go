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

func (u *localUser) login(db *sql.DB){

}

func (u *localUser) signUp(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err; 
	}
	query := ""+string(hashedPassword)
	_, err = db.Query(query) 

	return err
}