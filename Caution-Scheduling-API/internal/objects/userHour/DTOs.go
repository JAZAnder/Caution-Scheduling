package userHour

type UserHour struct {
	Id        int    `json:"id"`
	HourId    int    `json:"hourId"`
	Tutor     string `json:"tutor"`
	Available bool   `json:"available"`
}