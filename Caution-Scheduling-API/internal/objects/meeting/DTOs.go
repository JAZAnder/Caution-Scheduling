package meeting

import (
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/note"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/topic"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"

)

type Meeting struct {
	Id         int `json:"id"`
	UserHourId int `json:"userHourId"`
	StudentId  int `json:"studentId"`
	Date       int `json:"date"`
	TopicId    int `json:"topicId"`
}

type BasicMeetingDto struct {
	Id      int `json:"id"`
	Hour    hour.TimeOnlyDto
	Tutor   user.StandardUserInformationId
	Student user.StandardUserInformationId
	Date    int `json:"date"`
	Topic   topic.Topic
	Notes   []note.Note
}
