package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	. "github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	. "github.com/JAZAnder/Caution-Scheduling/internal/helpers"
	"github.com/gorilla/mux"
)

type HourController struct {
	DB     *sql.DB
}

func AddHourRoutes(a *mux.Router){
	a.HandleFunc("/api/hour", createHour).Methods("POST")
	a.HandleFunc("/api/hour/{id:[0-9]+}", getHour).Methods("GET")
	a.HandleFunc("/api/hours", getHours).Methods("GET")
	a.HandleFunc("/api/hour/day/{id:[0-9]+}", getHoursByDay).Methods("GET")
	a.HandleFunc("/api/hour/{id:[0-9]+}", deleteHour).Methods("DELETE")
	a.HandleFunc("/api/hour/availability/{id:[0-9]+}", getUsersByHour).Methods("GET")
}

var(
	database = GetDatabase()
)

func getHour(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid hour Id")
		return
	}
	h := Hour{Id: id}
	err = h.GetHour(database)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Timeslot not Found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	RespondWithJSON(w, http.StatusOK, h)
}

func createHour(w http.ResponseWriter, r *http.Request){
	var c sessionCookie

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

	currentUser, err := c.checkSession(a.DB)
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
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ "Not an Admin")
		RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	var h Hour
	h.StartTime = r.PostFormValue("startTime")
	h.EndTime = r.PostFormValue("endTime")
	h.DayOfWeek, err = strconv.Atoi(r.PostFormValue("dayOfWeek"))

	if (err != nil || h.DayOfWeek > 6 || h.DayOfWeek < 0) {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if IllegalString(h.StartTime) ||IllegalString(h.EndTime){
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ "Invalid request payload")
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.CreateHour(a.DB)
	if err != nil {
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ err.Error())
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("	Time Created by " + currentUser.UserName)

	RespondWithJSON(w, http.StatusCreated, h)
}

func getHours(w http.ResponseWriter, r *http.Request){
	hours, err := GetHours(a.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, hours)
}

func getHoursByDay(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid hour Id")
		return
	}

	hours, err := GetHoursByDay(a.DB, id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, hours)
}

func deleteHour(w http.ResponseWriter, r *http.Request){
	var c sessionCookie

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

	currentUser, err := c.checkSession(a.DB)
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
		fmt.Println("	Fail : Time Not Deleted by " + currentUser.UserName +" : "+ "Not an Admin")
		RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid hour ID")
		return
	}
	h := Hour{Id: id}
	if err := h.DeleteHour(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})

}

func getUsersByHour(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid hour Id")
		return
	}

	users, err := GetUsersByHour(a.DB, id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, users)
}