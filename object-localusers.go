package main
import()

type localUser struct{
	WNumber int `json:"wNumber"`
	UserName string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	EncryptedPassword string `json:"password"`
	PasswordSalt string `json:"salt"`
	IsAdmin bool `json:"isAdmin"`
	
}