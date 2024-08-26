package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	. "github.com/JAZAnder/Caution-Scheduling/internal/helpers"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {

	a.DB = GetDatabase()


	checkDatabase(a.DB)

	//Creates Router
	a.Router = mux.NewRouter()

	a.initializeRoutes()

 }

func (a *App) Run(addr string) { 

	http.ListenAndServe(":" + addr, a.Router)
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