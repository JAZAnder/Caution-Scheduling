package hours

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

func createTimeslots(w http.ResponseWriter, r *http.Request) {
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

	

	monday, err := strconv.ParseBool(r.PostFormValue("Monday"))
	if err != nil {responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload"); return; }
	tuesday, err := strconv.ParseBool(r.PostFormValue("Tuesday"))
	if err != nil {responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload"); return; }
	wednesday, err := strconv.ParseBool(r.PostFormValue("Wednesday"))
	if err != nil {responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload"); return; }
	thursday, err := strconv.ParseBool(r.PostFormValue("Thursday"))
	if err != nil {responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload"); return; }
	friday, err := strconv.ParseBool(r.PostFormValue("Friday"))
	if err != nil {responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload"); return; }
	saturday, err := strconv.ParseBool(r.PostFormValue("Saturday"))
	if err != nil {responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload"); return; }
	sunday, err := strconv.ParseBool(r.PostFormValue("Sunday"))
	if err != nil {responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload"); return; }


	newTimeslotMultiDay := hour.TimeslotsMultiDay{
		StartTime: r.PostFormValue("startTime"),
		EndTime: r.PostFormValue("endTime"),
		Monday: monday,
		Tuesday: tuesday,
		Wednesday: wednesday,
		Thursday: thursday,
		Friday: friday,
		Saturday: saturday,
		Sunday: sunday,
	}



	if helpers.IllegalString(newTimeslotMultiDay.StartTime) || helpers.IllegalString(newTimeslotMultiDay.EndTime) {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	timeslots := newTimeslotMultiDay.ToTimeslotArray()


	for _, timeslot := range timeslots {
		err := timeslot.CreateHour(database)
		if err != nil {
			logger.Log(3, "Admin Tasks", "Timeslots", currentUser.UserName, "Failed to Create New Timeslot from" + timeslot.StartTime + " to "+ timeslot.EndTime + " on the" + strconv.Itoa(timeslot.DayOfWeek) + "day of the week")

		}else{
			logger.Log(2, "Admin Tasks", "Timeslots", currentUser.UserName, "Created New Timeslot from" + timeslot.StartTime + " to "+ timeslot.EndTime + " on the" + strconv.Itoa(timeslot.DayOfWeek) + "day of the week")
		}
		
	}

	responses.RespondWithJSON(w, http.StatusCreated, timeslots)
}

func toggleTimeslot(w http.ResponseWriter, r *http.Request) {
	// Checking for Authentication and Authorization
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


	//Doing the Stuff
	var h hour.Hour

	h.Id, err = strconv.Atoi(r.PostFormValue("id")) 

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Cannot Parse Id")
		return
	}

	h.Active, err = strconv.ParseBool(r.PostFormValue("active")) 

	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Cannot Parse active")
		return
	}

	if h.Active {
		h.MakeHourActive(database)
	}else {
		h.MakeHourDisable(database)
	}


	responses.RespondWithJSON(w, http.StatusCreated, h)
}