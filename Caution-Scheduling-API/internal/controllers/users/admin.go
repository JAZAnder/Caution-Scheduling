package users

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

func createLocalUser(w http.ResponseWriter, r *http.Request) {
	

	var c user.SessionCookie
	userToCreateDto := user.CreateLocalUserDto{
		UserName: r.PostFormValue("userName"),
		FirstName: r.PostFormValue("firstName"),
		LastName: r.PostFormValue("lastName"),
		Email: r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
		Role: r.PostFormValue("role"),
	}

	userToCreate, _ := userToCreateDto.ToLocalUser()


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

	authorized, _ := currentUser.HasAdministratorRights()
	if  !authorized{
		responses.RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	goodPassword, err := helpers.PasswordTest(userToCreate.Password)

	if !goodPassword {
		responses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	if helpers.IllegalString(userToCreate.UserName) || len(userToCreate.Password) <= 0 || helpers.IllegalString(userToCreate.FirstName) || helpers.IllegalString(userToCreate.LastName) || helpers.IllegalString(userToCreate.Email) {
		responses.RespondWithError(w, http.StatusBadRequest, "All User Felid Must Be Valid")
		return
	}

	err = userToCreate.SignUp(database)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	viewableResult, _:=userToCreate.ToAdminViewUserInformation()

	responses.RespondWithJSON(w, http.StatusCreated, viewableResult)
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

	authorized, err := currentUser.HasSupervisorRights()

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


func updateUser(w http.ResponseWriter, r *http.Request) {
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

	authorized, err := currentUser.HasAdministratorRights()

	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !authorized {
		responses.RespondWithError(w, http.StatusForbidden, "Insufficient Permissions")
		return
	}

	
	u.UserId, err = strconv.Atoi(r.PostFormValue("userId"))
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Id")
		return
	}
	u.UserName = r.PostFormValue("userName")
	u.FirstName = r.PostFormValue("firstName")
	u.LastName = r.PostFormValue("lastName")
	u.Email = r.PostFormValue("email")
	u.Role, err = strconv.Atoi(r.PostFormValue("role"))
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Role")
		return
	}

	u.Update(database)
	responses.RespondWithJSON(w, http.StatusCreated, "Password Changed")
}