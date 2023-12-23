package console

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/thuongtruong1009/gouse/strings"
)

func Cmd(defaultCmmand string, windowsCmmand ...string) {
	var cmd *exec.Cmd

	splitCmd := strings.Split(defaultCmmand, " ")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c")
		if len(windowsCmmand) > 0 {
			cmd.Args = append(cmd.Args, windowsCmmand...)
		} else {
			cmd.Args = append(cmd.Args, splitCmd...)
		}
	} else {
		cmd = exec.Command("sh", "-c")
		cmd.Args = append(cmd.Args, splitCmd...)
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Clear() {
	Cmd("clear", "cls")
}

func WithColor[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](color string, chain T) {
	fmt.Printf("%s%v\n", color, chain)
}
