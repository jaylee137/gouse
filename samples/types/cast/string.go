package cast

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/types"
)

func typeStringConvert() {
	println("Convert string to int: ", types.StringToInt("1"))
	fmt.Println("Convert string to float: ", types.StringToFloat("1.1"))
	println("Convert string to bool: ", types.StringToBool("true"))
	fmt.Println("Convert string to bytes: ", string(types.StringToBytes("1")), "->", types.StringToBytes("1"))
	fmt.Println("Convert strings to bytes: ", string(types.StringsToBytes([]string{"1", "2", "3"})), "->", types.StringsToBytes([]string{"1", "2", "3"}))
}

func typeCastToString() {
	println("Cast int to string: ", types.IntToString(1))
	fmt.Println("Cast float to string: ", types.FloatToString(1.1))
	println("Cast bool to string: ", types.BoolToString(true))
	println("Cast interface to string: ", types.InterfaceToString([]int{1, 2, 3}))
	println("Cast bytes to string: ", types.BytesToString([]byte{49, 50, 51}))
	println("Cast rune to string: ", types.RuneToString('a'))
}
