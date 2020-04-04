package env

import (
	"bytes"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	os.Setenv("STRING2", "STRING2")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  string
		expectedValue string
	}{
		{"test-default-value", "STRING1", "string1", "string1"},
		{"test-value-from-envvar", "STRING2", "string2", "STRING2"},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetString(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetString(\"%s\", \"%s\"): expected %s, actual %s", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("INT2", "2")
	os.Setenv("INT3", "três")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  int
		expectedValue int
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt32(t *testing.T) {
	os.Setenv("INT2", "2")
	os.Setenv("INT3", "três")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  int32
		expectedValue int32
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt32(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt32(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt64(t *testing.T) {
	os.Setenv("INT2", "2")
	os.Setenv("INT3", "três")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  int64
		expectedValue int64
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt64(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt64(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	os.Setenv("BOOL2", "true")
	os.Setenv("BOOL3", "tru")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  bool
		expectedValue bool
	}{
		{"test-default-value", "BOOL1", true, true},
		{"test-value-from-envvar", "BOOL2", false, true},
		{"test-invalid-value-from-envvar", "BOOL3", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetBool(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetBool(\"%s\", %t): expected %t, actual %t", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetFloat32(t *testing.T) {
	os.Setenv("FLOAT2", "2.1")
	os.Setenv("FLOAT3", "três-ponto-um")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  float32
		expectedValue float32
	}{
		{"test-default-value", "FLOAT1", 1.5, 1.5},
		{"test-value-from-envvar", "FLOAT2", 1, 2.1},
		{"test-invalid-value-from-envvar", "FLOAT3", 3.1, 3.1},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetFloat32(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetFloat32(\"%s\", %b): expected %b, actual %b", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	os.Setenv("FLOAT2", "2.1")
	os.Setenv("FLOAT3", "três-ponto-um")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  float64
		expectedValue float64
	}{
		{"test-default-value", "FLOAT1", 1.5, 1.5},
		{"test-value-from-envvar", "FLOAT2", 1, 2.1},
		{"test-invalid-value-from-envvar", "FLOAT3", 3.1, 3.1},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetFloat64(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetFloat64(\"%s\", %b): expected %b, actual %b", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetBytes(t *testing.T) {
	os.Setenv("STRING2", "STRING2")

	var tests = []struct {
		kind          string
		key           string
		defaultValue  []byte
		expectedValue []byte
	}{
		{"test-default-value", "STRING1", []byte("string1"), []byte("string1")},
		{"test-value-from-envvar", "STRING2", []byte("string2"), []byte("STRING2")},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetBytes(tt.key, tt.defaultValue)
			if bytes.Compare(result, tt.expectedValue) != 0 {
				t.Errorf("GetBytes(\"%s\", %#v): expected %#v, actual %#v", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}
