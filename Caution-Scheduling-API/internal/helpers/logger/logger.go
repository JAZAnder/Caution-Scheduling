package logger

import (
	"database/sql"
	"errors"
	"fmt"

	"time"
)

// LogLevel represents the severity of the log
type LogLevel int

const (
	Debug LogLevel = iota + 1
	Info
	Warning
	Error
	Critical
	NoLogging
)

// Config holds the logger configuration
type Config struct {
	MinLogLevelConsole LogLevel
	MinLogLevelDB      LogLevel
	Database           *sql.DB
}

// Logger represents the logger instance
type Logger struct {
	config Config
}

// New creates a new Logger instance
func New(config Config) *Logger {
	return &Logger{config: config}
}

// String converts LogLevel to string
func (l LogLevel) String() string {
	return [...]string{"Debug", "Info", "Warning", "Error", "Critical", "NoLogging"}[l-1]
}

// Log sends a log message
func (l *Logger) Log(level LogLevel, category, subCategory, user, message string) error {
	if err := l.validate(level, category, subCategory); err != nil {
		return err
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	if level >= l.config.MinLogLevelConsole {
		if err := l.logToConsole(currentTime, level, category, subCategory, user, message); err != nil {
			return fmt.Errorf("console logging failed: %w", err)
		}
	}

	if level >= l.config.MinLogLevelDB && l.config.Database != nil {
		if err := l.logToDB(currentTime, level, category, subCategory, user, message); err != nil {
			return fmt.Errorf("database logging failed: %w", err)
		}
	}

	return nil
}

func (l *Logger) validate(level LogLevel, category, subCategory string) error {
	if level < Debug || level > NoLogging {
		return errors.New("invalid log level")
	}
	if len(category) > 20 {
		return errors.New("category is too long (max 20 characters)")
	}
	if len(subCategory) > 20 {
		return errors.New("subCategory is too long (max 20 characters)")
	}
	return nil
}

func LogSetUpDb(db *sql.DB) {

}

func (l *Logger) logToConsole(timestamp string, level LogLevel, category, subCategory, user, message string) error {
	fmt.Printf("\n %s || %s || Category: %s - %s || %s\n\t%s\n\n",
		level.String(), timestamp, category, subCategory, user, message)
	return nil
}

func (l *Logger) logToDB(timestamp string, level LogLevel, category, subCategory, user, message string) error {
	query := `
		INSERT INTO logs (timestamp, level, category, subcategory, user, message)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := l.config.Database.Exec(query, timestamp, level.String(), category, subCategory, user, message)
	return err
}

// GetMeetingsByTutor retrieves meetings from the database filtered by tutor
func (l *Logger) GetMeetingsByTutor(tutor string) ([]Meeting, error) {
	if l.config.Database == nil {
		return nil, errors.New("database connection not initialized")
	}

	query := `
		SELECT timestamp, category, subcategory, message
		FROM logs
		WHERE category = 'Meeting' AND user = ?
		ORDER BY timestamp DESC
	`

	rows, err := l.config.Database.Query(query, tutor)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var meetings []Meeting
	for rows.Next() {
		var m Meeting
		err := rows.Scan(&m.Timestamp, &m.Category, &m.SubCategory, &m.Message)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		meetings = append(meetings, m)
	}

	return meetings, nil
}

// Meeting represents a meeting log entry
type Meeting struct {
	Timestamp   string
	Category    string
	SubCategory string
	Message     string
}
