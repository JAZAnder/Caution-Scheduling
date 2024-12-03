package models

import (
	"database/sql"
)

type LogEntry struct {
	ID          int    `json:"id"`
	Level       string `json:"level"`
	Category    string `json:"category"`
	SubCategory string `json:"subCategory"`
	User        string `json:"user"`
	Message     string `json:"message"`
	Timestamp   string `json:"timestamp"`
}

type LogsFilters struct {
	Level       string `form:"level"`
	Category    string `form:"category"`
	SubCategory string `form:"subCategory"`
	User        string `form:"user"`
}

func GetLogEntries(db *sql.DB, filters LogsFilters) ([]LogEntry, error) {
	query := "SELECT id, level, category, subCategory, user, message, timestamp FROM logs"
	args := []interface{}{}

	if filters.Level != "" {
		query += " WHERE level = ?"
		args = append(args, filters.Level)
	}
	if filters.Category != "" {
		query += " AND category = ?"
		args = append(args, filters.Category)
	}
	if filters.SubCategory != "" {
		query += " AND subCategory = ?"
		args = append(args, filters.SubCategory)
	}
	if filters.User != "" {
		query += " AND user = ?"
		args = append(args, filters.User)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []LogEntry
	for rows.Next() {
		var log LogEntry
		if err := rows.Scan(&log.ID, &log.Level, &log.Category, &log.SubCategory, &log.User, &log.Message, &log.Timestamp); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
