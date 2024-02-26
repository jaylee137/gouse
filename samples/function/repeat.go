package function

import "github.com/thuongtruong1009/gouse/function"

func SampleFuncTimes() {
	function.Times(func() {
		println("Times")
	}, 3)
}
