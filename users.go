package main

import()

type user struct{
	FistName string `json:"firstname"`
	LastName string `json:"lastname"`
	UserName string `json:"usernaem"`
	Password string `json:"password"`
	Salt string `json:"salt"`
	IsAdmin bool `json:"isadmin"`
}
