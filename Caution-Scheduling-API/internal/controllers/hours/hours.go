package hours

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers"
	db "github.com/JAZAnder/Caution-Scheduling/internal/helpers/database"
	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"

)

type HourController struct {
	DB *sql.DB
}

func AddHourRoutes(a *mux.Router) {
	a.HandleFunc("/api/hour", createHour).Methods("POST")
	a.HandleFunc("/api/hour/{id:[0-9]+}", getHour).Methods("GET")
	a.HandleFunc("/api/hours", getHours).Methods("GET")
	a.HandleFunc("/api/hour/day/{id:[0-9]+}", getHoursByDay).Methods("GET")
	a.HandleFunc("/api/hour/{id:[0-9]+}", deleteHour).Methods("DELETE")
	a.HandleFunc("/api/hour/availability/{id:[0-9]+}", getUsersByHour).Methods("GET")
}

var (
	database = db.GetDatabase()
)

func getHour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid hour Id")
		return
	}
	h := hour.Hour{Id: id}
	err = h.GetHour(database)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			responses.RespondWithError(w, http.StatusNotFound, "Timeslot not Found")
		default:
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, h)
}

func createHour(w http.ResponseWriter, r *http.Request) {
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

	if !currentUser.IsAdmin {
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName + " : " + "Not an Admin")
		responses.RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	var h hour.Hour
	h.StartTime = r.PostFormValue("startTime")
	h.EndTime = r.PostFormValue("endTime")
	h.DayOfWeek, err = strconv.Atoi(r.PostFormValue("dayOfWeek"))

	if err != nil || h.DayOfWeek > 6 || h.DayOfWeek < 0 {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if helpers.IllegalString(h.StartTime) || helpers.IllegalString(h.EndTime) {
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName + " : " + "Invalid request payload")
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.CreateHour(database)
	if err != nil {
		fmt.Println("	Fail : Time Not Created by " + currentUser.UserName + " : " + err.Error())
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("	Time Created by " + currentUser.UserName)

	responses.RespondWithJSON(w, http.StatusCreated, h)
}

func getHours(w http.ResponseWriter, r *http.Request) {
	hours, err := hour.GetHours(database)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, hours)
}

func getHoursByDay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid hour Id")
		return
	}

	hours, err := hour.GetHoursByDay(database, id)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, hours)
}

func deleteHour(w http.ResponseWriter, r *http.Request) {
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

	if !currentUser.IsAdmin {
		fmt.Println("	Fail : Time Not Deleted by " + currentUser.UserName + " : " + "Not an Admin")
		responses.RespondWithError(w, http.StatusForbidden, "Not an Admin")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid hour ID")
		return
	}
	h := hour.Hour{Id: id}
	if err := h.DeleteHour(database); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})

}

func getUsersByHour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid hour Id")
		return
	}

	users, err := user.GetUsersByHour(database, id)
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondWithJSON(w, http.StatusOK, users)
}
