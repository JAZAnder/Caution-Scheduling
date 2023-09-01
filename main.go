package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/joho/godotenv"
	"os"
	"html/template"
)

//Entry Point
func main()  {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error Loading .env file")
	}
	//Grabs Port from .env if Empty defults 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	
	r := newRouter()
	http.ListenAndServe(":"+port, r)
}

//Routes
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)

	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets/").Handler(fs)

	return r
}

//handlers
func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("./assets/index.html"))
	tmpl.Execute(w, nil)
}