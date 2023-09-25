package main
import()

type sessionCookie struct{
	Id int `json:"id"`
	WNumber int `json:"wnumber"`
	Cookie string `json:"cookie"`
	user localUser
}
