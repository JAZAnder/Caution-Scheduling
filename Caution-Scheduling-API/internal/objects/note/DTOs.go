package note

type Note struct{
	Id int `json:"id"`
	UserId int `json:"userId"`
	MeetingId int `json:"meetingId"`
	Note string `json:"note"`
}