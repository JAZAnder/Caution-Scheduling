package userTimeslots

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"
)

func getEverythingByFilter(w http.ResponseWriter, r *http.Request) {

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


	filter := userHour.TutorsAndHours{
		HourId: r.URL.Query().Get("hourId"),
		TutorId: r.URL.Query().Get("tutorId"),
		DayOfWeek: r.URL.Query().Get("dayOfWeek"),
	}

	UsersAndTimeSlots, err := userHour.GetUserTimeslotByFilter(database, filter)

	if err != nil {
		//TODO Return non Generic Error
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, UsersAndTimeSlots)
}

func addTimeAdmin(w http.ResponseWriter, r *http.Request) {
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

	authorized, err := currentUser.HasSupervisorRights()

	if (!authorized){
		if err.Error() == "not valid user" {
			responses.RespondWithError(w, http.StatusUnauthorized, err.Error())
		}else{
			responses.RespondWithError(w, http.StatusForbidden, err.Error())
		}
		
	}


	hourId := r.PostFormValue("timeslotId")
	uh.HourId, err = strconv.Atoi(hourId)
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}

	userId := r.PostFormValue("userId")
	uh.TutorId, err = strconv.Atoi(userId)
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid Time Slot")
		return
	}

	err = uh.CreateUserHour(database)

	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	responses.RespondWithJSON(w, http.StatusCreated, uh)

}