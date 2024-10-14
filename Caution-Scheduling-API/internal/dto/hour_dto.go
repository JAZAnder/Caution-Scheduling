package dto

type HourDTO struct {
	Id        int    `json:"id"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	DayOfWeek int    `json:"dayOfWeek"`
}
