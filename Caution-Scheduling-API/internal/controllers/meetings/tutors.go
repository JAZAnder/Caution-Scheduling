package meetings

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/meeting"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	"github.com/gorilla/mux"
)

func archiveMeeting(w http.ResponseWriter, r *http.Request) {
	var c user.SessionCookie

	cookie, err := r.Cookie("key")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not found")
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
			responses.RespondWithError(w, http.StatusUnauthorized, "Session expired")
			return
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}

	authorized, err := currentUser.HasSupervisorRights()

	if !authorized {
		if err.Error() == "not valid user" {
			responses.RespondWithError(w, http.StatusUnauthorized, err.Error())
		} else {
			responses.RespondWithError(w, http.StatusForbidden, err.Error())
		}

	}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid meeting ID")
		return
	}

	if err := meeting.ArchiveMeeting(database, id); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
