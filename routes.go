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

func (a *App) staticRoutes(){
	//Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	a.Router.PathPrefix("/assets/").Handler(fs)

	//Pages
	wp := http.StripPrefix("/pages/", http.FileServer(http.Dir("./pages/")))
	a.Router.PathPrefix("/pages/").Handler(wp)

	a.Router.HandleFunc("/hello", helloHandler).Methods("GET")
	a.Router.HandleFunc("/", indexHandler)
}