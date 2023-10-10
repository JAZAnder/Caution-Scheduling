package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func(a *App) getHour(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invaid hour Id")
		return
	}

	h := hour{Id: id}
	err = h.getHour(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Lab not Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, h)
}

func (a *App) createHour(w http.ResponseWriter, r *http.Request){
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

	var h hour
	h.StartTime = r.PostFormValue("startTime")
	h.EndTime = r.PostFormValue("endTime")

	if len(h.StartTime) <= 0 || len(h.EndTime) <= 0 {
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ "Invalid request payload")
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.createHour(a.DB)
	if err != nil {
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName +" : "+ err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("	Time Created by " + currentUser.UserName)

	respondWithJSON(w, http.StatusCreated, h)
}

func (a *App) getHours(w http.ResponseWriter, r *http.Request){
	hours, err := getHours(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, hours)
}

func (a *App) deleteHour(w http.ResponseWriter, r *http.Request){
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
	h := hour{Id: id}
	if err := h.deleteHour(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})

}