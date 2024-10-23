package app

import (
	"net/http"

	hourController "github.com/JAZAnder/Caution-Scheduling/internal/controllers/hours"
	labController "github.com/JAZAnder/Caution-Scheduling/internal/controllers/labs"
	meetingController "github.com/JAZAnder/Caution-Scheduling/internal/controllers/meetings"
	userTimeslots "github.com/JAZAnder/Caution-Scheduling/internal/controllers/user-timeslots"
	userController "github.com/JAZAnder/Caution-Scheduling/internal/controllers/users"
	"github.com/gorilla/mux"
)

//Routes

func (a *App) initializeRoutes() {
	labController.AddLabRoutes(a.Router)
	userController.AddUserRoutes(a.Router)
	hourController.AddHourRoutes(a.Router)
	meetingController.AddMeetingRoutes(a.Router)
	userTimeslots.AddUseTimeslotRoutes(a.Router)
	AddStaticRoutes(a.Router)
}

func AddStaticRoutes(a *mux.Router) {
	//Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	a.PathPrefix("/assets/").Handler(fs)

	// Serve the homepage when the root URL ("/") is accessed
	//rf := http.StripPrefix("/", http.FileServer(http.Dir("./pages/")))

	a.PathPrefix("/").HandlerFunc(serveIndex)

}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
