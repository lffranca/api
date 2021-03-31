package repository

import (
	"database/sql"
	"fmt"

	"github.com/lffranca/api/environment"
	_ "github.com/mattn/go-sqlite3"
)

// sqliteConnection sqliteConnection
func sqliteConnection(env *environment.Environment) (*sql.DB, error) {
	var pathDB = "./app.db"
	if env.SQLitePath != "" {
		pathDB = env.SQLitePath
	}

	db, err := sql.Open("sqlite3", pathDB)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(queryBase); err != nil {
		return nil, fmt.Errorf("[ERROR] Query base: %v", err)
	}

	return db, nil
}
