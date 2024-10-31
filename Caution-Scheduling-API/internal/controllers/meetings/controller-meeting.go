package meetings

import (
	//"database/sql"
	//"errors"
	//"fmt"
	//"net/http"
	//"strconv"

	"github.com/gorilla/mux"

	//"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
	//"github.com/JAZAnder/Caution-Scheduling/internal/objects/meeting"
	//"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"

)

var(
	database = db.GetDatabase()
)

func AddMeetingRoutes(a *mux.Router) {
	// a.HandleFunc("/api/meeting", createMeeting).Methods("POST")
	// a.HandleFunc("/api/meeting/{id:[0-9]+}", getMeeting).Methods("GET")
	// a.HandleFunc("/api/meetings", getMeetings).Methods("GET")
	// a.HandleFunc("/api/meeting/{id:[0-9]+}", deleteMeeting).Methods("DELETE")
	// a.HandleFunc("/api/meetings/mine", getMyMeetings).Methods("GET")
}

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

// func getMyMeetings(w http.ResponseWriter, r *http.Request) {
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

// 	meetings, err := meeting.GetMyMeetings(database, currentUser.UserName)
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	responses.RespondWithJSON(w, http.StatusOK, meetings)
// }

// func createMeeting(w http.ResponseWriter, r *http.Request) {
// 	var m meeting.Meeting
// 	var err error

// 	m.UserHourId, err = strconv.Atoi(r.PostFormValue("userHourId"))
// 	m.LabId, err = strconv.Atoi(r.PostFormValue("labId"))
// 	m.StudentName = r.PostFormValue("studentName")
// 	m.StudentEmail = r.PostFormValue("studentEmail")
// 	m.Date, err = strconv.Atoi(r.PostFormValue("date"))

// 	if err != nil || len(m.StudentName) <= 0 || len(m.StudentEmail) <= 0 {
// 		responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	err = m.CreateMeeting(database)
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	responses.RespondWithJSON(w, http.StatusCreated, m)
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

// func deleteMeeting(w http.ResponseWriter, r *http.Request) {
// 	var c user.SessionCookie

// 	cookie, err := r.Cookie("key")
// 	if err != nil {
// 		if errors.Is(err, http.ErrNoCookie) {
// 			responses.RespondWithError(w, http.StatusUnauthorized, "Cookie not found")
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
// 			responses.RespondWithError(w, http.StatusUnauthorized, "Session expired")
// 			return
// 		} else {
// 			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 	}

// 	if !currentUser.IsAdmin {
// 		fmt.Println("	Fail : Meeting not deleted by " + currentUser.UserName + " : " + "Not an Admin")
// 		responses.RespondWithError(w, http.StatusForbidden, "Not an Admin")
// 		return
// 	}

// 	vars := mux.Vars(r)

// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		responses.RespondWithError(w, http.StatusBadRequest, "Invalid meeting ID")
// 		return
// 	}
// 	m := meeting.Meeting{Id: id}
// 	if err := m.DeleteMeeting(database); err != nil {
// 		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	responses.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
// }
