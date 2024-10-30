package hour

type Hour struct{
	Id int `json:"id"`
	TimeCode int `json:"timeCode"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	DayOfWeek int `json:"dayOfWeek"`
}

type FilterHour struct{
	TimeCode string `json:"timeCode"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	DayOfWeek string `json:"dayOfWeek"`
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