package userHour


type UserHour struct {
	Id        int  `json:"id"`
	HourId    int  `json:"hourId"`
	TutorId   int  `json:"tutor"`
	Available bool `json:"available"`
}

type TutorHour struct {
	Id        int  `json:"id"`
	HourId    int  `json:"hourId"`
	TutorId   int  `json:"tutor"`
	Available bool `json:"available"`
}

type TutorsAndHours struct {
	Id        string  `json:"id"`
	HourId    string  `json:"hourId"`
	TutorId   string  `json:"tutor"`
	FirstName string `json:"firstName"` 
	LastName string `json:"lastName"` 
	StartTime string `json:"startTime"` 
	EndTime string `json:"endTime"` 
	DayOfWeek string `json:"dayOfWeek"`
}