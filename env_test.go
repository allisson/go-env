package env

import (
	"bytes"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestGetString(t *testing.T) {
	os.Setenv("STRING2", "STRING2") //nolint:errcheck

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

func TestGetStringSlice(t *testing.T) {
	os.Setenv("STRING2", "string1,string2")         //nolint:errcheck
	os.Setenv("STRING3", "string1 string2 string3") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []string
		expectedValue []string
	}{
		{"test-default-value-sep-comma", "STRING1", ",", []string{"string1"}, []string{"string1"}},
		{"test-value-from-envvar-sep-comma", "STRING2", ",", []string{"string2"}, []string{"string1", "string2"}},
		{"test-value-from-envvar-sep-space", "STRING3", " ", []string{"string3"}, []string{"string1", "string2", "string3"}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetStringSlice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetStringSlice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

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

func TestGetIntSlice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []int
		expectedValue []int
	}{
		{"test-default-value-sep-comma", ",", "INT1", []int{1}, []int{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []int{1}, []int{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []int{1}, []int{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []int{3}, []int{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetIntSlice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetIntSlice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt8(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  int8
		expectedValue int8
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt8(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt8(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt8Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []int8
		expectedValue []int8
	}{
		{"test-default-value-sep-comma", ",", "INT1", []int8{1}, []int8{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []int8{1}, []int8{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []int8{1}, []int8{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []int8{3}, []int8{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt8Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt8Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt16(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  int16
		expectedValue int16
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt16(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt16(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt16Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []int16
		expectedValue []int16
	}{
		{"test-default-value-sep-comma", ",", "INT1", []int16{1}, []int16{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []int16{1}, []int16{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []int16{1}, []int16{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []int16{3}, []int16{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt16Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt16Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt32(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

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

func TestGetInt32Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []int32
		expectedValue []int32
	}{
		{"test-default-value-sep-comma", ",", "INT1", []int32{1}, []int32{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []int32{1}, []int32{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []int32{1}, []int32{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []int32{3}, []int32{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt32Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt32Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetInt64(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

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

func TestGetInt64Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []int64
		expectedValue []int64
	}{
		{"test-default-value-sep-comma", ",", "INT1", []int64{1}, []int64{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []int64{1}, []int64{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []int64{1}, []int64{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []int64{3}, []int64{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetInt64Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt64Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  uint
		expectedValue uint
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt64(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUintSlice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []uint
		expectedValue []uint
	}{
		{"test-default-value-sep-comma", ",", "INT1", []uint{1}, []uint{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []uint{1}, []uint{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []uint{1}, []uint{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []uint{3}, []uint{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUintSlice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt64Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint8(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  uint8
		expectedValue uint8
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint8(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt8(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint8Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []uint8
		expectedValue []uint8
	}{
		{"test-default-value-sep-comma", ",", "INT1", []uint8{1}, []uint8{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []uint8{1}, []uint8{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []uint8{1}, []uint8{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []uint8{3}, []uint8{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint8Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt8Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint16(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  uint16
		expectedValue uint16
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint16(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt16(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint16Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []uint16
		expectedValue []uint16
	}{
		{"test-default-value-sep-comma", ",", "INT1", []uint16{1}, []uint16{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []uint16{1}, []uint16{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []uint16{1}, []uint16{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []uint16{3}, []uint16{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint16Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt16Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint32(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  uint32
		expectedValue uint32
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint32(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt64(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint32Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []uint32
		expectedValue []uint32
	}{
		{"test-default-value-sep-comma", ",", "INT1", []uint32{1}, []uint32{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []uint32{1}, []uint32{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []uint32{1}, []uint32{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []uint32{3}, []uint32{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint32Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt64Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint64(t *testing.T) {
	os.Setenv("INT2", "2")    //nolint:errcheck
	os.Setenv("INT3", "três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  uint64
		expectedValue uint64
	}{
		{"test-default-value", "INT1", 1, 1},
		{"test-value-from-envvar", "INT2", 1, 2},
		{"test-invalid-value-from-envvar", "INT3", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint64(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetInt64(\"%s\", %d): expected %d, actual %d", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetUint64Slice(t *testing.T) {
	os.Setenv("INT2", "1,2")      //nolint:errcheck
	os.Setenv("INT3", "1 2 3")    //nolint:errcheck
	os.Setenv("INT4", "1,2,três") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []uint64
		expectedValue []uint64
	}{
		{"test-default-value-sep-comma", ",", "INT1", []uint64{1}, []uint64{1}},
		{"test-value-from-envvar-sep-comma", "INT2", ",", []uint64{1}, []uint64{1, 2}},
		{"test-value-from-envvar-sep-space", "INT3", " ", []uint64{1}, []uint64{1, 2, 3}},
		{"test-invalid-value-from-envvar", "INT4", ",", []uint64{3}, []uint64{3}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetUint64Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetInt64Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	os.Setenv("BOOL2", "true") //nolint:errcheck
	os.Setenv("BOOL3", "tru")  //nolint:errcheck

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

func TestGetBoolSlice(t *testing.T) {
	os.Setenv("BOOL2", "true,true")       //nolint:errcheck
	os.Setenv("BOOL3", "true true false") //nolint:errcheck
	os.Setenv("BOOL4", "true,true,falso") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []bool
		expectedValue []bool
	}{
		{"test-default-value-sep-comma", "BOOL1", ",", []bool{true}, []bool{true}},
		{"test-value-from-envvar-sep-comma", "BOOL2", ",", []bool{false}, []bool{true, true}},
		{"test-value-from-envvar-sep-space", "BOOL3", " ", []bool{false}, []bool{true, true, false}},
		{"test-invalid-value-from-envvar", "BOOL4", ",", []bool{true}, []bool{true}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetBoolSlice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetBoolSlice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetFloat32(t *testing.T) {
	os.Setenv("FLOAT2", "2.1")           //nolint:errcheck
	os.Setenv("FLOAT3", "três-ponto-um") //nolint:errcheck

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

func TestGetFloat32Slice(t *testing.T) {
	os.Setenv("FLOAT2", "2.1,2.2")                 //nolint:errcheck
	os.Setenv("FLOAT3", "2.3 2.4 2.5")             //nolint:errcheck
	os.Setenv("FLOAT4", "2.6,2.7,dois-ponto-oito") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []float32
		expectedValue []float32
	}{
		{"test-default-value-sep-comma", "FLOAT1", ",", []float32{2.0}, []float32{2.0}},
		{"test-value-from-envvar-sep-comma", "FLOAT2", ",", []float32{1.0}, []float32{2.1, 2.2}},
		{"test-value-from-envvar-sep-space", "FLOAT3", " ", []float32{1.0}, []float32{2.3, 2.4, 2.5}},
		{"test-invalid-value-from-envvar", "FLOAT4", ",", []float32{1.0}, []float32{1.0}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetFloat32Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetFloat32Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	os.Setenv("FLOAT2", "2.1")           //nolint:errcheck
	os.Setenv("FLOAT3", "três-ponto-um") //nolint:errcheck

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

func TestGetFloat64Slice(t *testing.T) {
	os.Setenv("FLOAT2", "2.1,2.2")                 //nolint:errcheck
	os.Setenv("FLOAT3", "2.3 2.4 2.5")             //nolint:errcheck
	os.Setenv("FLOAT4", "2.6,2.7,dois-ponto-oito") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		sep           string
		defaultValue  []float64
		expectedValue []float64
	}{
		{"test-default-value-sep-comma", "FLOAT1", ",", []float64{2.0}, []float64{2.0}},
		{"test-value-from-envvar-sep-comma", "FLOAT2", ",", []float64{1.0}, []float64{2.1, 2.2}},
		{"test-value-from-envvar-sep-space", "FLOAT3", " ", []float64{1.0}, []float64{2.3, 2.4, 2.5}},
		{"test-invalid-value-from-envvar", "FLOAT4", ",", []float64{1.0}, []float64{1.0}},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetFloat64Slice(tt.key, tt.sep, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expectedValue) {
				t.Errorf("GetFloat64Slice(\"%s\", \"%s\", %#v): expected %#v, actual %#v", tt.key, tt.sep, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetBytes(t *testing.T) {
	os.Setenv("STRING2", "STRING2") //nolint:errcheck

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
			if !bytes.Equal(result, tt.expectedValue) {
				t.Errorf("GetBytes(\"%s\", %#v): expected %#v, actual %#v", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetDuration(t *testing.T) {
	os.Setenv("DURATION2", "10") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  int64
		duration      time.Duration
		expectedValue time.Duration
	}{
		{"test-default-value", "DURATION1", 1, time.Second, 1 * time.Second},
		{"test-value-from-envvar", "DURATION2", 1, time.Hour, 10 * time.Hour},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetDuration(tt.key, tt.defaultValue, tt.duration)
			if result != tt.expectedValue {
				t.Errorf("GetDuration(\"%s\", %d, %v): expected %v, actual %v", tt.key, tt.defaultValue, tt.duration, tt.expectedValue, result)
			}
		})
	}
}

func TestGetBase64ToBytes(t *testing.T) {
	os.Setenv("BASE64-2", "SGVsbG8gV29ybGQgMg==") //nolint:errcheck
	os.Setenv("BASE64-3", "invalid-base64-value") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  []byte
		expectedValue []byte
	}{
		{"test-default-value", "BASE64-1", []byte("Hello World 1"), []byte("Hello World 1")},
		{"test-value-from-envvar", "BASE64-2", []byte("Hello World 1"), []byte("Hello World 2")},
		{"test-invalid-value-from-envvar", "BASE64-3", []byte("Hello World 1"), []byte("Hello World 1")},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetBase64ToBytes(tt.key, tt.defaultValue)
			if !bytes.Equal(result, tt.expectedValue) {
				t.Errorf("GetBase64ToBytes(\"%s\", %#v): expected %#v, actual %#v", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}

func TestGetBase64ToString(t *testing.T) {
	os.Setenv("BASE64-2", "SGVsbG8gV29ybGQgMg==") //nolint:errcheck
	os.Setenv("BASE64-3", "invalid-base64-value") //nolint:errcheck

	var tests = []struct {
		kind          string
		key           string
		defaultValue  string
		expectedValue string
	}{
		{"test-default-value", "BASE64-1", "Hello World 1", "Hello World 1"},
		{"test-value-from-envvar", "BASE64-2", "Hello World 1", "Hello World 2"},
		{"test-invalid-value-from-envvar", "BASE64-3", "Hello World 1", "Hello World 1"},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			result := GetBase64ToString(tt.key, tt.defaultValue)
			if result != tt.expectedValue {
				t.Errorf("GetBase64ToString(\"%s\", %#v): expected %#v, actual %#v", tt.key, tt.defaultValue, tt.expectedValue, result)
			}
		})
	}
}
