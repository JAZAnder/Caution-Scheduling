package meeting

type Meeting struct{
	Id int `json:"id"`
	UserHourId int `json:"userHourId"`
	StudentId int `json:"studentId"`
	Date int `json:"date"`
}
