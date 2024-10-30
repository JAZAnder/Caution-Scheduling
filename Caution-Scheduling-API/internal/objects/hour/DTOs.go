package hour

type Hour struct{
	Id int `json:"id"`
	TimeCode int `json:"timeCode"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	DayOfWeek int `json:"dayOfWeek"`
	Active bool `json:"active"`
}

type FilterHour struct{
	TimeCode string `json:"timeCode"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	DayOfWeek string `json:"dayOfWeek"`
}

type SQLHour struct{
	TimeCode string `json:"timeCode"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	DayOfWeek string `json:"dayOfWeek"`
	Active string `json:"active"`
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