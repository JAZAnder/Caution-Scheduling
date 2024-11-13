package userTimeslots

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"

)

func getTutorsAvailability(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tutorId, err := strconv.Atoi(vars["tutorId"]) //TODO Return non Generic Error
	if err != nil { responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	date, err := strconv.Atoi(vars["date"]) //TODO Return non Generic Error
	if err != nil { responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userHours, err := userHour.GetAvailableHoursByUserAndDay(database, tutorId, date)



	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, userHours)
}

func getTutorsAvailabilityByDateOnly(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	date, err := strconv.Atoi(vars["date"]) //TODO Return non Generic Error
	if err != nil { responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userHours, err := userHour.GetAvailableHoursByDay(database, date)



	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, userHours)
}



