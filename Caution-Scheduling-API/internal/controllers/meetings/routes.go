package meetings

import "github.com/gorilla/mux"

func AddMeetingRoutes(a *mux.Router) {
	addEveryoneRoutes(a)
	addStudentRoutes(a)
	addTutorRoutes(a)
	addSupervisorRoutes(a)
	addAdminRoutes(a)
	// a.HandleFunc("/api/meeting/{id:[0-9]+}", getMeeting).Methods("GET")
	// a.HandleFunc("/api/meetings", getMeetings).Methods("GET")
	// a.HandleFunc("/api/meeting/{id:[0-9]+}", deleteMeeting).Methods("DELETE")
	// a.HandleFunc("/api/meetings/mine", getMyMeetings).Methods("GET")
}

func addEveryoneRoutes(a *mux.Router){

}

func addStudentRoutes(a *mux.Router){
	a.HandleFunc("/api/meeting", createMeeting).Methods("POST")
	a.HandleFunc("/api/meetings", getMyMeetings).Methods("GET")
}

func addTutorRoutes(a *mux.Router){
	
}

func addSupervisorRoutes(a *mux.Router){
	a.HandleFunc("/api/meetings/filter", getAllMeetingsByFilter).Methods("GET")

}

func addAdminRoutes(a *mux.Router){

}