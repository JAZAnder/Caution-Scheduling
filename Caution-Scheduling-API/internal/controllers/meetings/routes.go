package meetings

import "github.com/gorilla/mux"

func AddMeetingRoutes(a *mux.Router) {
	a.HandleFunc("/api/meeting", createMeeting).Methods("POST")
	// a.HandleFunc("/api/meeting/{id:[0-9]+}", getMeeting).Methods("GET")
	// a.HandleFunc("/api/meetings", getMeetings).Methods("GET")
	// a.HandleFunc("/api/meeting/{id:[0-9]+}", deleteMeeting).Methods("DELETE")
	// a.HandleFunc("/api/meetings/mine", getMyMeetings).Methods("GET")
}