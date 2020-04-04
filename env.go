package env

import (
	"os"
	"strconv"
)

// GetString returns a string value from environment variable or default value
func GetString(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}

// GetInt returns a int value from environment variable or default value
func GetInt(key string, defaultValue int) int {
	result := 0

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}

	return result
}

// GetBool returns a boolean value from environment variable or default value
func GetBool(key string, defaultValue bool) bool {
	result := false

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}

	return result
}
