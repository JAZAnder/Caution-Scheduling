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
	a.Router.HandleFunc("/api/luser/delete/{userName}", a.deleteLocalUser).Methods("DELETE")
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