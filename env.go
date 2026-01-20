// Package env provides utilities for reading environment variables with type
// conversion and default value support. It handles various Go types including
// strings, integers (signed and unsigned), floats, booleans, durations, and
// their slice variants.
//
// All Get functions follow a consistent pattern: they attempt to read the
// specified environment variable, parse it to the appropriate type, and return
// the default value if the variable is not set or cannot be parsed.
package env

import (
	b64 "encoding/base64"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetString retrieves a string value from the environment variable specified by key.
// If the environment variable is not set, it returns defaultValue.
func GetString(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}

// GetStringSlice retrieves a string slice from the environment variable specified by key,
// splitting the value using sep as the separator. If the environment variable is not set,
// it returns defaultValue.
func GetStringSlice(key, sep string, defaultValue []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []string
	slice = append(slice, strings.Split(val, sep)...)

	return slice
}

// GetInt retrieves an int value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as an integer,
// it returns defaultValue.
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

// GetIntSlice retrieves an int slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as an integer.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetInt8 retrieves an int8 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as an int8,
// it returns defaultValue.
func GetInt8(key string, defaultValue int8) int8 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(val, 10, 8)
	if err != nil {
		return defaultValue
	}

	return int8(result)
}

// GetInt8Slice retrieves an int8 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as an int8.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
func GetInt8Slice(key, sep string, defaultValue []int8) []int8 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []int8
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, int8(result))
	}

	return slice
}

// GetInt16 retrieves an int16 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as an int16,
// it returns defaultValue.
func GetInt16(key string, defaultValue int16) int16 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(val, 10, 16)
	if err != nil {
		return defaultValue
	}

	return int16(result)
}

// GetInt16Slice retrieves an int16 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as an int16.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
func GetInt16Slice(key, sep string, defaultValue []int16) []int16 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []int16
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, int16(result))
	}

	return slice
}

// GetInt32 retrieves an int32 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as an int32,
// it returns defaultValue.
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

// GetInt32Slice retrieves an int32 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as an int32.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetInt64 retrieves an int64 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as an int64,
// it returns defaultValue.
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

// GetInt64Slice retrieves an int64 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as an int64.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetUint retrieves a uint value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a uint,
// it returns defaultValue.
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

// GetUintSlice retrieves a uint slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a uint.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetUint8 retrieves a uint8 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a uint8,
// it returns defaultValue.
func GetUint8(key string, defaultValue uint8) uint8 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return defaultValue
	}

	return uint8(result)
}

// GetUint8Slice retrieves a uint8 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a uint8.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
func GetUint8Slice(key, sep string, defaultValue []uint8) []uint8 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []uint8
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, uint8(result))
	}

	return slice
}

// GetUint16 retrieves a uint16 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a uint16,
// it returns defaultValue.
func GetUint16(key string, defaultValue uint16) uint16 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		return defaultValue
	}

	return uint16(result)
}

// GetUint16Slice retrieves a uint16 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a uint16.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
func GetUint16Slice(key, sep string, defaultValue []uint16) []uint16 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	var slice []uint16
	for _, s := range strings.Split(val, sep) {
		result, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return defaultValue
		}
		slice = append(slice, uint16(result))
	}

	return slice
}

// GetUint32 retrieves a uint32 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a uint32,
// it returns defaultValue.
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

// GetUint32Slice retrieves a uint32 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a uint32.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetUint64 retrieves a uint64 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a uint64,
// it returns defaultValue.
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

// GetUint64Slice retrieves a uint64 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a uint64.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetBool retrieves a bool value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a boolean,
// it returns defaultValue. Accepted boolean values are defined by strconv.ParseBool
// (1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False).
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

// GetBoolSlice retrieves a bool slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a boolean.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetFloat32 retrieves a float32 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a float32,
// it returns defaultValue.
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

// GetFloat32Slice retrieves a float32 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a float32.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetFloat64 retrieves a float64 value from the environment variable specified by key.
// If the environment variable is not set or cannot be parsed as a float64,
// it returns defaultValue.
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

// GetFloat64Slice retrieves a float64 slice from the environment variable specified by key,
// splitting the value using sep as the separator and parsing each element as a float64.
// If the environment variable is not set or any element cannot be parsed, it returns defaultValue.
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

// GetBytes retrieves a byte slice from the environment variable specified by key,
// converting the string value to []byte. If the environment variable is not set,
// it returns defaultValue.
func GetBytes(key string, defaultValue []byte) []byte {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return []byte(val)
}

// GetDuration retrieves an int64 value from the environment variable specified by key
// and converts it to a time.Duration by multiplying with the provided duration unit.
// For example, to read seconds: GetDuration("TIMEOUT", 30, time.Second).
// If the environment variable is not set or cannot be parsed, it uses defaultValue.
func GetDuration(key string, defaultValue int64, duration time.Duration) time.Duration {
	value := GetInt64(key, defaultValue)
	return time.Duration(value) * duration
}

// GetBase64ToBytes retrieves a base64-encoded string from the environment variable
// specified by key and decodes it to a byte slice. If the environment variable is not
// set or the value cannot be decoded as base64, it returns defaultValue.
func GetBase64ToBytes(key string, defaultValue []byte) []byte {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := b64.StdEncoding.DecodeString(val)
	if err != nil {
		return defaultValue
	}

	return result
}

// GetBase64ToString retrieves a base64-encoded string from the environment variable
// specified by key and decodes it to a string. If the environment variable is not
// set or the value cannot be decoded as base64, it returns defaultValue.
func GetBase64ToString(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	result, err := b64.StdEncoding.DecodeString(val)
	if err != nil {
		return defaultValue
	}

	return string(result)
}
