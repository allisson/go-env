# go-env
[![Build Status](https://github.com/allisson/go-env/workflows/Release/badge.svg)](https://github.com/allisson/go-env/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/allisson/go-env)](https://goreportcard.com/report/github.com/allisson/go-env)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/allisson/go-env)

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
	os.Setenv("STRING_SLICE", "string1,string2,string3")
	os.Setenv("DURATION", 10, time.Seconds)

	// Get a environment variable
	str := env.GetString("STRING", "default-value")
	fmt.Printf("str=%#v\n", str)

	// Get a not set environment variable (return default value as fallback)
	str2 := env.GetString("STRING2", "default-value-for-string2")
	fmt.Printf("str2=%#v\n", str2)

	// Get a slice environment variable
	str3 := env.GetStringSlice("STRING_SLICE", ",", []string{"default-value"})
	fmt.Printf("str3=%#v\n", str3)

	duration := env.GetDuration("DURATION", 1, time.Seconds)
	fmt.Printf("duration=%v\n", duration)

	// Other functions:
	// GetInt()
	// GetIntSlice()
	// GetInt8()
	// GetInt8Slice()
	// GetInt16()
	// GetInt16Slice()
	// GetInt32()
	// GetInt32Slice()
	// GetInt64()
	// GetInt64Slice()
	// GetUInt()
	// GetUIntSlice()
	// GetUInt8()
	// GetUInt8Slice()
	// GetUInt16()
	// GetUInt16Slice()
	// GetUInt32()
	// GetUInt32Slice()
	// GetUInt64()
	// GetUInt64Slice()
	// GetBool()
	// GetBoolSlice()
	// GetFloat32()
	// GetFloat32Slice()
	// GetFloat64()
	// GetFloat64Slice()
	// GetBytes()
	// GetBase64ToBytes()
	// GetBase64ToString()
}
```

```bash
go run main.go
str="string"
str2="default-value-for-string2"
str3=[]string{"string1", "string2", "string3"}
duration=10h0m0s
```
