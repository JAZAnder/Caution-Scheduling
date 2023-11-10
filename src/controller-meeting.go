package main

import(
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (a *App) getMeeting(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "Invalid meeting Id")
		return
	}

	m := meeting{Id: id}
	err = m.getMeeting(a.DB)
	if err != nil{
		switch err{
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Meeting not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, m)
}

// func (a *App) getMyMeetings(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil{
// 		respondWithError(w, http.StatusBadRequest, "Invalid meeting Id")
// 		return
// 	}

// 	meetings, err := getMyMeetings(a.DB)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, meetings)
// }

func (a *App) createMeeting(w http.ResponseWriter, r *http.Request){
	var m meeting
	var err error

	m.UserHourId, err = strconv.Atoi(r.PostFormValue("userHourId"))
	m.LabId, err = strconv.Atoi(r.PostFormValue("labId")) 
	m.StudentName = r.PostFormValue("studentName")
	m.StudentEmail = r.PostFormValue("studentEmail")
	if err != nil || len(m.StudentName) <= 0 || len(m.StudentEmail) <= 0{
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = m.createMeeting(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}


	respondWithJSON(w, http.StatusCreated, m)
}

func (a *App) getMeetings(w http.ResponseWriter, r *http.Request){
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
	
	meetings, err := getMeetings(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, meetings)
}

func (a *App) deleteMeeting(w http.ResponseWriter, r *http.Request){
	var c sessionCookie

	cookie, err := r.Cookie("key")
	if err != nil{
		if errors.Is(err, http.ErrNoCookie){
			respondWithError(w, http.StatusUnauthorized, "Cookie not found")
			return
		}else{
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.checkSession(a.DB)
	if err != nil{
		if err == sql.ErrNoRows{
			respondWithError(w, http.StatusUnauthorized, "Session expired")
			return
		}else{
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}

	if !currentUser.IsAdmin{
		fmt.Println("	Fail : Meeting not deleted by " + currentUser.UserName + " : " + "Not an Admin")
		respondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "Invalid meeting ID")
		return
	}
	m := meeting{Id: id}
	if err := m.deleteMeeting(a.DB); err != nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}