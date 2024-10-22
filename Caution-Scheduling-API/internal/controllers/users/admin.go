package users

import (
	"net/http"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var c user.SessionCookie
	isAdmin := false
	cookie, err := r.Cookie("key")
	if err == nil {
		c.Cookie = cookie.Value
		currentUser, err := c.CheckSession(database)
		if err == nil {
			isAdmin = currentUser.IsAdmin
		}
	}

	Users, err := user.GetLusers(database, isAdmin)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, Users)
}