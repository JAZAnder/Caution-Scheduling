package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	checkDatabase(a.DB)

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

func checkDatabase(db *sql.DB){
	dir := "./mysql-tables"

	files, err := os.ReadDir(dir)
	if err != nil{
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir(){
			filePath := fmt.Sprintf("%s/%s", dir, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", file.Name(), err)
			} else {
				_ ,err := db.Exec(string(content))
				if err != nil {fmt.Println(err)}
			}

		}
	}
}

func illegalString(test string) bool{
	if(len(test) <= 0){return true}
	if(strings.ContainsAny(test, "/\\-;<>'\"\b\n\r\t%_")){return true}
	return false
}