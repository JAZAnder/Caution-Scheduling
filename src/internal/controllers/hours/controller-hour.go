package hours

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	. "github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	"github.com/gorilla/mux"
)
type HourController struct {
	DB     *sql.DB
}


func(a *HourController) GetHour(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid hour Id")
		return
	}
	h := Hour{Id: id}
	err = h.GetHour(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Timeslot not Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, h)
}

func (a *HourController) CreateHour(w http.ResponseWriter, r *http.Request){
	var c sessionCookie

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
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ "Not an Admin")
		respondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	var h Hour
	h.StartTime = r.PostFormValue("startTime")
	h.EndTime = r.PostFormValue("endTime")
	h.DayOfWeek, err = strconv.Atoi(r.PostFormValue("dayOfWeek"))

	if (err != nil || h.DayOfWeek > 6 || h.DayOfWeek < 0) {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if illegalString(h.StartTime) ||illegalString(h.EndTime){
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ "Invalid request payload")
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.CreateHour(a.DB)
	if err != nil {
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("	Time Created by " + currentUser.UserName)

	respondWithJSON(w, http.StatusCreated, h)
}

func (a *HourController) GetHours(w http.ResponseWriter, r *http.Request){
	hours, err := GetHours(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, hours)
}

func (a *HourController) GetHoursByDay(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid hour Id")
		return
	}

	hours, err := GetHoursByDay(a.DB, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, hours)
}

func (a *HourController) DeleteHour(w http.ResponseWriter, r *http.Request){
	var c sessionCookie

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
		fmt.Println("	Fail : Time Not Deleted by " + currentUser.UserName +" : "+ "Not an Admin")
		respondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid hour ID")
		return
	}
	h := Hour{Id: id}
	if err := h.FeleteHour(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})

}

func (a *HourController) GetUsersByHour(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid hour Id")
		return
	}

	users, err := GetUsersByHour(a.DB, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}