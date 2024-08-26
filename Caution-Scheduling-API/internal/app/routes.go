package app

import (
	"net/http"
	//"github.com/gorilla/mux"
	//"github.com/gin-gonic/gin"
	//"example.com/controllers"
)

//Routes

func (a *App) initializeRoutes() {
    a.labRoutes()
	a.userRoutes()
	a.hourRoutes()
	a.meetingRoutes()
	a.staticRoutes()
}

func (a *App) labRoutes(){
	a.Router.HandleFunc("/api/labs", a.getLabs).Methods("GET")
    a.Router.HandleFunc("/api/lab", a.createLab).Methods("POST")
    a.Router.HandleFunc("/api/lab/{id:[0-9]+}", a.getLab).Methods("GET")
    a.Router.HandleFunc("/api/lab/{id:[0-9]+}", a.updateLab).Methods("PUT")
    a.Router.HandleFunc("/api/lab/{id:[0-9]+}", a.deleteLab).Methods("DELETE")
	a.Router.HandleFunc("/api/lab/timeslot/{id:[0-9]+}", a.openLabTimeSlot).Methods("POST")
	a.Router.HandleFunc("/api/lab/timeslots", a.getAllLabHours).Methods("GET")
	a.Router.HandleFunc("/api/lab/timeslot/{id:[0-9]+}", a.removeLabTimeSlot).Methods("DELETE")
}

func (a *App) userRoutes(){
	a.Router.HandleFunc("/api/luser", a.createLocalUser).Methods("POST")
	a.Router.HandleFunc("/api/luser/login", a.loginLocalUser).Methods("POST")
	a.Router.HandleFunc("/api/luser/whoami", a.whoami).Methods("GET")
	a.Router.HandleFunc("/api/luser/logout", a.logoutLocalUser).Methods("DELETE")
	a.Router.HandleFunc("/api/lusers", a.getAllUsers).Methods("GET")
	a.Router.HandleFunc("/api/luser/resetmypasswd",a.changePassword).Methods("PUT")
	a.Router.HandleFunc("/api/luser/admin/resetpasswd", a.resetPassword).Methods("PUT")
	a.Router.HandleFunc("/api/luser/timeslot", a.addTime).Methods("POST")
	a.Router.HandleFunc("/api/luser/admin/timeslot", a.addTimeAdmin).Methods("POST")
	a.Router.HandleFunc("/api/luser/admin/timeslot/{id:[0-9]+}", a.removeTimeAdmin).Methods("DELETE")
	a.Router.HandleFunc("/api/tutor/availability/{username}", a.getluserAvalibleTime).Methods("GET")
	a.Router.HandleFunc("/api/tutor/hours/{username}",a.getluserTime).Methods("GET")
	a.Router.HandleFunc("/api/tutor/timeslot/whois/{id:[0-9]+}", a.getUserHourById).Methods("GET")
	a.Router.HandleFunc("/api/tutor/timeslots",a.getAllUserHours).Methods("GET")
	a.Router.HandleFunc("/api/tutor/whois/{id}", a.getUserInfo).Methods("GET")

}

func (a *App) hourRoutes(){
	a.Router.HandleFunc("/api/hour", a.createHour).Methods("POST")
	a.Router.HandleFunc("/api/hour/{id:[0-9]+}", a.getHour).Methods("GET")
	a.Router.HandleFunc("/api/hours", a.getHours).Methods("GET")
	a.Router.HandleFunc("/api/hour/day/{id:[0-9]+}", a.getHoursByDay).Methods("GET")
	a.Router.HandleFunc("/api/hour/{id:[0-9]+}", a.deleteHour).Methods("DELETE")
	a.Router.HandleFunc("/api/hour/availability/{id:[0-9]+}", a.getUsersByHour).Methods("GET")
}

func (a *App) meetingRoutes(){
	a.Router.HandleFunc("/api/meeting", a.createMeeting).Methods("POST")
	a.Router.HandleFunc("/api/meeting/{id:[0-9]+}",a.getMeeting).Methods("GET")
	a.Router.HandleFunc("/api/meetings", a.getMeetings).Methods("GET")
	a.Router.HandleFunc("/api/meeting/{id:[0-9]+}",a.deleteMeeting).Methods("DELETE")
	a.Router.HandleFunc("/api/meetings/mine", a.getMyMeetings).Methods("GET")
}

func (a *App) staticRoutes(){
	//Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	a.Router.PathPrefix("/assets/").Handler(fs)


    // Serve the homepage when the root URL ("/") is accessed
	rf := http.StripPrefix("/", http.FileServer(http.Dir("./pages/")))
	a.Router.PathPrefix("/").Handler(rf)


    a.Router.HandleFunc("/hello", helloHandler).Methods("GET")
}
