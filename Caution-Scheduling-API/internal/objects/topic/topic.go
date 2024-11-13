package topic

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"

)

func (topic *Topic) AddTopic( db *sql.DB) error {
	query := "INSERT INTO `topic` (`description`) VALUES ('" + topic.Description + "');"
	logger.Log(1, "Database", "Topics", "databaseManager", query)
	errSql := db.QueryRow(query)

	if errSql.Err() != nil {
		return errSql.Err()
	}

	return nil
}
