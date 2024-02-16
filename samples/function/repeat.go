package function

import "github.com/thuongtruong1009/gouse/function"

func funcTimes() {
	function.Times(func() {
		println("Times")
	}, 3)
}
