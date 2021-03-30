package repository

import (
	"database/sql"

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

	return db, nil
}
