package main

import (
	"database/sql"
	"errors"

	//"fmt"
	"net/http"
	"strconv"

	//"github.com/bytedance/sonic/decoder"
	"github.com/gorilla/mux"
)

func (a *App) getLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid lab Id")
		return
	}

	l := lab{Id: id}
	err = l.getLab(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Lab not Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, l)

}

func (a *App) getLabs(w http.ResponseWriter, r *http.Request) {
	labs, err := getLabs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, labs)
}

func (a *App) createLab(w http.ResponseWriter, r *http.Request) {
	var l lab

	l.Location = r.PostFormValue("location")
	l.Name = r.PostFormValue("name")

	if (illegalString(l.Name) || illegalString(l.Location)) {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := l.createLab(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, l)
}

func (a *App) updateLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid lab Id")
		return
	}

	var l lab

	l.Location = r.PostFormValue("location")
	l.Name = r.PostFormValue("name")

	if len(l.Name) <= 0 || len(l.Location) <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	l.Id = id

	if err := l.updateLab(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, l)

}

func (a *App) deleteLab(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid lab ID")
		return
	}

	l := lab{Id: id}
	if err := l.deleteLab(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) openLabTimeSlot(w http.ResponseWriter, r *http.Request){
	var lh labHour
	var uh userHour
	var c sessionCookie
	var err error
	vars := mux.Vars(r)

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			respondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.checkSession(a.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if !currentUser.IsAdmin {
		respondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}


	lh.LabId, err = strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid lab Id")
		return
	}
	

	lh.HourId, err = strconv.Atoi(r.PostFormValue("hourId"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid Timeslot")
		return
	}

	uh.HourId = lh.HourId
	uh.Tutor = r.PostFormValue("TutorName")

	err = uh.getUserHourId(a.DB)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "This User cannot be sceduled at this Time")
		return
	}
	lh.UserHourId = uh.Id
	lh.createLabTimeSlot(a.DB);
	respondWithJSON(w, http.StatusCreated, lh)
	return 

}
func (a *App) getAllLabHours(w http.ResponseWriter, r *http.Request) {
	labhours, err := getLabHous(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, labhours)
}
func removeLabTimeSlot(){
	//TODOAdds-labhour(Requires-Admin)
}