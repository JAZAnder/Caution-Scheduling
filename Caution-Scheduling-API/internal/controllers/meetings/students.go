package meetings

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/email"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/meeting"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"

)

func createMeeting(w http.ResponseWriter, r *http.Request) {
	var m meeting.Meeting
	var err error
	var c user.SessionCookie

	m.UserHourId, err = strconv.Atoi(r.PostFormValue("userHourId"))
	m.Date, err = strconv.Atoi(r.PostFormValue("date"))
	m.UserHourId, err = strconv.Atoi(r.PostFormValue("userHourId"))
	m.TopicId, err = strconv.Atoi(r.PostFormValue("topicId"))

	if err != nil {
		m.TopicId = 1
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

	authorized, err := currentUser.HasStudentRights()

	if !authorized {
		if err.Error() == "not valid user" {
			responses.RespondWithError(w, http.StatusUnauthorized, err.Error())
		} else {
			responses.RespondWithError(w, http.StatusForbidden, err.Error())
		}

	}

	m.StudentId = currentUser.UserId

	err = m.CreateMeeting(database)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tutor,_ := userHour.GetUserByUserHour(database, m.UserHourId)
	email.NewMeeting(currentUser,tutor , m)

	responses.RespondWithJSON(w, http.StatusCreated, m)

}

func getMyMeetings(w http.ResponseWriter, r *http.Request) {
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

	meetings, err := meeting.GetMyMeetings(database, currentUser.UserId)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, meetings)
}
