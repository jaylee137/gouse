package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/types"
)

func checkType() {
	fmt.Println("Check type start:")
	fmt.Println(types.IsInt(1))
	fmt.Println(types.IsFloat(1.0))
	fmt.Println(types.IsString("1"))
	fmt.Println(types.IsBool(true))
	fmt.Println(types.IsSlice([]int{1, 2, 3}))
	fmt.Println(types.IsMap(map[string]int{"a": 1}))
	fmt.Println(types.IsStruct(struct{}{}))
	fmt.Println(types.IsArray([3]int{1, 2, 3}))
	fmt.Println(types.IsPointer(&struct{}{}))
	fmt.Println(types.IsFunc(func() {}))
	fmt.Println(types.IsNil(nil))
	fmt.Println(types.IsEmpty(""))
	fmt.Println(types.IsUndefined(1))
}

func TypeExample() {
	checkType()
}