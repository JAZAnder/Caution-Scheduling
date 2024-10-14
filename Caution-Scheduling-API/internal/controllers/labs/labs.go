package labs

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/JAZAnder/Caution-Scheduling/internal/dto"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers"
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/lab"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

var (
	database   = db.GetDatabase()
	labService = lab.NewLabService(database)
)

func AddLabRoutes(a *mux.Router) {
	a.HandleFunc("/api/labs", getLabs).Methods("GET")
	a.HandleFunc("/api/lab", createLab).Methods("POST")
	a.HandleFunc("/api/lab/{id:[0-9]+}", getLab).Methods("GET")
	a.HandleFunc("/api/lab/{id:[0-9]+}", updateLab).Methods("PUT")
	a.HandleFunc("/api/lab/{id:[0-9]+}", deleteLab).Methods("DELETE")
	a.HandleFunc("/api/lab/timeslot/{id:[0-9]+}", openLabTimeSlot).Methods("POST")
	a.HandleFunc("/api/lab/timeslots", getAllLabHours).Methods("GET")
	a.HandleFunc("/api/lab/timeslot/{id:[0-9]+}", removeLabTimeSlot).Methods("DELETE")
}

func getLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	l, err := labService.GetLab(id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			responses.RespondWithError(w, http.StatusNotFound, "Lab not Found")
		default:
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, l)
}

func getLabs(w http.ResponseWriter, r *http.Request) {
	labs, err := labService.GetLabs()
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, labs)
}

func createLab(w http.ResponseWriter, r *http.Request) {
	var l dto.LabDTO

	l.Location = r.PostFormValue("location")
	l.Name = r.PostFormValue("name")

	if helpers.IllegalString(l.Name) || helpers.IllegalString(l.Location) {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := labService.CreateLab(&l)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusCreated, l)
}

func updateLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	var l dto.LabDTO
	l.ID = id
	l.Location = r.PostFormValue("location")
	l.Name = r.PostFormValue("name")

	if len(l.Name) <= 0 || len(l.Location) <= 0 {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := labService.UpdateLab(&l); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, l)
}

func deleteLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid lab ID")
		return
	}

	if err := labService.DeleteLab(id); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func openLabTimeSlot(w http.ResponseWriter, r *http.Request) {
	var lh lab.LabHour
	var uh user.UserHour
	var c user.SessionCookie
	var err error
	vars := mux.Vars(r)

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.CheckSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if !currentUser.IsAdmin {
		responses.RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	lh.LabId, err = strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	lh.HourId, err = strconv.Atoi(r.PostFormValue("hourId"))
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Timeslot")
		return
	}

	uh.HourId = lh.HourId
	uh.Tutor = r.PostFormValue("TutorName")

	err = uh.GetUserHourId(database)

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "This User cannot be scheduled at this Time")
		return
	}
	lh.UserHourId = uh.Id
	lh.CreateLabTimeSlot(database)
	responses.RespondWithJSON(w, http.StatusCreated, lh)
	return

}
func getAllLabHours(w http.ResponseWriter, r *http.Request) {
	labHours, err := lab.GetLabHours(database)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, labHours)
}
func removeLabTimeSlot(w http.ResponseWriter, r *http.Request) {
	var lh lab.LabHour
	var c user.SessionCookie
	var err error
	vars := mux.Vars(r)

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.CheckSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if !currentUser.IsAdmin {
		responses.RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	lh.LabId, err = strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	lh.DeleteLabTimeSlot(database)

	responses.RespondWithJSON(w, http.StatusOK, lh)
	return

}
