package logger

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetLogs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filters := handlers.LogsFilters{
			Level:       r.URL.Query().Get("level"),
			Category:    r.URL.Query().Get("category"),
			SubCategory: r.URL.Query().Get("subCategory"),
			User:        r.URL.Query().Get("user"),
		}

		logs, err := handlers.GetLogEntries(db, filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(logs)
	}
}
