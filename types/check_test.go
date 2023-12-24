package types

import (
	"errors"
	"testing"
)

func TestIsInt(t *testing.T) {
	var i int
	if !IsInt(i) {
		t.Error("IsInt(i) should return true")
	}
}

func TestIsUnInt(t *testing.T) {
	var i uint
	if !IsUnInt(i) {
		t.Error("IsUnInt(i) should return true")
	}
}

func TestIsFloat(t *testing.T) {
	var i float32
	if !IsFloat(i) {
		t.Error("IsFloat(i) should return true")
	}
}

func TestIsComplex(t *testing.T) {
	var i complex64
	if !IsComplex(i) {
		t.Error("IsComplex(i) should return true")
	}
}

func TestIsNumber(t *testing.T) {
	var i int
	if !IsNumber(i) {
		t.Error("IsNumber(i) should return true")
	}
}

func TestIsString(t *testing.T) {
	var i string
	if !IsString(i) {
		t.Error("IsString(i) should return true")
	}
}

func TestIsRune(t *testing.T) {
	var i rune
	if !IsRune(i) {
		t.Error("IsRune(i) should return true")
	}
}

func TestIsByte(t *testing.T) {
	var i byte
	if !IsByte(i) {
		t.Error("IsByte(i) should return true")
	}
}

func TestIsUintptr(t *testing.T) {
	var i uintptr
	if !IsUintptr(i) {
		t.Error("IsUintptr(i) should return true")
	}
}

func TestIsError(t *testing.T) {
	var i error = errors.New("test")
	if !IsError(i) {
		t.Error("IsError(i) should return true")
	}
}

func TestIsChannel(t *testing.T) {
	var i chan int
	if !IsChannel(i) {
		t.Error("IsChannel(i) should return true")
	}
}

func TestIsUnsafePointer(t *testing.T) {
	var i uintptr
	if !IsUnsafePointer(i) {
		t.Error("IsUnsafePointer(i) should return true")
	}
}

func TestIsPointer(t *testing.T) {
	var i *int
	if !IsPointer(i) {
		t.Error("IsPointer(i) should return true")
	}
}

func TestIsBool(t *testing.T) {
	var i bool
	if !IsBool(i) {
		t.Error("IsBool(i) should return true")
	}
}

func TestIsSlice(t *testing.T) {
	var i []int
	if !IsSlice(i) {
		t.Error("IsSlice(i) should return true")
	}
}

func TestIsMap(t *testing.T) {
	var i map[string]int
	if !IsMap(i) {
		t.Error("IsMap(i) should return true")
	}
}

func TestIsStruct(t *testing.T) {
	type testStruct struct {
		a int
	}
	var i = testStruct{
		a: 1,
	}
	if !IsStruct(i) {
		t.Error("IsStruct(i) should return true")
	}
}

func TestIsArray(t *testing.T) {
	var i [3]int
	if !IsArray(i) {
		t.Error("IsArray(i) should return true")
	}
}

func TestIsFunc(t *testing.T) {
	var i func()
	if !IsFunc(i) {
		t.Error("IsFunc(i) should return true")
	}
}

func TestIsNil(t *testing.T) {
	var i interface{}
	if !IsNil(i) {
		t.Error("IsNil(i) should return true")
	}
}

func TestIsEmpty(t *testing.T) {
	var i string
	if !IsEmpty(i) {
		t.Error("IsEmpty(i) should return true")
	}
}

func TestIsZero(t *testing.T) {
	var i int
	if !IsZero(i) {
		t.Error("IsZero(i) should return true")
	}
}

func TestUndefined(t *testing.T) {
	var i bool
	if !IsUndefined(i) {
		t.Error("IsUndefined(i) should return true")
	}
}
