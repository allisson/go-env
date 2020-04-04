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

// GetInt32 returns a int32 value from environment variable or default value
func GetInt32(key string, defaultValue int32) int32 {
	result := int64(0)

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return defaultValue
	}

	return int32(result)
}

// GetInt64 returns a int64 value from environment variable or default value
func GetInt64(key string, defaultValue int64) int64 {
	result := int64(0)

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(val, 10, 64)
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

// GetFloat32 returns a float32 value from environment variable or default value
func GetFloat32(key string, defaultValue float32) float32 {
	result := float64(0)

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return defaultValue
	}

	return float32(result)
}

// GetFloat64 returns a float64 value from environment variable or default value
func GetFloat64(key string, defaultValue float64) float64 {
	result := float64(0)

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defaultValue
	}

	return result
}
