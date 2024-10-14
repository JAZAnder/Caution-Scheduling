package log

type Log struct {
	Id        int    `json:"id"`
	Time    string    `json:"time"`
	Level     int `json:"level"`
	Category string   `json:"category"`
	SubCategory string   `json:"subCategory"`
	User string   `json:"user"`
	Message string   `json:"message"`
}