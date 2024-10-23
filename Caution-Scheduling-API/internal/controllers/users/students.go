package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

func changePassword(w http.ResponseWriter, r *http.Request) {
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

	u.UserName = user.UserName
	u.Password = r.PostFormValue("oldPassword")
	err = u.Login(database)

	if err != nil {
		responses.RespondWithError(w, http.StatusUnauthorized, "Password is Incorrect")
		return
	}

	u.Password = r.PostFormValue("newPassword")
	u.ChangePassword(database)

	responses.RespondWithJSON(w, http.StatusCreated, "Password Changed")

}