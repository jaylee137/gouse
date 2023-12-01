package examples

import (
	"github.com/thuongtruong1009/gouse/constants"
	"github.com/thuongtruong1009/gouse/helper"
)

func helperClearConsole() {
	helper.ClearConsole()
	println("console cleared")
}

func helperOutputColor() {
	helper.OutputColor(constants.DEFAULT, "this is default")
	helper.OutputColor(constants.WHITE, "this is white")
	helper.OutputColor(constants.RED, "this is red")
	helper.OutputColor(constants.GREEN, "this is green")
	helper.OutputColor(constants.YELLOW, "this is yellow")
	helper.OutputColor(constants.BLUE, "this is blue")
	helper.OutputColor(constants.MAGENTA, "this is magenta")
	helper.OutputColor(constants.CYAN, "this is cyan")
}

func helperRandomID() {
	randomID := helper.RandomID()
	println("randomID: ", randomID)
}

func HelperExample() {
	helperClearConsole()
	helperOutputColor()
	helperRandomID()
}