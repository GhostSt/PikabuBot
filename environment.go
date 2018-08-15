package bot

import (
	"errors"
	"fmt"
	"os"
)

// Structure provide access to environment variables
type env struct {
	parameters map[string]string
}

func (env *env) get(key string) (string, error) {
	for itemKey := range env.parameters {
		if key == itemKey {
			return env.parameters[key], nil
		}
	}

	value := os.Getenv(key)

	if value != "" {
		return value, nil
	}

	return "", errors.New(fmt.Sprintf("Environment variable %s doesn't exist", key))
}