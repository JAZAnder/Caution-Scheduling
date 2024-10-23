package users

import (
	"github.com/gorilla/mux"
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
)

func AddUserRoutes(a *mux.Router) {

	//Everyone Routes
	a.HandleFunc("/api/luser/login", loginLocalUser).Methods("POST")

	//Student Routes
	a.HandleFunc("/api/luser/whoami", whoami).Methods("GET")
	a.HandleFunc("/api/luser/logout", logoutLocalUser).Methods("DELETE")
	a.HandleFunc("/api/luser/resetmypasswd", changePassword).Methods("PUT")

	//Tutor Routes
	a.HandleFunc("/api/luser/timeslot", addTime).Methods("POST")

	//Supervisor Routes
	a.HandleFunc("/api/luser/resetpasswd", resetPassword).Methods("PUT")//Only Admins can reset another Admins password
	a.HandleFunc("/api/luser/admin/timeslot", addTimeAdmin).Methods("POST")

	//Administrator Routes
	a.HandleFunc("/api/luser/admin/timeslot/{id:[0-9]+}", removeTimeAdmin).Methods("DELETE")
	a.HandleFunc("/api/lusers", getAllUsers).Methods("GET")


	a.HandleFunc("/api/luser", createLocalUser).Methods("POST")
	a.HandleFunc("/api/tutor/availability/{username}", getluserAvalibleTime).Methods("GET")
	a.HandleFunc("/api/tutor/hours/{username}", getluserTime).Methods("GET")
	a.HandleFunc("/api/tutor/timeslot/whois/{id:[0-9]+}", getUserHourById).Methods("GET")
	a.HandleFunc("/api/tutor/timeslots", getAllUserHours).Methods("GET")
	a.HandleFunc("/api/tutor/whois/{id}", getUserInfo).Methods("GET")

}

var (
	database = db.GetDatabase()
)