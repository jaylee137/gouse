package examples

import "github.com/thuongtruong1009/gouse/helper"

func helperRandomID() {
	randomID := helper.RandomID()
	println("randomID: ", randomID)
}

func HelperExample() {
	helperRandomID()
}
