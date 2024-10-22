package user

// Role: 0 - Deactivated
// Role: 1 - Student
// Role: 2 - Tutor
// Role: 3 - Supervisor
// Role: 4 - Administrator
type LocalUser struct {
	UserId int `json:"userId"`
	GoogleId   string `json:"googleId"`
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FullName string `json:"fullName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       int `json:"role"`
	Attempts       int `json:"attempts"`
	Disabled    bool   `json:"disabled"`
	IsAdmin    bool   `json:"isAdmin"`
	Settings   userSettings
}

type SQLLocalUser struct {
	UserId string `json:"userId"`
	GoogleId   string `json:"googleId"`
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FullName string `json:"fullName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	//Attempts       string `json:"attempts"`
	//Disabled    string   `json:"disabled"`
	IsAdmin    string   `json:"isAdmin"`
}


type SelfViewInformation struct{
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FullName string `json:"fullName"`
	Email      string `json:"email"`
	Role       string `json:"role"`
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