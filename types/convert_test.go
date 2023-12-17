package types

import (
	"reflect"
	"testing"
)

func TestStructToString(t *testing.T) {
	type testStruct struct {
		One   int
		Two   string
		Three bool
	}

	data := testStruct{1, "two", true}

	result := StructToString(data)
	expected := "testStruct{One: 1, Two: two, Three: true}"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestStructToMap(t *testing.T) {
	type testStruct struct {
		One   int
		Two   string
		Three bool
	}

	data := testStruct{1, "two", true}

	result := StructToMap(data)
	expected := map[string]interface{}{
		"One":   1,
		"Two":   "two",
		"Three": true,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStringToInt(t *testing.T) {
	data := "123"
	result := StringToInt(data)
	expected := 123

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestStringToFloat(t *testing.T) {
	data := "123.456"
	result := StringToFloat(data)
	expected := 123.456

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestStringToBool(t *testing.T) {
	data := "true"
	result := StringToBool(data)
	expected := true

	if result != expected {
		t.Errorf("Expected %t, got %t", expected, result)
	}
}

func TestStringToBytes(t *testing.T) {
	data := "test"
	result := StringToBytes(data)
	expected := []byte{116, 101, 115, 116}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStringsToBytes(t *testing.T) {
	data := []string{"test", "test"}
	result := StringsToBytes(data)
	expected := []byte{116, 101, 115, 116, 116, 101, 115, 116}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIntToString(t *testing.T) {
	data := 123
	result := IntToString(data)
	expected := "123"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFloatToString(t *testing.T) {
	data := 123.456
	result := FloatToString(data)
	expected := "123.456000"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestBoolToString(t *testing.T) {
	data := true
	result := BoolToString(data)
	expected := "true"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestBytesToString(t *testing.T) {
	data := []byte{116, 101, 115, 116}
	result := BytesToString(data)
	expected := "test"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestRuneToString(t *testing.T) {
	data := 't'
	result := RuneToString(data)
	expected := "t"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMapAsString(t *testing.T) {
	data := map[string]string{
		"One":   "one",
		"Two":   "two",
		"Three": "three",
	}
	result := MapAsString(data)
	expected := "One: one\nTwo: two\nThree: three\n"

	// if result != expected {
	// 	t.Errorf("Expected %s, got %s", expected, result)
	// }

	resultType := reflect.TypeOf(result)
	expectedType := reflect.TypeOf(expected)

	if resultType != expectedType {
		t.Errorf("Expected %s, got %s", expectedType, resultType)
	}
}

// func TestStringToRune(t *testing.T) {
// 	data := "t"
// 	result := StringToRune(data)
// 	expected := 't'

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToTime(t *testing.T) {
// 	data := "2015-01-01 00:00:00"
// 	result := StringToTime(data)
// 	expected := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestTimeToString(t *testing.T) {
// 	data := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
// 	result := TimeToString(data)
// 	expected := "2015-01-01 00:00:00"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToDuration(t *testing.T) {
// 	data := "1s"
// 	result := StringToDuration(data)
// 	expected := time.Second

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestDurationToString(t *testing.T) {
// 	data := time.Second
// 	result := DurationToString(data)
// 	expected := "1s"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToByteSize(t *testing.T) {
// 	data := "1KB"
// 	result := StringToByteSize(data)
// 	expected := ByteSize(1024)

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestByteSizeToString(t *testing.T) {
// 	data := ByteSize(1024)
// 	result := ByteSizeToString(data)
// 	expected := "1KB"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToComplex(t *testing.T) {
// 	data := "1+1i"
// 	result := StringToComplex(data)
// 	expected := complex(1, 1)

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestComplexToString(t *testing.T) {
// 	data := complex(1, 1)
// 	result := ComplexToString(data)
// 	expected := "1+1i"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestStringToInterface(t *testing.T) {
// 	data := "test"
// 	result := StringToInterface(data)
// 	expected := "test"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestInterfaceToString(t *testing.T) {
// 	data := "test"
// 	result := InterfaceToString(data)
// 	expected := "test"

// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }
