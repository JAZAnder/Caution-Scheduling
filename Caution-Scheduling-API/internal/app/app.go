package app

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	database "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {

	a.DB = database.GetDatabase()

	//Creates Router
	a.Router = mux.NewRouter()

	a.initializeRoutes()

}

func (a *App) Run(addr string) {

	http.ListenAndServe(":"+addr, a.Router)
}
