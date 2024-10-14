package dto

type LabHourDTO struct {
	Id         int `json:"id"`
	LabId      int `json:"labId"`
	HourId     int `json:"hourId"`
	UserHourId int `json:"userHourId"`
}
