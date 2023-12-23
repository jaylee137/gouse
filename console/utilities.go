package console

import (
	"flag"
	"fmt"

	// "fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

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

type HelpOptions struct {
	Opt    string
	Desc   string
	Val    interface{}
	Action interface{}
}

type StringAction struct {
	Action func(string)
}

type IntAction struct {
	Action func(int)
}

type BoolAction struct {
	Action func(bool)
}

func autoDetectAction(f interface{}) interface{} {
	switch f.(type) {
	case func(string):
		return StringAction{Action: f.(func(string))}
	case func(int):
		return IntAction{Action: f.(func(int))}
	case func(bool):
		return BoolAction{Action: f.(func(bool))}
	default:
		fmt.Println("Unsupported action type")
		os.Exit(1)
		return nil
	}
}

func Help(name string, options []*HelpOptions) {
	for _, option := range options {
		switch autoDetectAction(option.Action).(type) {
		case StringAction:
			flag.String(option.Opt, option.Val.(string), option.Desc)
		case IntAction:
			flag.Int(option.Opt, option.Val.(int), option.Desc)
		case BoolAction:
			flag.Bool(option.Opt, option.Val.(bool), option.Desc)
		default:
			fmt.Println("Unsupported action type")
			os.Exit(1)
		}
	}

	flag.Parse()

	if flag.NArg() > 0 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
		fmt.Printf("Usage: %s [options]\n", name)
		fmt.Println("Options:")
		for _, option := range options {
			fmt.Printf("  -%s\n", option.Opt)
			fmt.Printf("        %s\n", option.Desc)
		}

		os.Exit(0)
	}

	// Run the associated actions for all set flags
	for _, option := range options {
		flag.Visit(func(f *flag.Flag) {
			if f.Name == option.Opt {
				switch a := autoDetectAction(option.Action).(type) {
				case StringAction:
					a.Action(f.Value.String())
				case IntAction:
					val, _ := strconv.Atoi(f.Value.String())
					a.Action(val)
				case BoolAction:
					val, _ := strconv.ParseBool(f.Value.String())
					a.Action(val)
				default:
					fmt.Println("Unsupported action type")
					os.Exit(1)
				}
			}
		})
	}
}
