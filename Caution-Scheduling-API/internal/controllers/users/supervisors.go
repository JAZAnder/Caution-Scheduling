package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

func resetPassword(w http.ResponseWriter, r *http.Request) {
	var c user.SessionCookie
	var u user.LocalUser
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

	authorized, err := currentUser.HasSupervisorRights()

	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !authorized {
		responses.RespondWithError(w, http.StatusForbidden, "Insufficient Permissions")
		return
	}

	u.UserName = r.PostFormValue("UserName")

	u.GetUser(database)

	if u.Role > currentUser.Role {
		responses.RespondWithError(w, http.StatusForbidden, "Supervisors are Unable to reset Administrator")
		return
	}

	u.Password = r.PostFormValue("password")

	goodPassword, err := helpers.PasswordTest(u.Password)

	if !goodPassword {
		responses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.ChangePassword(database)
	responses.RespondWithJSON(w, http.StatusCreated, "Password Changed")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
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

	authorized, err := currentUser.HasSupervisorRights()

	if (!authorized){
		if err.Error() == "not valid user" {
			responses.RespondWithError(w, http.StatusUnauthorized, err.Error())
		}else{
			responses.RespondWithError(w, http.StatusForbidden, err.Error())
		}
		
	}

	Users, err := user.GetLusers(database)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, Users)
}