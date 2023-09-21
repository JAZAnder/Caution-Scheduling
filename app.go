package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	//"fmt"
	"github.com/gorilla/mux"
	//"log"
	"net/http"
	//"github.com/gorilla/sessions"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, db, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, db ,dbname)

	var err error

	//Creates MySQL Connection
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	//Creates Router
	a.Router = mux.NewRouter()

	a.initializeRoutes()

 }

func (a *App) Run(addr string) { 

	http.ListenAndServe(":" + addr, a.Router)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}