package userHour

type UserHour struct {
	Id        int    `json:"id"`
	HourId    int    `json:"hourId"`
	TutorId     int `json:"tutor"`
	Available bool   `json:"available"`
}

type tutorHour struct {
	Id        int    `json:"id"`
	HourId    int    `json:"hourId"`
	TutorId   int `json:"tutor"`
	Available bool   `json:"available"`
}