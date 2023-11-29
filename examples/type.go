package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/types"
)

func typeCheck() {
	println("Check is int: ", types.IsInt(1))
	println("Check is float: ", types.IsFloat(1.1))
	println("Check is string: ", types.IsString("1"))
	println("Check is bool: ", types.IsBool(true))
	println("Check is slice: ", types.IsSlice([]int{1, 2, 3}))
	println("Check is map: ", types.IsMap(map[string]int{"a": 1}))
	println("Check is struct: ", types.IsStruct(struct{}{}))
	println("Check is array: ", types.IsArray([3]int{1, 2, 3}))
	println("Check is pointer: ", types.IsPointer(new(int)))
	println("Check is func: ", types.IsFunc(func() {}))
	println("Check is nil: ", types.IsNil(nil))
	println("Check is empty: ", types.IsEmpty(""))
	println("Check is empty: ", types.IsEmpty(0))
	println("Check is zero: ", types.IsZero(0))
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
}

func TypeExample() {
	typeCheck()

	typeStructToString()
	typeStringConvert()
	typeCastToString()
}
