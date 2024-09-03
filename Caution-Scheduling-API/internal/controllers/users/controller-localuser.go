package users

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	. "github.com/JAZAnder/Caution-Scheduling/internal/helpers"
	. "github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

func AddUserRoutes(a *mux.Router){
	a.HandleFunc("/api/luser", createLocalUser).Methods("POST")
	a.HandleFunc("/api/luser/login", loginLocalUser).Methods("POST")
	a.HandleFunc("/api/luser/whoami", whoami).Methods("GET")
	a.HandleFunc("/api/luser/logout", logoutLocalUser).Methods("DELETE")
	a.HandleFunc("/api/lusers", getAllUsers).Methods("GET")
	a.HandleFunc("/api/luser/resetmypasswd",changePassword).Methods("PUT")
	a.HandleFunc("/api/luser/admin/resetpasswd", resetPassword).Methods("PUT")
	a.HandleFunc("/api/luser/timeslot", addTime).Methods("POST")
	a.HandleFunc("/api/luser/admin/timeslot", addTimeAdmin).Methods("POST")
	a.HandleFunc("/api/luser/admin/timeslot/{id:[0-9]+}", removeTimeAdmin).Methods("DELETE")
	a.HandleFunc("/api/tutor/availability/{username}", getluserAvalibleTime).Methods("GET")
	a.HandleFunc("/api/tutor/hours/{username}",getluserTime).Methods("GET")
	a.HandleFunc("/api/tutor/timeslot/whois/{id:[0-9]+}", getUserHourById).Methods("GET")
	a.HandleFunc("/api/tutor/timeslots",getAllUserHours).Methods("GET")
	a.HandleFunc("/api/tutor/whois/{id}", getUserInfo).Methods("GET")

}

var(
	database = GetDatabase()
)

func isLoggedIn(w http.ResponseWriter, r *http.Request) {
	if true {
		RespondWithJSON(w, http.StatusOK, map[string]string{"loggedin": "true"})
	}

	RespondWithError(w, http.StatusUnauthorized, "Not Logged In")
}

func whoami(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
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

	user, err := c.CheckSession(database)

	if err != nil {
		if err == sql.ErrNoRows {
			RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	RespondWithJSON(w, http.StatusOK, user)
}

func createLocalUser(w http.ResponseWriter, r *http.Request) {
	var u LocalUser
	var c SessionCookie
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
			RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.CheckSession(database)
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
	if IllegalString(u.UserName) || len(u.Password) <= 0 || IllegalString(u.FirstName) || IllegalString(u.LastName) || IllegalString(u.Email) {
		RespondWithError(w, http.StatusBadRequest, "All User Feilds Must Be Vaild")
		return
	}
	err = u.SignUp(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("	User Created by " + currentUser.UserName)
	RespondWithJSON(w, http.StatusCreated, u)
	return
}

func loginLocalUser(w http.ResponseWriter, r *http.Request) {
	var u LocalUser
	var c SessionCookie

	u.UserName = r.PostFormValue("userName")
	u.Password = r.PostFormValue("password")
	err := u.Login(database)
	u.Password = "REDACTED"

	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Username or Password Incorrect")
		return
	}

	c.UserName = u.UserName
	c.CreateSession(database)
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
	RespondWithJSON(w, http.StatusOK, u)
	return

}

func logoutLocalUser(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
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

	err = c.DeleteSession(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
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

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	return
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
	isAdmin := false
	cookie, err := r.Cookie("key")
	if err == nil {
		c.Cookie = cookie.Value
		currentUser, err := c.CheckSession(database)
		if err == nil {
			isAdmin = currentUser.IsAdmin
		}
	}

	Users, err := GetLusers(database, isAdmin)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, Users)
}

func isAdmin(r *http.Request) (bool, error) {
	var c SessionCookie

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return false, errors.New("Cookie not Found")
		} else {
			return false, errors.New(err.Error())
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.CheckSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("Session Expired")
		} else {
			return false, errors.New(err.Error())
		}
	}

	return currentUser.IsAdmin, nil

}

func changePassword(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
	var u LocalUser
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
	user, err := c.CheckSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	u.UserName = user.UserName
	u.Password = r.PostFormValue("oldPassword")
	err = u.Login(database)

	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Password is Incorrect")
		return
	}

	u.Password = r.PostFormValue("newPassword")
	u.ChangePassword(database)

	RespondWithJSON(w, http.StatusCreated, "Password Changed")
	return

}
func resetPassword(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
	var u LocalUser
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

	currentUser, err := c.CheckSession(database)
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
		RespondWithError(w, http.StatusForbidden, "Insufficent Permissions")
		return
	}

	u.UserName = r.PostFormValue("UserName")
	u.Password = r.PostFormValue("password")
	u.ChangePassword(database)
	RespondWithJSON(w, http.StatusCreated, "Password Changed")
	return
}
func addTime(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
	var uh UserHour
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

	currentUser, err := c.CheckSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	uh.Tutor = currentUser.UserName
	uh.HourId, err = strconv.Atoi(r.PostFormValue("hourId"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}

	err = uh.CreateUserHour(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	uh.Available = true

	RespondWithJSON(w, http.StatusCreated, uh)
}

func addTimeAdmin(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
	var uh UserHour
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

	currentUser, err := c.CheckSession(database)
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
		RespondWithError(w, http.StatusForbidden, "Insufficent Permissions")
		return
	}

	hourId := r.PostFormValue("hourId")
	uh.HourId, err = strconv.Atoi(hourId)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}
	uh.Tutor = r.PostFormValue("username")

	err = uh.CreateUserHour(database)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	uh.Available = true
	RespondWithJSON(w, http.StatusCreated, uh)

}
func removeTime() {
	//TODORemove-Time-Userhour

}
func removeTimeAdmin(w http.ResponseWriter, r *http.Request) {
	var c SessionCookie
	var uh UserHour
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

	currentUser, err := c.CheckSession(database)
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
		RespondWithError(w, http.StatusForbidden, "Insufficent Permissions")
		return
	}

	vars := mux.Vars(r)
	uh.Id, err = strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}

	err = uh.DeleteUserHourById(database)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	uh.Available = true
	RespondWithJSON(w, http.StatusOK, uh)

}
func getluserTime(w http.ResponseWriter, r *http.Request) {
	var uh UserHour
	vars := mux.Vars(r)
	uh.Tutor = vars["username"]

	userHours, err := uh.GetHours(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, userHours)
}
func getluserAvalibleTime(w http.ResponseWriter, r *http.Request) {
	var uh UserHour
	vars := mux.Vars(r)
	uh.Tutor = vars["username"]

	userHours, err := uh.GetAvailableHours(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, userHours)
}

func getUserHourById(w http.ResponseWriter, r *http.Request) {
	var uh UserHour
	var err error
	vars := mux.Vars(r)
	uh.Id, err = strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Id")
		return
	}

	err = uh.GetUserHour(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, uh)
}

func getAllUserHours(w http.ResponseWriter, r *http.Request) {
	userHours, err := GetUserHours(database)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, userHours)
}

func getUserInfo(w http.ResponseWriter, r *http.Request) {
	var u LocalUser
	vars := mux.Vars(r)
	u.UserName = vars["id"]

	if IllegalString(u.UserName) {
		RespondWithError(w, http.StatusBadRequest, "Invaild UserId")
		return
	}

	err := u.GetUser(database)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, u)

}
