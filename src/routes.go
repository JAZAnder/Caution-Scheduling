package main

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
}

func (a *App) userRoutes(){
	a.Router.HandleFunc("/api/luser", a.createLocalUser).Methods("POST")
	a.Router.HandleFunc("/api/luser/login", a.loginLocalUser).Methods("POST")
	a.Router.HandleFunc("/api/luser/whoami", a.whoami).Methods("GET")
	a.Router.HandleFunc("/api/luser/logout", a.logoutLocalUser).Methods("DELETE")
	
}

func (a *App) hourRoutes(){
	a.Router.HandleFunc("/api/hour", a.createHour).Methods("POST")
	a.Router.HandleFunc("/api/hour/{id:[0-9]+}", a.getHour).Methods("GET")
	a.Router.HandleFunc("/api/hours", a.getHours).Methods("GET")
	a.Router.HandleFunc("/api/hour/{id:[0-9]+}", a.deleteHour).Methods("DELETE")
}

func (a *App) meetingRoutes(){
	a.Router.HandleFunc("/api/meeting", a.createMeeting).Methods("POST")
	a.Router.HandleFunc("/api/meeting/{id:[0-9]+}",a.getMeeting).Methods("GET")
	a.Router.HandleFunc("/api/meetings", a.getMeetings).Methods("GET")
	a.Router.HandleFunc("/api/meeting/{id:[0-9]+}",a.deleteMeeting).Methods("DELETE")
}

func (a *App) staticRoutes(){
	//Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	a.Router.PathPrefix("/assets/").Handler(fs)

	//Pages
	wp := http.StripPrefix("/", http.FileServer(http.Dir("./pages/")))
	a.Router.PathPrefix("/").Handler(wp)

	a.Router.HandleFunc("/hello", helloHandler).Methods("GET")
	//a.Router.HandleFunc("/", indexHandler)
}