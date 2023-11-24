package examples

import "github.com/thuongtruong1009/gouse/utils"

func utilDelay() {
	println("Delay start:")

	result := utils.DelayF(func() string {
		return "Delayed return after 2s"
	}, 2)

	if result.HasReturn{
		println(result.Value)
	} else {
		println("No result")
	}

	utils.Delay(func() {
		println("Delayed not return")
	}, 3)
}

func UtilExample() {
	utilDelay()
}