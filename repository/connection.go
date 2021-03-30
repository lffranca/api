package repository

import (
	"errors"

	"github.com/lffranca/api/environment"
)

type ConnectionType int

const (
	ConnectionSQLite  ConnectionType = iota
	ConnectionMongoDB ConnectionType = iota
)

// GetInput GetInput
type GetInput struct {
	Type ConnectionType
	Env  *environment.Environment
}

// Get Get
func Get(input *GetInput) (repositoryItem Repository, err error) {
	repositoryItem = &repository{}

	if db != nil {
		if err = db.Ping(); err == nil {
			return repositoryItem, nil
		}
	}

	switch input.Type {
	case ConnectionSQLite:
		db, err = sqliteConnection(input.Env)
		if err != nil {
			return nil, err
		}

		return repositoryItem, nil
	case ConnectionMongoDB:
		return nil, nil
	}

	return nil, errors.New("[ERROR] Invalid connection type")
}
