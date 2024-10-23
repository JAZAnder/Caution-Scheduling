package hours

import(
		db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
		"github.com/gorilla/mux"
)

func AddHourRoutes(a *mux.Router) {
	a.HandleFunc("/api/hours", createTimeslots).Methods("POST")
	a.HandleFunc("/api/hour/{id:[0-9]+}", getHour).Methods("GET")
	a.HandleFunc("/api/hours", getHours).Methods("GET") //Old
	a.HandleFunc("/api/timeslots", getHours).Methods("GET") //New
	a.HandleFunc("/api/hour/day/{id:[0-9]+}", getHoursByDay).Methods("GET")
	a.HandleFunc("/api/hour/{id:[0-9]+}", deleteHour).Methods("DELETE")
	//a.HandleFunc("/api/hour/availability/{id:[0-9]+}", getUsersByHour).Methods("GET")
}

var (
	database = db.GetDatabase()
)