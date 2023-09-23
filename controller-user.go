package main

import (
	"net/http"
)

func (a *App) isLoggedIn(w http.ResponseWriter, r *http.Request){
	if true {
		respondWithJSON(w, http.StatusOK, map[string]string{"loggedin": "true"})
	}

	respondWithError(w, http.StatusUnauthorized, "Not Logged In")
	
}


func (a *App) isAdmin(w http.ResponseWriter, r *http.Request){
	if true {
		respondWithJSON(w, http.StatusOK, map[string]string{"isadmin": "true"})
	}
	respondWithError(w, http.StatusForbidden, "Not Admin")
	
}