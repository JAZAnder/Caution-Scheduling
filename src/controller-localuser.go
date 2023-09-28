package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (a *App) isLoggedIn(w http.ResponseWriter, r *http.Request) {
	if true {
		respondWithJSON(w, http.StatusOK, map[string]string{"loggedin": "true"})
	}

	respondWithError(w, http.StatusUnauthorized, "Not Logged In")

}

func (a *App) whoami(w http.ResponseWriter, r *http.Request) {
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

	user, err := c.checkSession(a.DB)

	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (a *App) createLocalUser(w http.ResponseWriter, r *http.Request) {
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

	err = u.signUp(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("	User Created by " + currentUser.UserName)
	respondWithJSON(w, http.StatusCreated, u)
	return
}

func (a *App) loginLocalUser(w http.ResponseWriter, r *http.Request) {
	var u localUser
	var c sessionCookie

	u.UserName = r.PostFormValue("userName")
	u.Password = r.PostFormValue("password")
	err := u.login(a.DB)
	u.Password = "REDACTED"

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Username or Password Incorrect")
		return
	}

	c.UserName = u.UserName
	c.createSession(a.DB)

	cookie := http.Cookie{
		Name:     "key",
		Value:    c.Cookie,
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	respondWithJSON(w, http.StatusOK, u)
	return

}

func (a *App) isAdmin(w http.ResponseWriter, r *http.Request, name string) bool {
	var c sessionCookie

	cookie, err := r.Cookie(name)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			respondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return false
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return false
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.checkSession(a.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusUnauthorized, "Session Expired")
			return false
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return false
		}
	}

	return currentUser.IsAdmin

}
