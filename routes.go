package main

import (
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/gin-gonic/gin"
	//"example.com/controllers"
)

//Routes

//func newRouter() *gin.Engine {
// r := gin.Default()

// r.GET("/ping", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// })

// r.GET("/api/labs/test", controllers)
//return r
//}

func (a *App) initializeRoutes() {
    a.labRoutes()
}

func (a *App) labRoutes(){
	a.Router.HandleFunc("/api/labs", a.getLabs).Methods("GET")
    a.Router.HandleFunc("/api/lab", a.createLab).Methods("POST")
    a.Router.HandleFunc("/api/lab/{id:[0-9]+}", a.getLab).Methods("GET")
    a.Router.HandleFunc("/api/lab/{id:[0-9]+}", a.updateLab).Methods("PUT")
    a.Router.HandleFunc("/api/lab/{id:[0-9]+}", a.deleteLab).Methods("DELETE")
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/testsql", helloHandler)

	// - - - - - - Static Reasourses - - - - - -
	//Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets/").Handler(fs)

	//Pages
	wp := http.StripPrefix("/pages/", http.FileServer(http.Dir("./pages/")))
	r.PathPrefix("/pages/").Handler(wp)

	r.HandleFunc("/", indexHandler)

	// // - - - - - - API Endpoints - - - - - -

	// //Labs
	// 	//Get List of All Labs
	// 	r.HandleFunc("/api/labs",helloHandler).Methods("GET")
	// 	//Adds a new lab based on Info in POST (Name) and Cookie (Session_Token)
	// 	r.HandleFunc("/api/labs/add", helloHandler).Methods("POST")

	// //LabHours
	// 	//Gets a labs hours based on LabId
	// 	r.HandleFunc("/api/labHours/{labId}", helloHandler).Methods("GET")
	// 	//Changes lab hours based on LabId POST(action, hours) and Cookie(Session_Token)
	// 	r.HandleFunc("/api/labHours/add/{labId}", helloHandler).Methods("POST")
	return r
}
