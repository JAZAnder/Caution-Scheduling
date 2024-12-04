
package userTimeslots

import(
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
		"github.com/gorilla/mux"
)

func AddUseTimeslotRoutes(a *mux.Router) {
	//Admin Routes

	//Supervisor Routes
	a.HandleFunc("/api/availability", getEverythingByFilter).Methods("GET")
	a.HandleFunc("/api/availability", addTimeAdmin).Methods("POST")
	//Tutor Routes

	//Student Routes
	a.HandleFunc("/api/availability/{tutorId}/{date}", getTutorsAvailability).Methods("GET")
	a.HandleFunc("/api/availability/{date}", getTutorsAvailabilityByDateOnly).Methods("GET")
	//Global Routes

}



var (
	database = db.GetDatabase()
)