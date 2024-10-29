
package userTimeslots

import(
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
		"github.com/gorilla/mux"
)

func AddUseTimeslotRoutes(a *mux.Router) {
	//Admin Routes

	//Supervisor Routes
	a.HandleFunc("/api/availability", getEverythingByFilter).Methods("GET")
	//Tutor Routes

	//Student Routes
	a.HandleFunc("/api/availability/{userId}", getTutorsAvailability).Methods("GET")
	//Global Routes

}



var (
	database = db.GetDatabase()
)