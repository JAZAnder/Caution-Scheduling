package userTimeslots

import (
	"net/http"
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"
	"github.com/gorilla/mux"
)

func getTutorsAvailability(w http.ResponseWriter, r *http.Request) {
	var uh userHour.UserHour
	var err error
	vars := mux.Vars(r)
	uh.TutorId, err = strconv.Atoi(vars["userId"]) 

	if err != nil {
		//TODO Return non Generic Error
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userHours, err := uh.GetHoursByUserId(database)

	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, userHours)
}