package check

import "github.com/thuongtruong1009/gouse/types"

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
