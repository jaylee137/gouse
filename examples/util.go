package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/utils"
)

func utilDelay() {
	fmt.Println("Delay start:")

	result := utils.DelayF(func() string {
		return "Delayed return after 2s"
	}, 2)

	if result.HasReturn{
		fmt.Println(result.Value)
	} else {
		fmt.Println("No result")
	}

	utils.DelayV(func() {
		fmt.Println("Delayed not return")
	}, 3)
}

func UtilExample() {
	utilDelay()
}