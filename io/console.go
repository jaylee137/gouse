package io

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	// "reflect"
    // "github.com/jedib0t/go-pretty/v6/table"
)

func ClearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func OutputColor[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](color string, chain T) {
	fmt.Printf("%s%v\n", color, chain)
}