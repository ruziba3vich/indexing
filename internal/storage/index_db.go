package storage

import (
	"database/sql"
	"fmt"
	"time"
)

func IndexDb(db *sql.DB, indexName, columnName string) (time.Duration, error) {
	query := fmt.Sprintf("EXPLAIN ANALYZE CREATE INDEX %s ON People(%s)", indexName, columnName)

	startTime := time.Now()

	_, err := db.Exec(query)
	if err != nil {
		return 0, err
	}
	executionTime := time.Since(startTime)

	return executionTime, nil
}
