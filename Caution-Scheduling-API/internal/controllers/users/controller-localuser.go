package users

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"
)





func IsLoggedIn(w http.ResponseWriter, r *http.Request) {
	if true {
		responses.RespondWithJSON(w, http.StatusOK, map[string]string{"loggedin": "true"})
	}

	responses.RespondWithError(w, http.StatusUnauthorized, "Not Logged In")
}

func whoami(w http.ResponseWriter, r *http.Request) {
	var c user.SessionCookie
	cookie, err := r.Cookie("key")

	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	user, err := c.CheckSession(database)

	if err != nil {
		if err == sql.ErrNoRows {
			responses.RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	dto, _ := user.ToSelfViewInformation()

	responses.RespondWithJSON(w, http.StatusOK, dto)
}




func logoutLocalUser(w http.ResponseWriter, r *http.Request) {
	var c user.SessionCookie
	cookie, err := r.Cookie("key")

	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	err = c.DeleteSession(database)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
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

	responses.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}



func isAdmin(r *http.Request) (bool, error) {
	var c user.SessionCookie

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return false, errors.New("cookie not Found")
		} else {
			return false, errors.New(err.Error())
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.CheckSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("session Expired")
		} else {
			return false, errors.New(err.Error())
		}
	}

	return currentUser.IsAdmin, nil

}



// func addTime(w http.ResponseWriter, r *http.Request) {
// 	var c user.SessionCookie
// 	var uh userHour.UserHour
// 	cookie, err := r.Cookie("key")
// 	if err != nil {
// 		if errors.Is(err, http.ErrNoCookie) {
// 			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
// 			return
// 		} else {
// 			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 			return
// 		}
// 	}

// 	c.Cookie = cookie.Value

// 	currentUser, err := c.CheckSession(database)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			responses.RespondWithError(w, http.StatusUnauthorized, "Session Expired")
// 			return
// 		} else {
// 			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 			return
// 		}
// 	}

// 	uh.Tutor = currentUser.UserName
// 	uh.HourId, err = strconv.Atoi(r.PostFormValue("hourId"))
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
// 		return
// 	}

// 	err = uh.CreateUserHour(database)
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	uh.Available = true

// 	responses.RespondWithJSON(w, http.StatusCreated, uh)
// }


func removeTime() {
	//TODORemove-Time-Userhour

}
func removeTimeAdmin(w http.ResponseWriter, r *http.Request) {
	var c user.SessionCookie
	var uh userHour.UserHour
	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not Found")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Cookie = cookie.Value

	currentUser, err := c.CheckSession(database)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.RespondWithError(w, http.StatusUnauthorized, "Session Expired")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	authorization, _ := currentUser.HasSupervisorRights()

	if !authorization {
		responses.RespondWithError(w, http.StatusForbidden, "Insufficient Permissions")
		return
	}

	vars := mux.Vars(r)
	uh.Id, err = strconv.Atoi(vars["id"])

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}

	err = uh.DeleteUserHourById(database)

	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	uh.Available = true
	responses.RespondWithJSON(w, http.StatusOK, uh)

}
// func getluserTime(w http.ResponseWriter, r *http.Request) {
// 	var uh userHour.UserHour
// 	vars := mux.Vars(r)
// 	uh.Tutor = vars["username"]

// 	userHours, err := uh.GetHours(database)
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	responses.RespondWithJSON(w, http.StatusOK, userHours)
// }
// func getluserAvalibleTime(w http.ResponseWriter, r *http.Request) {
// 	var uh userHour.UserHour
// 	vars := mux.Vars(r)
// 	uh.Tutor = vars["username"]

// 	userHours, err := uh.GetAvailableHours(database)
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	responses.RespondWithJSON(w, http.StatusOK, userHours)
// }

func getUserHourById(w http.ResponseWriter, r *http.Request) {
	var uh userHour.UserHour
	var err error
	vars := mux.Vars(r)
	uh.Id, err = strconv.Atoi(vars["id"])

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Id")
		return
	}

	err = uh.GetUserHour(database)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusCreated, uh)
}

// func getAllUserHours(w http.ResponseWriter, r *http.Request) {
// 	userHours, err := userHour.GetUserHours(database)
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	responses.RespondWithJSON(w, http.StatusOK, userHours)
// }

func getUserInfo(w http.ResponseWriter, r *http.Request) {
	var u user.LocalUser
	vars := mux.Vars(r)
	u.UserName = vars["id"]

	if helpers.IllegalString(u.UserName) {
		responses.RespondWithError(w, http.StatusBadRequest, "Invaild UserId")
		return
	}

	err := u.GetUser(database)

	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, u)

}
