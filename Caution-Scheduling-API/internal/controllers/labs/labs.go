package main

import (
	"database/sql"
	"errors"

	//"fmt"
	"net/http"
	"strconv"

	//"github.com/bytedance/sonic/decoder"
	"github.com/gorilla/mux"
	. "github.com/JAZAnder/Caution-Scheduling/internal/helpers"
	. "github.com/JAZAnder/Caution-Scheduling/internal/objects/lab"
)

func AddLabRoutes(a *mux.Router){
	a.HandleFunc("/api/labs", getLabs).Methods("GET")
    a.HandleFunc("/api/lab", createLab).Methods("POST")
    a.HandleFunc("/api/lab/{id:[0-9]+}", getLab).Methods("GET")
    a.HandleFunc("/api/lab/{id:[0-9]+}", updateLab).Methods("PUT")
    a.HandleFunc("/api/lab/{id:[0-9]+}", deleteLab).Methods("DELETE")
	a.HandleFunc("/api/lab/timeslot/{id:[0-9]+}", openLabTimeSlot).Methods("POST")
	a.HandleFunc("/api/lab/timeslots", getAllLabHours).Methods("GET")
	a.HandleFunc("/api/lab/timeslot/{id:[0-9]+}", removeLabTimeSlot).Methods("DELETE")
}

var(
	database = GetDatabase()
)

func getLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	l := Lab{Id: id}
	err = l.GetLab(database)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Lab not Found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondWithJSON(w, http.StatusOK, l)

}

func getLabs(w http.ResponseWriter, r *http.Request) {
	labs, err := GetLabs(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, labs)
}

func createLab(w http.ResponseWriter, r *http.Request) {
	var l Lab

	l.Location = r.PostFormValue("location")
	l.Name = r.PostFormValue("name")

	if (IllegalString(l.Name) || IllegalString(l.Location)) {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := l.CreateLab(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, l)
}

func updateLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	var l Lab

	l.Location = r.PostFormValue("location")
	l.Name = r.PostFormValue("name")

	if len(l.Name) <= 0 || len(l.Location) <= 0 {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	l.Id = id

	if err := l.UpdateLab(database); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, l)

}

func deleteLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid lab ID")
		return
	}

	l := Lab{Id: id}
	if err := l.DeleteLab(database); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func openLabTimeSlot(w http.ResponseWriter, r *http.Request){
	var lh LabHour
	var uh UserHour
	var c SessionCookie
	var err error
	vars := mux.Vars(r)

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.checkSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if !currentUser.IsAdmin {
		RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}


	lh.LabId, err = strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}
	

	lh.HourId, err = strconv.Atoi(r.PostFormValue("hourId"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Timeslot")
		return
	}

	uh.HourId = lh.HourId
	uh.Tutor = r.PostFormValue("TutorName")

	err = uh.getUserHourId(database)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "This User cannot be scheduled at this Time")
		return
	}
	lh.UserHourId = uh.Id
	lh.createLabTimeSlot(database);
	RespondWithJSON(w, http.StatusCreated, lh)
	return 

}
func getAllLabHours(w http.ResponseWriter, r *http.Request) {
	labHours, err := GetLabHours(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, labHours)
}
func removeLabTimeSlot(w http.ResponseWriter, r *http.Request){
	var lh labHour
	var c sessionCookie
	var err error
	vars := mux.Vars(r)

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.checkSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if !currentUser.IsAdmin {
		RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	lh.LabId, err = strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	lh.deleteLabTimeSlot(database);

	RespondWithJSON(w, http.StatusOK, lh)
	return 


}