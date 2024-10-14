package dto

type LocalUserDTO struct {
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"isAdmin"`
}
