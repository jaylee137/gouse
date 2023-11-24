package examples

import "github.com/thuongtruong1009/gouse/types"

func checkType() {
	println("Check type start:")
	println(types.IsInt(1))
	println(types.IsFloat(1.0))
	println(types.IsString("1"))
	println(types.IsBool(true))
	println(types.IsSlice([]int{1, 2, 3}))
	println(types.IsMap(map[string]int{"a": 1}))
	println(types.IsStruct(struct{}{}))
	println(types.IsArray([3]int{1, 2, 3}))
	println(types.IsPointer(&struct{}{}))
	println(types.IsFunc(func() {}))
	println(types.IsNil(nil))
	println(types.IsEmpty(""))
	println(types.IsUndefined(1))
}

func TypeExample() {
	checkType()
}