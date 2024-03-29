package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	if illegalString(u.UserName) || len(u.Password) <= 0 || illegalString(u.FirstName) || illegalString(u.LastName) || illegalString(u.Email) {
		respondWithError(w, http.StatusBadRequest, "All User Feilds Must Be Vaild")
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
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
	respondWithJSON(w, http.StatusOK, u)
	return

}

func (a *App) logoutLocalUser(w http.ResponseWriter, r *http.Request) {
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

	err = c.deleteSession(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newcookie := http.Cookie{
		Name:     "key",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	http.SetCookie(w, &newcookie)

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	return
}

func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var c sessionCookie
	isAdmin := false
	cookie, err := r.Cookie("key")
	if err == nil {
		c.Cookie = cookie.Value
		currentUser, err := c.checkSession(a.DB)
		if err == nil {
			isAdmin = currentUser.IsAdmin
		}
	}

	Users, err := getLusers(a.DB, isAdmin)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, Users)
}

func (a *App) isAdmin(r *http.Request) (bool, error) {
	var c sessionCookie

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return false, errors.New("Cookie not Found")
		} else {
			return false, errors.New(err.Error())
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.checkSession(a.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("Session Expired")
		} else {
			return false, errors.New(err.Error())
		}
	}

	return currentUser.IsAdmin, nil

}

func (a *App) changePassword(w http.ResponseWriter, r *http.Request) {
	var c sessionCookie
	var u localUser
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

	u.UserName = user.UserName
	u.Password = r.PostFormValue("oldPassword")
	err = u.login(a.DB)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Password is Incorrect")
		return
	}

	u.Password = r.PostFormValue("newPassword")
	u.changePassword(a.DB)

	respondWithJSON(w, http.StatusCreated, "Password Changed")
	return

}
func (a *App) resetPassword(w http.ResponseWriter, r *http.Request) {
	var c sessionCookie
	var u localUser
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
		respondWithError(w, http.StatusForbidden, "Insufficent Permissions")
		return
	}

	u.UserName = r.PostFormValue("UserName")
	u.Password = r.PostFormValue("password")
	u.changePassword(a.DB)
	respondWithJSON(w, http.StatusCreated, "Password Changed")
	return
}
func (a *App) addTime(w http.ResponseWriter, r *http.Request) {
	var c sessionCookie
	var uh userHour
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

	uh.Tutor = currentUser.UserName
	uh.HourId, err = strconv.Atoi(r.PostFormValue("hourId"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}

	err = uh.createUserHour(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	uh.Available = true

	respondWithJSON(w, http.StatusCreated, uh)
}

func (a *App) addTimeAdmin(w http.ResponseWriter, r *http.Request) {
	var c sessionCookie
	var uh userHour
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
		respondWithError(w, http.StatusForbidden, "Insufficent Permissions")
		return
	}

	hourId := r.PostFormValue("hourId")
	uh.HourId, err = strconv.Atoi(hourId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}
	uh.Tutor = r.PostFormValue("username")

	err = uh.createUserHour(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	uh.Available = true
	respondWithJSON(w, http.StatusCreated, uh)

}
func (a *App) removeTime() {
	//TODORemove-Time-Userhour

}
func (a *App) removeTimeAdmin(w http.ResponseWriter, r *http.Request) {
	var c sessionCookie
	var uh userHour
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
		respondWithError(w, http.StatusForbidden, "Insufficent Permissions")
		return
	}


	vars := mux.Vars(r)
	uh.Id, err = strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}

	err = uh.deleteUserHourById(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	uh.Available = true
	respondWithJSON(w, http.StatusOK, uh)


}
func (a *App) getluserTime(w http.ResponseWriter, r *http.Request) {
	var uh userHour
	vars := mux.Vars(r)
	uh.Tutor = vars["username"]

	userHours, err := uh.getHours(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, userHours)
}
func (a *App) getluserAvalibleTime(w http.ResponseWriter, r *http.Request) {
	var uh userHour
	vars := mux.Vars(r)
	uh.Tutor = vars["username"]

	userHours, err := uh.getAvailableHours(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, userHours)
}

func (a *App) getUserHourById(w http.ResponseWriter, r *http.Request){
	var uh userHour
	var err error
	vars := mux.Vars(r)
	uh.Id, err = strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Id")
		return
	}

	err = uh.getUserHour(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return 
	}

	respondWithJSON(w, http.StatusCreated, uh)
}

func (a *App) getAllUserHours(w http.ResponseWriter, r *http.Request){
	userHours, err := getUserHours(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, userHours)
}

func (a *App) getUserInfo(w http.ResponseWriter, r *http.Request){
	var u localUser
	vars := mux.Vars(r)
	u.UserName = vars["id"]

	if illegalString(u.UserName){
		respondWithError(w, http.StatusBadRequest, "Invaild UserId")
		return
	}

	err := u.getUser(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
	
}