package repository

import "database/sql"

var db *sql.DB

// Repository Repository
type Repository interface {
	Close() error
}

type repository struct{}

func (item *repository) Close() error {
	return db.Close()
}
