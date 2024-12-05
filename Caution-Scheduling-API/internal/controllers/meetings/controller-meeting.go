package meetings

import (
	//"database/sql"
	//"errors"
	//"fmt"
	//"net/http"
	//"strconv"
	// "database/sql"
	// "errors"
	// "fmt"
	// "net/http"

	//"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
	//"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	//"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	//"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"

)

var (
	database = db.GetDatabase()
)

// func getMeeting(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusBadRequest, "Invalid meeting Id")
// 		return
// 	}
// 	m := meeting.Meeting{Id: id}
// 	err = m.GetMeeting(database)
// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			responses.RespondWithError(w, http.StatusNotFound, "Meeting not found")
// 		default:
// 			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	responses.RespondWithJSON(w, http.StatusOK, m)
// }

// func getMeetings(w http.ResponseWriter, r *http.Request) {
// 	var c user.SessionCookie

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

// 	if !currentUser.IsAdmin {
// 		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName + " : " + "Not an Admin")
// 		responses.RespondWithError(w, http.StatusForbidden, "Not an Admin")
// 		return
// 	}

// 	meetings, err := meeting.GetMeetings(database)
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	responses.RespondWithJSON(w, http.StatusOK, meetings)
// }


