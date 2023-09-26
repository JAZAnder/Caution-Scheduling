package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

func (a *App) createLocalUser(w http.ResponseWriter, r *http.Request){
	var u localUser
	var c sessionCookie
	fmt.Println("CreateLocalUser - POST")
	u.UserName = r.PostFormValue("userName")
	u.FirstName = r.PostFormValue("firstName")
	u.LastName = r.PostFormValue("lastName")
	u.Email = r.PostFormValue("email")
	u.Password = r.PostFormValue("password")
	u.IsAdmin, _ = strconv.ParseBool(r.PostFormValue("isAdmin"))
	
	cookie, err := r.Cookie("key")
	if err != nil {
		if(errors.Is(err, http.ErrNoCookie)){
			respondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		}else{
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}


	c.Cookie = cookie.Value

	authtenticated, currentUser := c.checkSession(a.DB)

	if !authtenticated {
		respondWithError(w, http.StatusUnauthorized, "Session Expired")
		return
	}
	if !currentUser.IsAdmin {
		respondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	err = u.signUp(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("	User Created by "+ currentUser.UserName)
	respondWithJSON(w, http.StatusCreated, u)
	return
}