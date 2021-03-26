package environment

import (
	"errors"
	"os"
)

// Get Get
func Get() (*Environment, error) {
	apiVersion := os.Getenv("API_VERSION")
	apiPort := os.Getenv("API_PORT")

	switch "" {
	case apiVersion,
		apiPort:
		return nil, errors.New("[ERROR] Invalid envs var")
	}

	return &Environment{
		APIVersion: apiVersion,
		APIPort:    apiPort,
	}, nil
}
