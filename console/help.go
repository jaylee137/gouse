package console

import (
	"flag"
	"fmt"
	"strconv"
)

const desc = "Unsupported action type"

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
	var action interface{}

	switch value := f.(type) {
	case func(string):
		action = StringAction{Action: value}
	case func(int):
		action = IntAction{Action: value}
	case func(bool):
		action = BoolAction{Action: value}
	default:
		println(desc)
		return nil
	}

	return action
}

// func Help(name string, options []*HelpOptions) {
// 	for _, option := range options {
// 		switch autoDetectAction(option.Action).(type) {
// 		case StringAction:
// 			flag.String(option.Opt, option.Val.(string), option.Desc)
// 		case IntAction:
// 			flag.Int(option.Opt, option.Val.(int), option.Desc)
// 		case BoolAction:
// 			flag.Bool(option.Opt, option.Val.(bool), option.Desc)
// 		default:
// 			fmt.Println("Unsupported action type")
// 		}
// 	}

// 	flag.Parse()

// 	if flag.NArg() > 0 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
// 		fmt.Printf("Usage: %s [options]\n", name)
// 		fmt.Println("Options:")
// 		for _, option := range options {
// 			fmt.Printf("  -%s\n", option.Opt)
// 			fmt.Printf("        %s\n", option.Desc)
// 		}
// 	}

// 	// Run the associated actions for all set flags
// 	for _, option := range options {
// 		flag.Visit(func(f *flag.Flag) {
// 			if f.Name == option.Opt {
// 				switch a := autoDetectAction(option.Action).(type) {
// 				case StringAction:
// 					a.Action(f.Value.String())
// 				case IntAction:
// 					val, _ := strconv.Atoi(f.Value.String())
// 					a.Action(val)
// 				case BoolAction:
// 					val, _ := strconv.ParseBool(f.Value.String())
// 					a.Action(val)
// 				default:
// 					fmt.Println("Unsupported action type")
// 				}
// 			}
// 		})
// 	}
// }

func Help(name string, options []*HelpOptions) {
	parseFlags(options)

	printUsage(name, options)

	runActions(options)
}

func parseFlags(options []*HelpOptions) {
	for _, option := range options {
		switch autoDetectAction(option.Action).(type) {
		case StringAction:
			flag.String(option.Opt, option.Val.(string), option.Desc)
		case IntAction:
			flag.Int(option.Opt, option.Val.(int), option.Desc)
		case BoolAction:
			flag.Bool(option.Opt, option.Val.(bool), option.Desc)
		default:
			println(desc)
		}
	}

	flag.Parse()
}

func printUsage(name string, options []*HelpOptions) {
	if flag.NArg() > 0 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
		fmt.Printf("Usage: %s [options]\n", name)
		fmt.Println("Options:")
		for _, option := range options {
			fmt.Printf("  -%s\n", option.Opt)
			fmt.Printf("        %s\n", option.Desc)
		}
	}
}

func runActions(options []*HelpOptions) {
	flag.Visit(func(f *flag.Flag) {
		for _, option := range options {
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
					println(desc)
				}
			}
		}
	})
}
