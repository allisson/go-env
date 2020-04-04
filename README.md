# go-env
[![Build Status](https://github.com/allisson/go-env/workflows/tests/badge.svg)](https://github.com/allisson/go-env/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/allisson/go-env)](https://goreportcard.com/report/github.com/allisson/go-env)
[![Documentation](https://godoc.org/github.com/allisson/go-env?status.svg)](http://godoc.org/github.com/allisson/go-env)

A simple golang library to get values from environment variables.

Example:

```golang
package main

import (
	"fmt"
	"os"

	"github.com/allisson/go-env"
)

func main() {
	// Set a environment variable
	os.Setenv("STRING", "string")

	// Get a environment variable
	str := env.GetString("STRING", "default-value")
	fmt.Printf("str=%#v\n", str)

	// Get a not set environment variable (return default value as fallback)
	str2 := env.GetString("STRING2", "default-value-for-string2")
	fmt.Printf("str2=%#v\n", str2)

	// Other functions:
	// GetInt()
	// GetInt32()
	// GetInt64()
	// GetBool()
	// GetFloat32()
	// GetFloat64()
	// GetBytes()
}
```

```bash
go run main.go
str="string"
str2="default-value-for-string2"
```
