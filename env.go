package env

import (
	"os"
	"strconv"
	"strings"
)

// GetString returns a string value from environment variable or default value
func GetString(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}

// GetStringSlice returns a string slice from environment variable or default value
func GetStringSlice(key, sep string, defaultValue []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	slice := []string{}
	for _, s := range strings.Split(val, sep) {
		slice = append(slice, s)
	}

	return slice
}

// GetInt returns a int value from environment variable or default value
func GetInt(key string, defaultValue int) int {
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

// GetIntSlice returns a int slice from environment variable or default value
func GetIntSlice(key, sep string, defaultValue []int) []int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	slice := []int{}
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.Atoi(s)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, result)
	}

	return slice
}

// GetInt32 returns a int32 value from environment variable or default value
func GetInt32(key string, defaultValue int32) int32 {
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

// GetInt32Slice returns a int32 slice from environment variable or default value
func GetInt32Slice(key, sep string, defaultValue []int32) []int32 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	slice := []int32{}
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, int32(result))
	}

	return slice
}

// GetInt64 returns a int64 value from environment variable or default value
func GetInt64(key string, defaultValue int64) int64 {
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

// GetInt64Slice returns a int64 slice from environment variable or default value
func GetInt64Slice(key, sep string, defaultValue []int64) []int64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	slice := []int64{}
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, result)
	}

	return slice
}

// GetBool returns a boolean value from environment variable or default value
func GetBool(key string, defaultValue bool) bool {
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

// GetBoolSlice returns a boolean slice from environment variable or default value
func GetBoolSlice(key, sep string, defaultValue []bool) []bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	slice := []bool{}
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseBool(s)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, result)
	}

	return slice
}

// GetFloat32 returns a float32 value from environment variable or default value
func GetFloat32(key string, defaultValue float32) float32 {
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

// GetFloat32Slice returns a float32 slice from environment variable or default value
func GetFloat32Slice(key, sep string, defaultValue []float32) []float32 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	slice := []float32{}
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, float32(result))
	}

	return slice
}

// GetFloat64 returns a float64 value from environment variable or default value
func GetFloat64(key string, defaultValue float64) float64 {
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

// GetFloat64Slice returns a float64 slice from environment variable or default value
func GetFloat64Slice(key, sep string, defaultValue []float64) []float64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	slice := []float64{}
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, result)
	}

	return slice
}

// GetBytes returns a byte slice value from environment variable or default value
func GetBytes(key string, defaultValue []byte) []byte {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return []byte(val)
}
