package app

import (
	"net/http"
	. "github.com/JAZAnder/Caution-Scheduling/internal/controllers/hours"
	. "github.com/JAZAnder/Caution-Scheduling/internal/controllers/labs"
	. "github.com/JAZAnder/Caution-Scheduling/internal/controllers/users"
	. "github.com/JAZAnder/Caution-Scheduling/internal/controllers/meetings"

)

//Routes

func (a *App) initializeRoutes() {
	AddLabRoutes(a.Router)
	AddUserRoutes(a.Router)
	AddHourRoutes(a.Router)
	AddMeetingRoutes(a.Router)
	a.staticRoutes()
}





func (a *App) staticRoutes(){
	//Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	a.Router.PathPrefix("/assets/").Handler(fs)


    // Serve the homepage when the root URL ("/") is accessed
	rf := http.StripPrefix("/", http.FileServer(http.Dir("./pages/")))
	a.Router.PathPrefix("/").Handler(rf)

}
