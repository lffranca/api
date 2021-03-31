package repository

import (
	"database/sql"
	"errors"
)

// GetDB GetDB
func GetDB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("")
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
