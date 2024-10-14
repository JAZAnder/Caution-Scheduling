package dto

type MeetingDTO struct {
	Id           int    `json:"id"`
	UserHourId   int    `json:"userHourId"`
	LabId        int    `json:"labId"`
	StudentName  string `json:"studentName"`
	StudentEmail string `json:"studentEmail"`
	Date         int    `json:"date"`
}
