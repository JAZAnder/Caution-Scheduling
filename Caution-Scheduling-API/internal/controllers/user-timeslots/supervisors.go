package userTimeslots

import (
	"net/http"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"

)

func getEverythingByFilter(w http.ResponseWriter, r *http.Request) {

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