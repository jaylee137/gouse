package function

import "github.com/thuongtruong1009/gouse/function"

func funcInterval() {
	function.Interval(func() {
		println("Interval")
	}, 1)
}
