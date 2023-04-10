package env

import (
	"os"
	"strconv"
	"strings"
	"time"
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

	var slice []string
	slice = append(slice, strings.Split(val, sep)...)

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

	var slice []int
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

	var slice []int32
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

	var slice []int64
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, result)
	}

	return slice
}

// GetUint returns a uint value from environment variable or default value
func GetUint(key string, defaultValue uint) uint {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(val, 10, 0)
	if err != nil {
		return defaultValue
	}

	return uint(result)
}

// GetUintSlice returns a uint slice from environment variable or default value
func GetUintSlice(key, sep string, defaultValue []uint) []uint {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []uint
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, uint(result))
	}

	return slice
}

// GetUint32 returns a uint32 value from environment variable or default value
func GetUint32(key string, defaultValue uint32) uint32 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return defaultValue
	}

	return uint32(result)
}

// GetUint32Slice returns a uint32 slice from environment variable or default value
func GetUint32Slice(key, sep string, defaultValue []uint32) []uint32 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []uint32
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, uint32(result))
	}

	return slice
}

// GetUint64 returns a uint64 value from environment variable or default value
func GetUint64(key string, defaultValue uint64) uint64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return defaultValue
	}

	return result
}

// GetUint64Slice returns a uint64 slice from environment variable or default value
func GetUint64Slice(key, sep string, defaultValue []uint64) []uint64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []uint64
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseUint(s, 10, 64)
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

	var slice []bool
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

	var slice []float32
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

	var slice []float64
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

// GetDuration returns a time.Duration value from environment variable or default value
func GetDuration(key string, defaultValue int64, duration time.Duration) time.Duration {
	value := GetInt64(key, defaultValue)
	return time.Duration(value) * duration
}
