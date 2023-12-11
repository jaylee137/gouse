package sample

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/types"
)

func typeCheck() {
	println("Check type is int: ", types.IsInt(1))
	println("Check type is uint: ", types.IsUnInt(-1))
	println("Check type is float: ", types.IsFloat(1.1))
	println("Check type is complex: ", types.IsComplex(1.1))
	println("Check type is number: ", types.IsNumber(1.23))
	println("Check type is string: ", types.IsString("1"))
	println("Check type is rune: ", types.IsRune('a'))
	println("Check type is byte: ", types.IsByte(byte('a')))
	println("Check type is uintptr: ", types.IsUintptr(uintptr(1)))
	println("Check type is bool: ", types.IsBool(true))
	println("Check type is slice: ", types.IsSlice([]int{1, 2, 3}))
	println("Check type is map: ", types.IsMap(map[string]int{"a": 1}))
	println("Check type is struct: ", types.IsStruct(struct{}{}))
	println("Check type is error: ", types.IsError(new(error)))
	println("Check type is interface: ", types.IsInterface(new(interface{})))
	println("Check type is channel: ", types.IsChannel(make(chan int)))
	println("Check type is array: ", types.IsArray([3]int{1, 2, 3}))
	println("Check type is unsafe pointer: ", types.IsUnsafePointer(new(*int)))
	println("Check type is pointer: ", types.IsPointer(new(int)))
	println("Check type is func: ", types.IsFunc(func() {}))
	println("Check type is nil: ", types.IsNil(nil))
	println("Check type is empty: ", types.IsEmpty(""))
	println("Check type is empty: ", types.IsEmpty(0))
	println("Check type is zero: ", types.IsZero(0))
}

func typeCheckUUID() {
	isValid, err := types.IsUUID("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		println(err.Error())
	}
	println("Check is valid uuid: ", isValid)
}

func typeStructToString() {
	println("Struct to string start:")
	type CompanyInfo struct {
		Company string
		Address string
		Contact string
	}

	companyInfo := CompanyInfo{
		Company: "GeeksforGeeks",
		Address: "Noida",
		Contact: "+91 9876543210",
	}

	println("Struct to string: ", types.StructToString(companyInfo))
	fmt.Println("Struct to map: ", types.StructToMap(companyInfo))
}

func typeStringConvert() {
	println("String to int: ", types.StringToInt("1"))
	fmt.Println("String to float: ", types.StringToFloat("1.1"))
	println("String to bool: ", types.StringToBool("true"))
}

func typeCastToString() {
	println("Cast int to string: ", types.IntToString(1))
	fmt.Println("Cast float to string: ", types.FloatToString(1.1))
	println("Cast bool to string: ", types.BoolToString(true))
	println("Cast interface to string: ", types.InterfaceToString([]int{1, 2, 3}))
}

func TypeExample() {
	typeCheck()
	// typeCheckUUID()

	// typeStructToString()
	// typeStringConvert()
	// typeCastToString()
}
