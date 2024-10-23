package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

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

	authorized, err := currentUser.HasAdministratorRights()

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

func getUsersByFilter(w http.ResponseWriter, r *http.Request) {
	var c user.SessionCookie
	filter := user.AdminViewUserInformation{
		UserName: r.URL.Query().Get("userName"),
		FirstName: r.URL.Query().Get("firstName"),
		LastName: r.URL.Query().Get("lastName"),
		Email: r.URL.Query().Get("email"),
		Role: r.URL.Query().Get("role"),
	}
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

	authorized, err := currentUser.HasAdministratorRights()

	if (!authorized){
		if err.Error() == "not valid user" {
			responses.RespondWithError(w, http.StatusUnauthorized, err.Error())
		}else{
			responses.RespondWithError(w, http.StatusForbidden, err.Error())
		}
		
	}

	Users, err := user.GetUsersByFilter(database, filter)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, Users)
}