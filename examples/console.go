package examples

import (
	"github.com/thuongtruong1009/gouse/console"
	"github.com/thuongtruong1009/gouse/console/choice"
	"github.com/thuongtruong1009/gouse/console/countdown"
	"github.com/thuongtruong1009/gouse/console/dir"
	"github.com/thuongtruong1009/gouse/console/glamour"
	"github.com/thuongtruong1009/gouse/console/inline"
	"github.com/thuongtruong1009/gouse/console/list"
	"github.com/thuongtruong1009/gouse/console/paper"
	"github.com/thuongtruong1009/gouse/console/parallel"
	"github.com/thuongtruong1009/gouse/console/progress"
	"github.com/thuongtruong1009/gouse/console/realtime"
	"github.com/thuongtruong1009/gouse/console/sequence"
	"github.com/thuongtruong1009/gouse/console/spinner"
	"github.com/thuongtruong1009/gouse/console/split"
	"github.com/thuongtruong1009/gouse/console/stopwatch"
	"github.com/thuongtruong1009/gouse/console/tab"
	"github.com/thuongtruong1009/gouse/console/table"
	"github.com/thuongtruong1009/gouse/shared"
)

func consoleCmd() {
	console.Cmd("echo command is working")

	// first param is default command, second param is windows command (default is empty)
	console.Cmd("ls", "clear")
}

func consoleClear() {
	println("console will be cleared now")
	console.Clear()
	println("console cleared")
}

func consoleWithColor() {
	console.WithColor(shared.DEFAULT_FG, "this is default")
	console.WithColor(shared.WHITE_FG, "this is white")
	console.WithColor(shared.RED_FG, "this is red")
	console.WithColor(shared.GREEN_FG, "this is green")
	console.WithColor(shared.YELLOW_FG, "this is yellow")
	console.WithColor(shared.BLUE_FG, "this is blue")
	console.WithColor(shared.MAGENTA_FG, "this is magenta")
	console.WithColor(shared.CYAN_FG, "this is cyan")
}

func consoleHelp() {
	console.Help()
}

func consoleSelect() {
	optconsolens := []string{"a", "b", "c"}
	selected, err := console.Select("Select an optconsolen:", optconsolens)
	if err != nil {
		panic(err)
	}
	console.Confirm("Are you sure?")
	println("You selected: ", selected)
}

func consoleList() {
	title := "My Fave Things"
	items := []list.Item{
		list.Item{Label: "Raspberry Pi’s", Desc: "I have ’em all over my house"},
		list.Item{Label: "Nutella", Desc: "It's good on toast"},
		list.Item{Label: "Bitter melon", Desc: "It cools you down"},
		list.Item{Label: "Nice socks", Desc: "And by that I mean socks without holes"},
		list.Item{Label: "Eight hours of sleep", Desc: "I had this once"},
		list.Item{Label: "Cats", Desc: "Usually"},
		list.Item{Label: "Plantasia, the album", Desc: "My plants love it too"},
		list.Item{Label: "Pour over coffee", Desc: "It takes forever to make though"},
		list.Item{Label: "VR", Desc: "Virtual reality...what is there to say?"},
		list.Item{Label: "Noguchi Lamps", Desc: "Such pleasing organic forms"},
		list.Item{Label: "Linux", Desc: "Pretty much the best OS"},
		list.Item{Label: "Business school", Desc: "Just kidding"},
		list.Item{Label: "Pottery", Desc: "Wet clay is a great feeling"},
		list.Item{Label: "Shampoo", Desc: "Nothing like clean hair"},
		list.Item{Label: "Table tennis", Desc: "It’s surprisingly exhausting"},
		list.Item{Label: "Milk crates", Desc: "Great for packing in your extra stuff"},
		list.Item{Label: "Afternoon tea", Desc: "Especially the tea sandwich part"},
		list.Item{Label: "Stickers", Desc: "The thicker the vinyl the better"},
		list.Item{Label: "20° Weather", Desc: "Celsius, not Fahrenheit"},
		list.Item{Label: "Warm light", Desc: "Like around 2700 Kelvin"},
		list.Item{Label: "The vernal equinox", Desc: "The autumnal equinox is pretty good too"},
		list.Item{Label: "Gaffer’s tape", Desc: "Basically sticky fabric"},
		list.Item{Label: "Terrycloth", Desc: "In other words, towel fabric"},
	}

	list.Default(title, items)
}

func consolePaper() {
	paper.Run("main.go")
}

func consoleProgress() {
	progress.Run()
}

func consoleRealtime() {
	realtime.Run()
}

func consoleChoice() {
	choice.Run()
}

func consoleSpinner() {
	spinner.Run()
}

func consoleSplit() {
	split.Run()
}

func consoleStopwatch() {
	stopwatch.Run()
}

func consoleTable() {
	table.Run()
}

func consoleTab() {
	tab.Run()
}

func consoleCountdown() {
	countdown.Run()
}

func consoleSequence() {
	sequence.Run()
}

func consoleInline() {
	inline.Run()
}

func consoleParallel() {
	parallel.Run()
}

func consoleDir() {
	dir.Run()
}

func consoleGlamour() {
	glamour.Run()
}

func ConsoleExample() {
	// consoleCmd()
	// consoleClear()
	// consoleWithColor()
	// consoleSelect()
	consoleHelp()

	// consoleList()
	// consolePaper()
	// consoleProgress()
	// consoleRealtime()
	// consoleChoice()
	// consoleSpinner()
	// consoleSplit()
	// consoleStopwatch()
	// consoleTable()
	// consoleTab()
	// consoleCountdown()
	// consoleSequence()
	// consoleInline()
	// consoleParallel()
	// consoleDir()
	// consoleGlamour()
}
