package user

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

type TutorInformation struct{
	UserId int `json:"userId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FullName string `json:"fullName"`
}

type StandardUserInformation struct{
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
}

type AdminViewUserInformation struct{
	UserId int `json:"userId"`
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FullName string `json:"fullName"`
	Email      string `json:"email"`
	Role       string `json:"role"`
}