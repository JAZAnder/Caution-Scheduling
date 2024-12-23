package users

import (
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
	"github.com/gorilla/mux"
)

func AddUserRoutes(a *mux.Router) {

	//Everyone Routes
	a.HandleFunc("/api/luser/login", loginLocalUser).Methods("POST")
	a.HandleFunc("/api/luser/google-login", googleLoginUser).Methods("POST")

	//Student Routes
	a.HandleFunc("/api/luser/whoami", whoami).Methods("GET")
	a.HandleFunc("/api/luser/logout", logoutLocalUser).Methods("DELETE")
	a.HandleFunc("/api/luser/resetmypasswd", changePassword).Methods("PUT")
	
	//a.HandleFunc("/api/lusers/tutors", getAllTutors).Methods("GET") //Works but IDK why it is here, endpoint returns student usernames

	//Tutor Routes
	//a.HandleFunc("/api/luser/timeslot", addTime).Methods("POST")

	//Supervisor Routes
	a.HandleFunc("/api/luser/resetpasswd", resetPassword).Methods("PUT") //Only Admins can reset another Admins password
	a.HandleFunc("/api/luser/update", updateUser).Methods("PUT")
	a.HandleFunc("/api/lusers", getAllUsers).Methods("GET")


	//Administrator Routes
	a.HandleFunc("/api/luser/admin/timeslot/{id:[0-9]+}", removeTimeAdmin).Methods("DELETE")
	a.HandleFunc("/api/luser", createLocalUser).Methods("POST")
	a.HandleFunc("/api/lusers/filter", getUsersByFilter).Methods("GET")

	//a.HandleFunc("/api/tutor/availability/{username}", getluserAvalibleTime).Methods("GET")
	//a.HandleFunc("/api/tutor/hours/{username}", getluserTime).Methods("GET")
	//a.HandleFunc("/api/tutor/timeslot/whois/{id:[0-9]+}", getUserHourById).Methods("GET")
	//a.HandleFunc("/api/tutor/timeslots", getAllUserHours).Methods("GET")
	a.HandleFunc("/api/tutor/whois/{id}", getUserInfo).Methods("GET")

}

var (
	database = db.GetDatabase()
)
