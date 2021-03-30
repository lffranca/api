package environment

import (
	"errors"
	"os"
)

// Get Get
func Get() (*Environment, error) {
	apiVersion := os.Getenv("API_VERSION")
	apiPort := os.Getenv("API_PORT")
	sqlitePath := os.Getenv("SQLITE_PATH")

	switch "" {
	case apiVersion,
		apiPort:
		return nil, errors.New("[ERROR] Invalid envs var")
	}

	return &Environment{
		APIVersion: apiVersion,
		APIPort:    apiPort,
		SQLitePath: sqlitePath,
	}, nil
}
