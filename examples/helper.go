package examples

import "github.com/thuongtruong1009/gouse/helper"

func helperRandomID() {
	randomID := helper.RandomID()
	println("randomID: ", randomID)
}

func helperSelect() {
	options := []string{"a", "b", "c"}
	selected, err := helper.Select("Select an option:", options)
	if err != nil {
		panic(err)
	}
	println("selected: ", selected)
}

func HelperExample() {
	// helperRandomID()
	helperSelect()
}
