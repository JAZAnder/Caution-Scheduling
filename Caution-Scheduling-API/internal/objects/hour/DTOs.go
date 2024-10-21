package hour

type Hour struct{
	Id int `json:"id"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	DayOfWeek int `json:"dayOfWeek"`
}

type TimeslotsMultiDay struct{
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	
	Monday bool 
	Tuesday bool
	Wednesday bool
	Thursday bool
	Friday bool
	Saturday bool
	Sunday bool
}