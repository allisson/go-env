# go-env

[![Build Status](https://github.com/allisson/go-env/workflows/Release/badge.svg)](https://github.com/allisson/go-env/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/allisson/go-env)](https://goreportcard.com/report/github.com/allisson/go-env)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/allisson/go-env)

A simple, type-safe Go library for reading environment variables with default values and automatic type conversion.

## Features

- Type-safe environment variable parsing
- Default value support for all types
- Comprehensive type coverage: strings, integers, floats, booleans, durations, and more
- Slice support with custom separators
- Base64 decoding utilities
- Zero dependencies (uses only Go standard library)
- Simple, intuitive API

## Installation

```bash
go get github.com/allisson/go-env
```

## Quick Start

```go
package main

import (
	"fmt"
	"time"

	"github.com/allisson/go-env"
)

func main() {
	// Read string with default value
	apiURL := env.GetString("API_URL", "https://api.example.com")
	fmt.Println("API URL:", apiURL)

	// Read integer with default value
	port := env.GetInt("PORT", 8080)
	fmt.Println("Port:", port)

	// Read boolean with default value
	debug := env.GetBool("DEBUG", false)
	fmt.Println("Debug mode:", debug)

	// Read duration with default value
	timeout := env.GetDuration("TIMEOUT", 30, time.Second)
	fmt.Println("Timeout:", timeout)
}
```

## Usage Examples

### String Functions

#### GetString

```go
// Get a string value
apiKey := env.GetString("API_KEY", "default-key")

// If API_KEY is not set, returns "default-key"
// If API_KEY="my-secret-key", returns "my-secret-key"
```

#### GetStringSlice

```go
// Get a comma-separated list of strings
hosts := env.GetStringSlice("ALLOWED_HOSTS", ",", []string{"localhost"})

// If ALLOWED_HOSTS is not set, returns []string{"localhost"}
// If ALLOWED_HOSTS="example.com,api.example.com", returns []string{"example.com", "api.example.com"}

// Use different separator (e.g., semicolon)
paths := env.GetStringSlice("SEARCH_PATHS", ";", []string{"/usr/local/bin"})
// If SEARCH_PATHS="/usr/bin;/usr/local/bin;/opt/bin", returns []string{"/usr/bin", "/usr/local/bin", "/opt/bin"}
```

### Integer Functions

#### GetInt

```go
// Get an integer value
maxConnections := env.GetInt("MAX_CONNECTIONS", 100)

// If MAX_CONNECTIONS is not set, returns 100
// If MAX_CONNECTIONS="500", returns 500
// If MAX_CONNECTIONS="invalid", returns 100 (default)
```

#### GetIntSlice

```go
// Get a comma-separated list of integers
ports := env.GetIntSlice("PORTS", ",", []int{8080})

// If PORTS is not set, returns []int{8080}
// If PORTS="8080,8081,8082", returns []int{8080, 8081, 8082}
// If PORTS="8080,invalid,8082", returns []int{8080} (default, due to parse error)
```

#### GetInt8, GetInt16, GetInt32, GetInt64

```go
// Get specific integer sizes
age := env.GetInt8("USER_AGE", 25)           // int8: -128 to 127
year := env.GetInt16("YEAR", 2024)           // int16: -32,768 to 32,767
population := env.GetInt32("POPULATION", 1000000) // int32: -2^31 to 2^31-1
bigNumber := env.GetInt64("BIG_NUMBER", 9223372036854775807) // int64: -2^63 to 2^63-1

// Slice variants
ages := env.GetInt8Slice("AGES", ",", []int8{25, 30})
years := env.GetInt16Slice("YEARS", ",", []int16{2020, 2021})
populations := env.GetInt32Slice("POPULATIONS", ",", []int32{1000000})
bigNumbers := env.GetInt64Slice("BIG_NUMBERS", ",", []int64{100000000000})
```

### Unsigned Integer Functions

#### GetUint, GetUint8, GetUint16, GetUint32, GetUint64

```go
// Get unsigned integer values (no negative numbers)
bufferSize := env.GetUint("BUFFER_SIZE", 1024)
statusCode := env.GetUint8("STATUS_CODE", 200)
port := env.GetUint16("PORT", 8080)
maxSize := env.GetUint32("MAX_SIZE", 4294967295)
fileSize := env.GetUint64("FILE_SIZE", 1099511627776)

// Slice variants
sizes := env.GetUintSlice("SIZES", ",", []uint{1024, 2048})
codes := env.GetUint8Slice("CODES", ",", []uint8{200, 201, 204})
ports := env.GetUint16Slice("PORTS", ",", []uint16{8080, 8081})
limits := env.GetUint32Slice("LIMITS", ",", []uint32{1000, 2000})
quotas := env.GetUint64Slice("QUOTAS", ",", []uint64{1000000000})
```

### Boolean Functions

#### GetBool

```go
// Get a boolean value
debug := env.GetBool("DEBUG", false)
enableCache := env.GetBool("ENABLE_CACHE", true)

// Accepted true values: "1", "t", "T", "TRUE", "true", "True"
// Accepted false values: "0", "f", "F", "FALSE", "false", "False"

// If DEBUG is not set, returns false
// If DEBUG="true", returns true
// If DEBUG="1", returns true
// If DEBUG="invalid", returns false (default)
```

#### GetBoolSlice

```go
// Get a comma-separated list of booleans
flags := env.GetBoolSlice("FEATURE_FLAGS", ",", []bool{false, false})

// If FEATURE_FLAGS is not set, returns []bool{false, false}
// If FEATURE_FLAGS="true,false,true", returns []bool{true, false, true}
// If FEATURE_FLAGS="1,0,1", returns []bool{true, false, true}
```

### Float Functions

#### GetFloat32 and GetFloat64

```go
// Get floating-point values
temperature := env.GetFloat32("TEMPERATURE", 20.5)
pi := env.GetFloat64("PI", 3.14159265359)

// If TEMPERATURE is not set, returns 20.5
// If TEMPERATURE="25.7", returns 25.7
// If TEMPERATURE="invalid", returns 20.5 (default)

// Slice variants
temperatures := env.GetFloat32Slice("TEMPERATURES", ",", []float32{20.5, 21.0})
coordinates := env.GetFloat64Slice("COORDINATES", ",", []float64{-73.935242, 40.730610})

// If COORDINATES="40.7128,-74.0060", returns []float64{40.7128, -74.0060}
```

### Byte Functions

#### GetBytes

```go
// Get raw bytes from environment variable
secretKey := env.GetBytes("SECRET_KEY", []byte("default-secret"))

// If SECRET_KEY is not set, returns []byte("default-secret")
// If SECRET_KEY="my-secret", returns []byte("my-secret")

// Useful for binary data or when you need byte manipulation
data := env.GetBytes("RAW_DATA", []byte{})
```

### Duration Functions

#### GetDuration

```go
import "time"

// Get duration values
timeout := env.GetDuration("TIMEOUT", 30, time.Second)
interval := env.GetDuration("POLL_INTERVAL", 5, time.Minute)
ttl := env.GetDuration("CACHE_TTL", 1, time.Hour)

// If TIMEOUT is not set, returns 30 * time.Second (30 seconds)
// If TIMEOUT="60", returns 60 * time.Second (1 minute)

// Common duration units:
// time.Nanosecond
// time.Microsecond
// time.Millisecond
// time.Second
// time.Minute
// time.Hour

// Example: Read milliseconds
responseTimeout := env.GetDuration("RESPONSE_TIMEOUT_MS", 500, time.Millisecond)
// If RESPONSE_TIMEOUT_MS="1000", returns 1000ms (1 second)
```

### Base64 Functions

#### GetBase64ToBytes

```go
// Decode base64 string to bytes
encryptionKey := env.GetBase64ToBytes("ENCRYPTION_KEY", []byte("default-key"))

// If ENCRYPTION_KEY is not set, returns []byte("default-key")
// If ENCRYPTION_KEY="aGVsbG8gd29ybGQ=", returns []byte("hello world")
// If ENCRYPTION_KEY="invalid-base64", returns []byte("default-key")

// Useful for storing binary data in environment variables
certificate := env.GetBase64ToBytes("TLS_CERT", []byte{})
```

#### GetBase64ToString

```go
// Decode base64 string to string
password := env.GetBase64ToString("DB_PASSWORD", "default-password")

// If DB_PASSWORD is not set, returns "default-password"
// If DB_PASSWORD="c2VjcmV0", returns "secret"
// If DB_PASSWORD="invalid-base64", returns "default-password"

// Useful for storing sensitive strings in base64 format
apiToken := env.GetBase64ToString("API_TOKEN", "")
```

## Complete Example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/allisson/go-env"
)

type Config struct {
	// Server configuration
	ServerHost string
	ServerPort int
	Debug      bool

	// Database configuration
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBTimeout  time.Duration

	// Cache configuration
	CacheEnabled bool
	CacheTTL     time.Duration

	// API configuration
	APIKeys        []string
	RateLimits     []int
	TrustedProxies []string

	// Feature flags
	EnableMetrics    bool
	EnableTracing    bool
	EnableCompression bool

	// Advanced settings
	MaxConnections   int32
	BufferSize       uint32
	RequestTimeout   time.Duration
	RetryAttempts    int8
	BackoffMultiplier float64
}

func LoadConfig() *Config {
	return &Config{
		// Server configuration
		ServerHost: env.GetString("SERVER_HOST", "0.0.0.0"),
		ServerPort: env.GetInt("SERVER_PORT", 8080),
		Debug:      env.GetBool("DEBUG", false),

		// Database configuration
		DBHost:     env.GetString("DB_HOST", "localhost"),
		DBPort:     env.GetInt("DB_PORT", 5432),
		DBName:     env.GetString("DB_NAME", "myapp"),
		DBUser:     env.GetString("DB_USER", "postgres"),
		DBPassword: env.GetBase64ToString("DB_PASSWORD", "postgres"),
		DBTimeout:  env.GetDuration("DB_TIMEOUT", 10, time.Second),

		// Cache configuration
		CacheEnabled: env.GetBool("CACHE_ENABLED", true),
		CacheTTL:     env.GetDuration("CACHE_TTL", 5, time.Minute),

		// API configuration
		APIKeys:        env.GetStringSlice("API_KEYS", ",", []string{}),
		RateLimits:     env.GetIntSlice("RATE_LIMITS", ",", []int{100, 1000, 10000}),
		TrustedProxies: env.GetStringSlice("TRUSTED_PROXIES", ",", []string{"127.0.0.1"}),

		// Feature flags
		EnableMetrics:    env.GetBool("ENABLE_METRICS", true),
		EnableTracing:    env.GetBool("ENABLE_TRACING", false),
		EnableCompression: env.GetBool("ENABLE_COMPRESSION", true),

		// Advanced settings
		MaxConnections:   env.GetInt32("MAX_CONNECTIONS", 1000),
		BufferSize:       env.GetUint32("BUFFER_SIZE", 8192),
		RequestTimeout:   env.GetDuration("REQUEST_TIMEOUT", 30, time.Second),
		RetryAttempts:    env.GetInt8("RETRY_ATTEMPTS", 3),
		BackoffMultiplier: env.GetFloat64("BACKOFF_MULTIPLIER", 1.5),
	}
}

func main() {
	// Set some example environment variables
	os.Setenv("SERVER_PORT", "3000")
	os.Setenv("DEBUG", "true")
	os.Setenv("DB_PASSWORD", "bXlzZWNyZXRwYXNzd29yZA==") // base64: "mysecretpassword"
	os.Setenv("API_KEYS", "key1,key2,key3")
	os.Setenv("CACHE_TTL", "10")
	os.Setenv("MAX_CONNECTIONS", "5000")

	// Load configuration
	config := LoadConfig()

	// Display configuration
	fmt.Printf("Server: %s:%d (Debug: %v)\n", config.ServerHost, config.ServerPort, config.Debug)
	fmt.Printf("Database: %s@%s:%d/%s (Timeout: %v)\n",
		config.DBUser, config.DBHost, config.DBPort, config.DBName, config.DBTimeout)
	fmt.Printf("Cache: Enabled=%v, TTL=%v\n", config.CacheEnabled, config.CacheTTL)
	fmt.Printf("API Keys: %v\n", config.APIKeys)
	fmt.Printf("Max Connections: %d\n", config.MaxConnections)
	fmt.Printf("Request Timeout: %v\n", config.RequestTimeout)
}
```

Output:
```
Server: 0.0.0.0:3000 (Debug: true)
Database: postgres@localhost:5432/myapp (Timeout: 10s)
Cache: Enabled=true, TTL=10m0s
API Keys: [key1 key2 key3]
Max Connections: 5000
Request Timeout: 30s
```

## Error Handling

All functions in this library handle errors gracefully by returning the default value when:

- The environment variable is not set
- The value cannot be parsed to the expected type
- For slice functions, if any element cannot be parsed

This makes it safe to use in production without explicit error checking, as you always get a valid value (either from the environment or the default).

```go
// These all safely return defaults on error
port := env.GetInt("INVALID_PORT", 8080)        // Returns 8080 if INVALID_PORT="abc"
enabled := env.GetBool("INVALID_BOOL", false)    // Returns false if INVALID_BOOL="xyz"
nums := env.GetIntSlice("INVALID_NUMS", ",", []int{1}) // Returns []int{1} if parsing fails
```

## Best Practices

1. **Always provide sensible defaults**: This ensures your application can start even without environment variables configured.

```go
// Good
timeout := env.GetDuration("TIMEOUT", 30, time.Second)

// Less ideal (but still works)
timeout := env.GetDuration("TIMEOUT", 0, time.Second) // 0 might not be a good default
```

2. **Use appropriate types**: Choose the smallest type that fits your needs.

```go
// For values 0-255, use uint8
statusCode := env.GetUint8("HTTP_STATUS", 200)

// For port numbers, use uint16 or int
port := env.GetUint16("PORT", 8080)

// For potentially large values, use int32/int64
userCount := env.GetInt64("TOTAL_USERS", 0)
```

3. **Validate critical values**: While go-env handles type conversion, you should still validate business logic.

```go
port := env.GetInt("PORT", 8080)
if port < 1024 || port > 65535 {
	log.Fatal("PORT must be between 1024 and 65535")
}
```

4. **Use base64 for sensitive data**: This provides a level of obfuscation (not encryption).

```go
apiSecret := env.GetBase64ToString("API_SECRET", "")
if apiSecret == "" {
	log.Fatal("API_SECRET is required")
}
```

5. **Document your environment variables**: Create a `.env.example` file in your repository.

```bash
# .env.example
SERVER_PORT=8080
DEBUG=false
DB_HOST=localhost
DB_PASSWORD=base64encodedpassword
API_KEYS=key1,key2,key3
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
