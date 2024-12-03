package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

type UserHour struct {
    ID        int  `json:"id"`
    HourID    int  `json:"hourId"`
    TutorID   int  `json:"tutorId"`
    IsEnabled bool `json:"isEnabled"`
}

func UpdateTimeslotStatus(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract ID from URL path
        pathParts := strings.Split(r.URL.Path, "/")
        id := pathParts[len(pathParts)-1]

        var timeslot UserHour
        if err := json.NewDecoder(r.Body).Decode(&timeslot); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err := updateTimeslotStatus(db, id, timeslot.IsEnabled)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]bool{"isEnabled": timeslot.IsEnabled})
    }
}

func updateTimeslotStatus(db *sql.DB, id string, isEnabled bool) error {
    query := "UPDATE user_hours SET isEnabled = ? WHERE id = ?"
    _, err := db.Exec(query, isEnabled, id)
    return err
}
