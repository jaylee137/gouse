# Struct

## Imports

```go
import (
	"fmt"

	"github.com/thuongtruong1009/gouse/array"
)
```
## Functions


### arrEqual

```go
func arrEqual() {
	println("--- Compare equal ---")
	println("[int]: ", array.Equal(1, 1))
	println("[uint]: ", array.Equal(uint(1), uint(1)))
	println("[float]: ", array.Equal(1.2, 1.1))
	println("[string]: ", array.Equal("1", "0"))
	println("[rune]: ", array.Equal('a', 'a'))
	println("[bool]: ", array.Equal(true, true))
	println("[complex]: ", array.Equal(1+2i, 1+2i))
	println("[struct]: ", array.Equal(struct{ a int }{1}, struct{ a int }{1}))
}
```

### arrIncludes

```go
func arrIncludes() {
	println("--- Check element is exist in array ---")
	println("[int]: ", array.Includes([]int{1, -2, 3}, 1))
	println("[uint]: ", array.Includes([]uint{1, 2, 3}, 1))
	println("[float]: ", array.Includes([]float64{1.2, 2.3, 3.4}, 1.2))
	println("[string]: ", array.Includes([]string{"1", "2", "3"}, "0"))
	println("[rune]: ", array.Includes([]rune{'a', 'b', 'c'}, 'a'))
	println("[complex]: ", array.Includes([]complex128{1 + 2i, 2 + 3i}, 1+2i))
	println("[struct]: ", array.Includes([]struct{ a int }{{1}, {2}}, struct{ a int }{3}))
}
```

### arrMost

```go
func arrMost() {
	println("--- Most frequency in array ---")
	println("[int]: ", array.Most([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	println("[uint]: ", array.Most([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	fmt.Println("[float]: ", array.Most([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	println("[string]: ", array.Most([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	println("[rune]: ", string(array.Most([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	fmt.Println("[complex]: ", array.Most([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	fmt.Println("[struct]: ", array.Most([]struct{ a int }{{1}, {2}, {3}, {2}, {2}, {1}, {2}, {3}}))
}
```

### arrSum

```go
func arrSum() {
	println("--- Sum elements in array ---")
	println("[int]: ", array.Sum([]int{1, -2, 3}))
	println("[uint]: ", array.Sum([]uint{1, 2, 3}))
	fmt.Println("[float]: ", array.Sum([]float64{1.2, 2.3, 3.4}))
	println("[rune]: ", array.Sum([]rune{'a', 'b', 'c'}))
	fmt.Println("[complex]: ", array.Sum([]complex128{1 + 2i, 2 + 3i, 3 + 4i}))
}
```

### arrChunk

```go
func arrChunk() {
	println("--- Chunk array ---")
	fmt.Println("[int]: ", array.Chunk([]int{1, -2, 3, -4, 5, 6}, 3))
	fmt.Println("[uint]: ", array.Chunk([]uint{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println("[float]: ", array.Chunk([]float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7}, 3))
	fmt.Println("[string]: ", array.Chunk([]string{"1", "2", "3", "4", "5", "6"}, 3))
	fmt.Println("[rune]: ", array.Chunk([]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3))
	fmt.Println("[complex]: ", array.Chunk([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 4 + 5i, 5 + 6i, 6 + 7i}, 3))
	fmt.Println("[struct]: ", array.Chunk([]struct{ a int }{{1}, {2}, {3}, {4}, {5}, {6}}, 3))
}
```

### arrDiff

```go
func arrDiff() {
	println("--- Difference array ---")
	fmt.Println("[int]: ", array.Diff([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println("[uint]: ", array.Diff([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	fmt.Println("[float]: ", array.Diff([]float64{1.2, 2.3, 3.4}, []float64{4.5, 5.6, 6.7}))
	fmt.Println("[string]: ", array.Diff([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	fmt.Println("[rune]: ", array.Diff([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	fmt.Println("[complex]: ", array.Diff([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	fmt.Println("[struct]: ", array.Diff([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
```

### arrDrop

```go
func arrDrop() {
	println("--- Drop elements in array (default n = 1) ---")
	fmt.Println("[int] with default: ", array.Drop([]int{1, -2, 3, -4, 5, 6}))
	fmt.Println("[int]: ", array.Drop([]int{1, -2, 3, -4, 5, 6}, 2))
	fmt.Println("[uint]: ", array.Drop([]uint{1, 2, 3, 4, 5, 7}, 2))
	fmt.Println("[float]: ", array.Drop([]float64{1.2, 2.3, 3.4}, 2))
	fmt.Println("[string]: ", array.Drop([]string{"1", "4", "5", "6"}, 2))
	fmt.Println("[rune]: ", array.Drop([]rune{'a', 'b', 'd', 'e', 'f'}, 2))
	fmt.Println("[complex]: ", array.Drop([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 2))
	fmt.Println("[struct]: ", array.Drop([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, 2))
}
```

### arrIndex

```go
func arrIndex() {
	println("--- Index of element in array ---")
	println("[int]: ", array.Index([]int{1, -2, 3, -4, 5, 6}, 3))
	println("[uint]: ", array.Index([]uint{1, 2, 3, 4, 5, 7}, 3))
	println("[float]: ", array.Index([]float64{1.2, 2.3, 3.4}, 3.4))
	println("[string]: ", array.Index([]string{"1", "4", "5", "6"}, "5"))
	println("[rune]: ", array.Index([]rune{'a', 'b', 'd', 'e', 'f'}, 'e'))
	println("[complex]: ", array.Index([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 5+6i))
	println("[struct]: ", array.Index([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, struct{ a int }{3}))
}
```

### arrMerge

```go
func arrMerge() {
	println("--- Merge arrays ---")
	println("[int]: ", array.Merge([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, []int{1, -2, 3, -4, 5, 6}))
	println("[uint]: ", array.Merge([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", array.Merge([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", array.Merge([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", array.Merge([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", array.Merge([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", array.Merge([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
```

### arrMin

```go
func arrMin() {
	println("--- Min element in array ---")
	println("[int]: ", array.Min([]int{1, -2, 3}))
	println("[uint]: ", array.Min([]uint{1, 2, 3}))
	println("[string]: ", array.Min([]string{"z", "d", "m"}))
	println("[rune]: ", string(array.Min([]rune{'z', 'd', 'm'})))
	println("[float]: ", array.Min([]float64{1.2, 2.3, 3.4}))
}
```

### arrMax

```go
func arrMax() {
	println("--- Max element in array ---")
	println("[int]: ", array.Max([]int{1, -2, 3}))
	println("[uint]: ", array.Max([]uint{1, 2, 3}))
	println("[string]: ", array.Max([]string{"z", "d", "m"}))
	println("[rune]: ", string(array.Max([]rune{'z', 'd', 'm'})))
	println("[float]: ", array.Max([]float64{1.2, 2.3, 3.4}))
}
```

### arrIntersect

```go
func arrIntersect() {
	println("--- Intersection arrays ---")
	println("[int]: ", array.Intersect([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	println("[uint]: ", array.Intersect([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	println("[float]: ", array.Intersect([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	println("[string]: ", array.Intersect([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	println("[rune]: ", array.Intersect([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	println("[complex]: ", array.Intersect([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	println("[struct]: ", array.Intersect([]struct{ a int }{{1}, {-2}, {3}, {4}, {5}, {6}}, []struct{ a int }{{1}, {2}}))
}
```

### arrKeyBy

```go
func arrKeyBy() {
	println("--- Find key of element pass condition in callback function ---")
	println("[int]: ", array.KeyBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", array.KeyBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", array.KeyBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", array.KeyBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", array.KeyBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", array.KeyBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	println("[struct]: ", array.KeyBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}
```

### arrIndexBy

```go
func arrIndexBy() {
	println("--- Find index of element pass condition in callback function ---")
	println("[int]: ", array.IndexBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", array.IndexBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", array.IndexBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", array.IndexBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", array.IndexBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", array.IndexBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	println("[struct]: ", array.IndexBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}
```

### arrFilterBy

```go
func arrFilterBy() {
	println("--- Filter elements in array by pass condition in callback function---")
	println("[int]: ", array.FilterBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v > 2
	}))

	println("[uint]: ", array.FilterBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v > 2
	}))

	println("[float]: ", array.FilterBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v > 2
	}))

	println("[string]: ", array.FilterBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v > "3"
	}))

	println("[rune]: ", array.FilterBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v > 'c'
	}))

	println("[complex]: ", array.FilterBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return real(v) > 3
	}))

	println("[struct]: ", array.FilterBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a > 0
	}))
}
```

### arrRejectBy

```go
func arrRejectBy() {
	println("--- Reject elements in array by pass condition in callback function---")
	println("[int]: ", array.RejectBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v > 2
	}))

	println("[uint]: ", array.RejectBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v > 2
	}))

	println("[float]: ", array.RejectBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v > 2
	}))

	println("[string]: ", array.RejectBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v > "3"
	}))

	println("[rune]: ", array.RejectBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v > 'c'
	}))

	println("[complex]: ", array.RejectBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return real(v) > 3
	}))
}
```

### arrFindBy

```go
func arrFindBy() {
	println("--- Find element in array by pass condition in callback function---")
	println("[int]: ", array.FindBy([]int{1, -2, 3, -4, 5, 6}, func(v int) bool {
		return v == 3
	}))

	println("[uint]: ", array.FindBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) bool {
		return v == 3
	}))

	println("[float]: ", array.FindBy([]float64{1.2, 2.3, 3.4}, func(v float64) bool {
		return v == 3.4
	}))

	println("[string]: ", array.FindBy([]string{"1", "4", "5", "6"}, func(v string) bool {
		return v == "5"
	}))

	println("[rune]: ", array.FindBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) bool {
		return v == 'e'
	}))

	println("[complex]: ", array.FindBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) bool {
		return v == 5+6i
	}))

	fmt.Println("[struct]: ", array.FindBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) bool {
		return v.a == 3
	}))
}
```

### arrForBy

```go
func arrForBy() {
	println("--- Loop array then handler with callback function ---")
	print("[int]: ")
	array.ForBy([]int{1, -2, 3, -4, 5, 6}, func(v int) {
		println(v)
	})

	print("[uint]: ")
	array.ForBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) {
		println(v)
	})

	print("[float]: ")
	array.ForBy([]float64{1.2, 2.3, 3.4}, func(v float64) {
		println(v)
	})

	print("[string]: ")
	array.ForBy([]string{"1", "4", "5", "6"}, func(v string) {
		println(v)
	})

	print("[rune]: ")
	array.ForBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) {
		println(v)
	})

	print("[complex]: ")
	array.ForBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) {
		println(v)
	})

	print("[struct]: ")
	array.ForBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) {
		fmt.Println(v)
	})
}
```

### arrMapBy

```go
func arrMapBy() {
	println("--- Map array then handler with callback function ---")
	fmt.Println("[int]: ", array.MapBy([]int{1, -2, 3, -4, 5, 6}, func(v int) int {
		return v * 2
	}))

	fmt.Println("[uint]: ", array.MapBy([]uint{1, 2, 3, 4, 5, 7}, func(v uint) uint {
		return v * 2
	}))

	fmt.Println("[float]: ", array.MapBy([]float64{1.2, 2.3, 3.4}, func(v float64) float64 {
		return v * 2
	}))

	fmt.Println("[string]: ", array.MapBy([]string{"1", "4", "5", "6"}, func(v string) string {
		return v + "1"
	}))

	fmt.Println("[rune]: ", array.MapBy([]rune{'a', 'b', 'd', 'e', 'f'}, func(v rune) rune {
		return v + 1
	}))

	fmt.Println("[complex]: ", array.MapBy([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func(v complex128) complex128 {
		return v * 2
	}))

	fmt.Println("[struct]: ", array.MapBy([]struct{ a int }{{-1}, {-2}, {3}, {4}, {5}, {6}}, func(v struct{ a int }) struct{ a int } {
		return struct{ a int }{v.a * 2}
	}))
}
```

### arrCompact

```go
func arrCompact() {
	result := array.Compact([]interface{}{1, -2, 3, -4, 5, 6, 0, 0.0, "", false, nil})
	fmt.Println("Compact remove all falsy values: ", result)
}
```

### ArrayExample

```go
func ArrayExample() {
	arrEqual()
	arrIncludes()
	arrMost()
	arrSum()
	arrChunk()
	arrDiff()
	arrDrop()
	arrIndex()
	arrMerge()
	arrCompact()

	arrIntersect()
	arrMin()
	arrMax()

	arrKeyBy()
	arrIndexBy()
	arrFilterBy()
	arrRejectBy()
	arrFindBy()
	arrForBy()
	arrMapBy()
}
```

### consoleCmd

```go
func consoleCmd() {
	console.Cmd("echo command is working")

	// first param is default command, second param is windows command (default is empty)
	console.Cmd("ls", "clear")
}
```

### consoleClear

```go
func consoleClear() {
	println("console will be cleared now")
	console.Clear()
	println("console cleared")
}
```

### consoleWithColor

```go
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
```

### consoleHelp

```go
func consoleHelp() {
	name := "myprogram"
	options := []*console.HelpOptions{
		{
			Opt:  "name",
			Desc: "Enter your name",
			Val:  "Example name",
			Action: func(name string) {
				println("Your name is: ", name)
			},
		},
		{
			Opt:  "age",
			Desc: "Enter your age",
			Val:  18,
			Action: func(age int) {
				println("this is age: ", age)
			},
		},
		{
			Opt:  "learning",
			Desc: "Enter your confirm",
			Val:  true,
			Action: func(confirm bool) {
				println("this is your confirm", confirm)
			},
		},
	}
	console.Help(name, options)
	// for _, option := range options {
	// 	fmt.Printf("Option %s: %v\n", option.Opt, option.Val)
	// }
}
```

### consoleSelect

```go
func consoleSelect() {
	optconsolens := []string{"a", "b", "c"}
	selected, err := console.Select("Select an optconsolen:", optconsolens)
	if err != nil {
		panic(err)
	}
	console.Confirm("Are you sure?")
	println("You selected: ", selected)
}
```

### consoleList

```go
func consoleList() {
	title := "My Fave Things"
	items := []list.Item{
		{Label: "Raspberry Pi’s", Desc: "I have ’em all over my house"},
		{Label: "Nutella", Desc: "It's good on toast"},
		{Label: "Bitter melon", Desc: "It cools you down"},
		{Label: "Nice socks", Desc: "And by that I mean socks without holes"},
		{Label: "Eight hours of sleep", Desc: "I had this once"},
		{Label: "Cats", Desc: "Usually"},
		{Label: "Plantasia, the album", Desc: "My plants love it too"},
		{Label: "Pour over coffee", Desc: "It takes forever to make though"},
		{Label: "VR", Desc: "Virtual reality...what is there to say?"},
		{Label: "Noguchi Lamps", Desc: "Such pleasing organic forms"},
		{Label: "Linux", Desc: "Pretty much the best OS"},
		{Label: "Business school", Desc: "Just kidding"},
		{Label: "Pottery", Desc: "Wet clay is a great feeling"},
		{Label: "Shampoo", Desc: "Nothing like clean hair"},
		{Label: "Table tennis", Desc: "It’s surprisingly exhausting"},
		{Label: "Milk crates", Desc: "Great for packing in your extra stuff"},
		{Label: "Afternoon tea", Desc: "Especially the tea sandwich part"},
		{Label: "Stickers", Desc: "The thicker the vinyl the better"},
		{Label: "20° Weather", Desc: "Celsius, not Fahrenheit"},
		{Label: "Warm light", Desc: "Like around 2700 Kelvin"},
		{Label: "The vernal equinox", Desc: "The autumnal equinox is pretty good too"},
		{Label: "Gaffer’s tape", Desc: "Basically sticky fabric"},
		{Label: "Terrycloth", Desc: "In other words, towel fabric"},
	}

	list.Default(title, items)
}
```

### consolePaper

```go
func consolePaper() {
	paper.Run("main.go")
}
```

### consoleProgress

```go
func consoleProgress() {
	// first param is fail message
	// second param is done message
	// third param is increment percent (default 0.25)
	progress.Run("^_^ Oh no, something went wrong", "✔️ Done!", 0.5)
}
```

### consoleRealtime

```go
func consoleRealtime() {
	realtime.Run()
}
```

### consoleChoice

```go
func consoleChoice() {
	question := "What's your favorite flavor?"
	options := []string{"Taro", "Coffee", "Lychee"}

	update := &choice.Model{}
	choice.Run(question, options, update)

	fmt.Printf("\n---\nYou chose %s!\n", update.Choice)
}
```

### consoleSpinner

```go
func consoleSpinner() {
	spinner.Run()
}
```

### consoleSplit

```go
func consoleSplit() {
	split.Run()
}
```

### consoleStopwatch

```go
func consoleStopwatch() {
	stopwatch.Run()
}
```

### consoleTable

```go
func consoleTable() {
	table.Run()
}
```

### consoleTab

```go
func consoleTab() {
	// tabContent items must be same length with tabs items
	// and according to similar order in tabs
	tabs := []string{"Lip Gloss", "Blush", "Eye Shadow", "Mascara", "Foundation"}
	tabContent := []string{"Lip Gloss Tab", "Blush Tab", "Eye Shadow Tab", "Mascara Tab", "Foundation Tab"}

	println("Use arrow keys to switch tabs. Press q to quit.")

	tab.Run(tabs, tabContent)
}
```

### consoleCountdown

```go
func consoleCountdown() {
	countdown.Run()
}
```

### consoleSequence

```go
func consoleSequence() {
	sequence.Run()
}
```

### consoleInline

```go
func consoleInline() {
	customMode := &inline.Mode{
		AltscreenMode: " altscreen mode ", // string will be displayed when switch to altscreen mode
		InlineMode:    " inline mode ",    // string will be displayed when switch to inline mode
		Key:           " ",                // space key to switch mode
	}
	inline.Run(customMode)
}
```

### consoleParallel

```go
func consoleParallel() {
	parallel.Run()
}
```

### consoleDir

```go
func consoleDir() {
	dir.Run()
}
```

### consoleGlamour

```go
func consoleGlamour() {
	const content = `
# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] René Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appétit!
`
	glamour.Run(content)
}
```

### ConsoleExample

```go
func ConsoleExample() {
	consoleCmd()
	consoleClear()
	consoleWithColor()
	consoleSelect()
	consoleHelp()

	consoleList()
	consolePaper()
	consoleProgress()
	consoleRealtime()
	consoleChoice()
	consoleSpinner()
	consoleSplit()
	consoleStopwatch()
	consoleTable()
	consoleTab()
	consoleCountdown()
	consoleSequence()
	consoleInline()
	consoleParallel()
	consoleDir()
	consoleGlamour()
}
```

### dateTime

```go
func dateTime() {
	println("Second:", date.Second())
	println("Minute:", date.Minute())
	println("Hour:", date.Hour())
	println("Day:", date.Day())
	println("Month:", date.Month())
	println("Year:", date.Year())
	println("Weekday:", date.Weekday())
	println("Unix:", date.Unix())
	println("UnixNano:", date.UnixNano())
	println("UnixMilli:", date.UnixMilli())
	println("UnixMicro:", date.UnixMicro())
	fmt.Println("UnixMilliToTime:", date.UnixMilliToTime(1000000000))
	fmt.Println("UnixMicroToTime:", date.UnixMicroToTime(1000000000))
	fmt.Println("UnixNanoToTime:", date.UnixNanoToTime(1000000000))
}
```

### dateISO

```go
func dateISO() {
	println("ISO:", date.ISO())
}
```

### dateShort

```go
func dateShort() {
	println("Short:", date.Short())
	println("ShortNormal:", date.ShortNormal())
	println("ShortReverse:", date.ShortReverse())
	println("ShortDash:", date.ShortDash())
	println("ShortDot:", date.ShortDot())
	println("ShortUnderscore:", date.ShortUnderscore())
	println("ShortSpace:", date.ShortSpace())
	println("ShortMonth:", date.ShortMonth())
}
```

### dateLong

```go
func dateLong() {
	println("Long:", date.Long())
}
```

### dateUTC

```go
func dateUTC() {
	println("UTC:", date.UTC())
}
```

### dateToSecond

```go
func dateToSecond() {
	println("ToSecond:", date.ToSecond(1))
}
```

### dateToMinute

```go
func dateToMinute() {
	println("ToMinute:", date.ToMinute(1))
}
```

### dateToHour

```go
func dateToHour() {
	println("ToHour:", date.ToHour(1))
}
```

### dateSleepSecond

```go
func dateSleepSecond() {
	date.SleepSecond(1)
}
```

### dateSleepMinute

```go
func dateSleepMinute() {
	date.SleepMinute(1)
}
```

### dateSleepHour

```go
func dateSleepHour() {
	date.SleepHour(1)
}
```

### DateExample

```go
func DateExample() {
	dateTime()

	dateISO()
	dateShort()
	dateLong()
	dateUTC()

	dateToSecond()
	dateToMinute()
	dateToHour()
	dateSleepSecond()
	dateSleepMinute()
	dateSleepHour()
}
```

### funcDelay

```go
func funcDelay() {
	println("Delay start:")

	result := function.DelayF(func() string {
		return "Delayed return after 2s"
	}, 2)

	if result.HasReturn {
		println(result.Value)
	} else {
		println("No result")
	}

	function.Delay(func() {
		println("Delayed not return")
	}, 3)
}
```

### funcRetry

```go
func funcRetry() {
	function.Retry(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}
```

### funcTimes

```go
func funcTimes() {
	function.Times(func() {
		println("Times")
	}, 3)
}
```

### funcInterval

```go
func funcInterval() {
	function.Interval(func() {
		println("Interval")
	}, 1)
}
```

### funcLock

```go
func funcLock() {
	oneInOneOutFuc := function.LockFunc(func(i interface{}) interface{} {
		return i
	}).(func(interface{}) interface{})("one input - one output")
	fmt.Println(oneInOneOutFuc)

	twoInTwoOutFunc1, twoInTwoOutFunc2 := function.LockFunc(func(i1, i2 interface{}) (interface{}, interface{}) {
		return i1, i2
	}).(func(interface{}, interface{}) (interface{}, interface{}))("two input - two output (a)", "two input - two output (b)")
	fmt.Println(twoInTwoOutFunc1, twoInTwoOutFunc2)

	function.LockFunc(func() {
		println("no input - no output")
	}).(func())()
}
```

### funcRunTime

```go
func funcRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task completed.")
	}

	duration := function.RunTime(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}
```

### FunctionExample

```go
func FunctionExample() {
	funcDelay()
	funcRetry()
	funcTimes()
	funcInterval()
	funcLock()
	funcRunTime()
}
```

### helperRandomID

```go
func helperRandomID() {
	println("Generate random ID: ", helper.RandomID())
}
```

### helperUUID

```go
func helperUUID() {
	println("New uuid: ", helper.UUID())
}
```

### helperAutoMdDoc

```go
func helperAutoMdDoc() {
	inputFilePath := "main.go"
	outputFilePath := "main.md"
	helper.AutoMdDoc(inputFilePath, outputFilePath)
}
```

### HelperExample

```go
func HelperExample() {
	helperRandomID()
	helperUUID()

	helperAutoMdDoc()
}
```

### ioCheckDir

```go
func ioCheckDir() {
	isExist, err1 := io.IsExistDir("tmp")
	if err1 != nil {
		println(err1.Error())
	}
	if isExist {
		println("dir exist")
	} else {
		println("dir not exist")
	}
}
```

### ioCreateDir

```go
func ioCreateDir() {
	err2 := io.CreateDir("tmp")
	if err2 != nil {
		println(err2.Error())
	}
	println("dir created")
}
```

### ioRemoveDir

```go
func ioRemoveDir() {
	err3 := io.RemoveDir("tmp")
	if err3 != nil {
		println(err3.Error())
	}
	println("dir removed")
}
```

### ioLsDir

```go
func ioLsDir() {
	data, err := io.LsDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}
```

### ioHierarchyDir

```go
func ioHierarchyDir() {
	data, err := io.HierarchyDir(".")
	if err != nil {
		println(err.Error())
		return
	}

	for _, v := range data {
		println(v)
	}
}
```

### ioCurrentDir

```go
func ioCurrentDir() {
	data, err := io.CurrentDir()
	if err != nil {
		println(err.Error())
		return
	}

	println(data)
}
```

### ioCheckFile

```go
func ioCheckFile() {
	isExist, err := io.IsExistFile("data.json")
	if err != nil {
		println(err.Error())
	}
	if isExist {
		println("file exist")
	} else {
		println("file not exist")
	}
}
```

### ioCreateFile

```go
func ioCreateFile() {
	err := io.CreateFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file created")
}
```

### ioRemoveFile

```go
func ioRemoveFile() {
	err := io.RemoveFile("data.json")
	if err != nil {
		println(err.Error())
	}
	println("file removed")
}
```

### ioReadFileByLine

```go
func ioReadFileByLine() {
	data, err := io.ReadFileByLine("main.go")
	if err != nil {
		println(err.Error())
	}
	for _, v := range data {
		println(v)
	}
}
```

### ioFileInfo

```go
func ioFileInfo() {
	data, err := io.FileInfo("main.go")
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("File info: %+v\n", data.All)
	fmt.Println("File info (with name):", data.Name)
	fmt.Printf("File info (with size): %d bytes\n", data.Size)
	fmt.Println("File info (with permissions):", data.Mode)
	fmt.Println("File info (with last modified):", data.ModTime)
	fmt.Println("File info (with directory check): ", data.IsDir)
	fmt.Printf("File info (with system process): %+v\n", data.Sys)
}
```

### ioRenameFile

```go
func ioRenameFile() {
	err := io.RenameFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file renamed")
}
```

### ioCopyFile

```go
func ioCopyFile() {
	err := io.CopyFile("data.json", "data2.json")
	if err != nil {
		println(err.Error())
	}
	println("file copied")
}
```

### ioTruncateFile

```go
func ioTruncateFile() {
	err := io.TruncateFile("data.json", 10)
	if err != nil {
		println(err.Error())
	}
	println("file truncated to 10 bytes")
}
```

### ioCleanFile

```go
func ioCleanFile() {
	err := io.CleanFile("data.json")
	if err != nil {
		println(err.Error())
	}

	// or using truncate with 0 bytes
	// err := helper.TruncateFile("data.json", 0)
	// if err != nil {
	// 	println(err.Error())
	// }

	println("file cleaned")
}
```

### ioWriteToFile

```go
func ioWriteToFile() {
	err := io.WriteFile("data.json", []string{"this is data 1", "this is data 2"})
	if err != nil {
		println(err.Error())
	}
	println("file written")
}
```

### ioAppendToFile

```go
func ioAppendToFile() {
	err := io.AppendFile("data.json", []string{"this is data 3", "this is data 4"})
	if err != nil {
		println(err.Error())
	}
	println("file appended")
}
```

### ioFileObj

```go
func ioFileObj() {
	type User struct {
		Name string
		Age  int
	}

	exampleFile := "data.json"

	// read file
	data, err := io.ReadFileObj[User](exampleFile)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("data: %+v\n", data)

	// write file
	updateData := append(data, User{
		Name: fmt.Sprintf("name %d", len(data)+1),
		Age:  len(data) + 1,
	})

	err2 := io.WriteFileObj(exampleFile, updateData)
	if err2 != nil {
		println(err2.Error())
	}
	println("data written")
}
```

### ioCreatePath

```go
func ioCreatePath() {
	relativePath := "tmp/example.txt"

	if err := io.CreatePath(relativePath); err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	println("File created successfully.")
}
```

### ioReadPath

```go
func ioReadPath() {
	relativePath := "tmp/example.txt"

	content, err := io.ReadPath(relativePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))
}
```

### ioWritePath

```go
func ioWritePath() {
	relativePath := "tmp/example.txt"

	newContent := []byte("This is a new content")

	if err := io.WritePath(relativePath, newContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	println("File updated successfully.")
}
```

### IOExample

```go
func IOExample() {
	ioCheckDir()
	ioCreateDir()
	ioRemoveDir()
	ioLsDir()
	ioHierarchyDir()
	ioCurrentDir()

	ioCheckFile()
	ioCreateFile()
	ioRemoveFile()
	ioReadFileByLine()
	ioFileInfo()
	ioRenameFile()
	ioCopyFile()
	ioTruncateFile()
	ioCleanFile()
	ioWriteToFile()
	ioAppendToFile()
	ioFileObj()

	ioCreatePath()
	ioReadPath()
	ioWritePath()
}
```

### mathIsPrime

```go
func mathIsPrime() {
	var num = 10
	println("Check prime number: ", math.IsPrime(num))
}
```

### mathIsEven

```go
func mathIsEven() {
	var num = 10
	println("Check even number: ", math.IsEven(num))
}
```

### mathIsOdd

```go
func mathIsOdd() {
	var num = 10
	println("Check odd number: ", math.IsOdd(num))
}
```

### mathIsPerfectSquare

```go
func mathIsPerfectSquare() {
	var num = 10
	println("Check perfect square number: ", math.IsPerfectSquare(num))
}
```

### mathAbs

```go
func mathAbs() {
	var num = -10
	println("Absolute number: ", math.Abs(num))
}
```

### mathFloor

```go
func mathFloor() {
	var num = 10.49
	println("Floor number: ", math.Floor(num))
}
```

### mathCeil

```go
func mathCeil() {
	var num = 10.49
	println("Ceil number: ", math.Ceil(num))
}
```

### mathRound

```go
func mathRound() {
	var num1, num2 = 10.49, 10.51
	println("Round number: ", math.Round(num1), math.Round(num2))
}
```

### mathMinMax

```go
func mathMinMax() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Max number: ", math.Max(num1, num2, num3, num4))
	println("Min number: ", math.Min(num1, num2, num3, num4))
}
```

### mathOperators

```go
func mathOperators() {
	var num1, num2, num3, num4 = 10, 20, 30, -2
	println("Sum of numbers: ", math.Sum(num1, num2, num3, num4))
	println("Add numbers: ", math.Add(num1, num2))
	println("Subtract of numbers: ", math.Sub(num1, num2))
	println("Multiply of numbers: ", math.Multi(num1, num2, num3, num4))
	println("Quotient of numbers: ", math.Divide(num1, num2))
	println("Remainder of numbers: ", math.Remainder(num1, num2))
	println("Average/Mean of numbers: ", math.Mean(num1, num2, num3, num4))
}
```

### mathPower

```go
func mathPower() {
	println("Power square of number: ", math.Pow2(2))
	println("Power of integer numbers: ", math.Pow(2, 3))
	println("Power of float numbers: ", math.PowF(2.0, 3.0))
}
```

### mathFactorial

```go
func mathFactorial() {
	var num = 5
	println("Factorial of number: ", math.Factorial(num))
}
```

### mathRoot

```go
func mathRoot() {
	println("Square-Root of integer number: ", math.Sqrt(16))
	println("Square-Root of float number: ", fmt.Sprintf("%f", math.SqrtF(20.0)))

	println("Cube-Root of integer number: ", math.Cbrt(27))
	println("Cube-Root of float number: ", fmt.Sprintf("%f", math.CbrtF(20.0)))

	println("Nth-Root of integer number: ", math.Root(4, 2))
	println("Nth-Root of float number: ", fmt.Sprintf("%f", math.RootF(20.0, 3.0)))
}
```

### mathLog

```go
func mathLog() {
	println("Logarithm of integer number: ", math.Log(16, 2))
	println("Logarithm of float number: ", fmt.Sprintf("%f", math.LogF(20.0, 2.0)))

	println("Logarithm of integer number (base 2): ", math.Log2(16))
	println("Logarithm of float number (base 2): ", fmt.Sprintf("%f", math.Log2F(20.0)))

	println("Logarithm of integer number (base 10): ", math.Log10(100))
	println("Logarithm of float number (base 10): ", fmt.Sprintf("%f", math.Log10F(20.0)))
}
```

### mathPytago

```go
func mathPytago() {
	println("Pytago of number (integer): ", fmt.Sprintf("%f", math.Pytago(3, 4)))
	println("Pytago of number (float): ", fmt.Sprintf("%f", math.PytagoF(3.0, 4.0)))
}
```

### mathTrigonometry

```go
func mathTrigonometry() {
	println("Sine of integer number: ", math.Sin(90))
	println("Sine of float number: ", fmt.Sprintf("%f", math.SinF(90.0)))
	println("Cosine of integer number: ", math.Cos(90))
	println("Cosine of float number: ", fmt.Sprintf("%f", math.CosF(90.0)))
	println("Tangent of integer number: ", math.Tan(90))
	println("Tangent of float number: ", fmt.Sprintf("%f", math.TanF(90.0)))
}
```

### mathRect

```go
func mathRect() {
	println("Area of rectangle: ", math.AreaRect(10, 20))
	println("Perimeter of rectangle (integer): ", math.PeriRect(10, 20))
	println("Perimeter of rectangle (float): ", fmt.Sprintf("%f", math.PeriRectF(10.0, 20.0)))
	println("Diagonal of rectangle (integer): ", fmt.Sprintf("%f", math.DiagRect(10, 20)))
	println("Diagonal of rectangle (float): ", fmt.Sprintf("%f", math.DiagRectF(10.0, 20.0)))
	println("Volume of rectangular (integer): ", math.VolRect(10, 20, 30))
	println("Volume of rectangular (float): ", fmt.Sprintf("%f", math.VolRectF(10.0, 20.0, 30.0)))
}
```

### mathCircle

```go
func mathCircle() {
	println("Area of circle (integer): ", fmt.Sprintf("%f", math.AreaCircle(10)))
	println("Area of circle (float): ", fmt.Sprintf("%f", math.AreaCircleF(10.0)))
	println("Perimeter of circle (integer): ", fmt.Sprintf("%f", math.PeriCircle(10)))
	println("Perimeter of circle (float): ", fmt.Sprintf("%f", math.PeriCircleF(10.0)))
}
```

### mathTriangle

```go
func mathTriangle() {
	println("Area of triangle (integer): ", math.AreaTriangle(10, 20))
	println("Area of triangle (float): ", fmt.Sprintf("%f", math.AreaTriangleF(10.0, 20.0)))
	println("Perimeter of triangle (integer): ", math.PeriTriangle(10, 20, 30))
	println("Perimeter of triangle (float): ", fmt.Sprintf("%f", math.PeriTriangleF(10.0, 20.0, 30.0)))
}
```

### mathSquare

```go
func mathSquare() {
	println("Area of square (integer): ", math.AreaSquare(10))
	println("Area of square (float): ", fmt.Sprintf("%f", math.AreaSquareF(10.0)))
	println("Perimeter of square (integer): ", math.PeriSquare(10))
	println("Perimeter of square (float): ", fmt.Sprintf("%f", math.PeriSquareF(10.0)))
}
```

### mathCube

```go
func mathCube() {
	println("Area of cube (integer): ", math.AreaCube(10))
	println("Area of cube (float): ", fmt.Sprintf("%f", math.AreaCubeF(10.0)))
	println("Perimeter of cube (integer): ", math.PeriCube(10))
	println("Perimeter of cube (float): ", fmt.Sprintf("%f", math.PeriCubeF(10.0)))
	println("Volume of cube (integer): ", math.VolCube(10))
	println("Volume of cube (float): ", fmt.Sprintf("%f", math.VolCubeF(10.0)))
}
```

### mathSphere

```go
func mathSphere() {
	println("Area of sphere (integer): ", fmt.Sprintf("%f", math.AreaSphere(10)))
	println("Area of sphere (float): ", fmt.Sprintf("%f", math.AreaSphereF(10.0)))
	println("Volume of sphere (integer): ", fmt.Sprintf("%f", math.VolSphere(10)))
	println("Volume of sphere (float): ", fmt.Sprintf("%f", math.VolSphereF(10.0)))
}
```

### mathCylinder

```go
func mathCylinder() {
	println("Area of cylinder (integer): ", fmt.Sprintf("%f", math.AreaCylinder(10, 20)))
	println("Area of cylinder (float): ", fmt.Sprintf("%f", math.AreaCylinderF(10.0, 20.0)))
	println("Volume of cylinder (integer): ", fmt.Sprintf("%f", math.VolCylinder(10, 20)))
	println("Volume of cylinder (float): ", fmt.Sprintf("%f", math.VolCylinderF(10.0, 20.0)))
}
```

### mathCone

```go
func mathCone() {
	println("Area of cone (integer): ", fmt.Sprintf("%f", math.AreaCone(10, 20)))
	println("Area of cone (float): ", fmt.Sprintf("%f", math.AreaConeF(10.0, 20.0)))
	println("Volume of cone (integer): ", fmt.Sprintf("%f", math.VolCone(10, 20)))
	println("Volume of cone (float): ", fmt.Sprintf("%f", math.VolConeF(10.0, 20.0)))
}
```

### mathTrapezoid

```go
func mathTrapezoid() {
	println("Area of trapezoid (integer): ", fmt.Sprintf("%f", math.AreaTrapezoid(10, 20, 30)))
	println("Area of trapezoid (float): ", fmt.Sprintf("%f", math.AreaTrapezoidF(10.0, 20.0, 30.0)))
}
```

### mathParallelogram

```go
func mathParallelogram() {
	println("Area of parallelogram (integer): ", math.AreaParallelogram(10, 20))
	println("Area of parallelogram (float): ", fmt.Sprintf("%f", math.AreaParallelogramF(10.0, 20.0)))
}
```

### mathRhombus

```go
func mathRhombus() {
	println("Area of rhombus (integer): ", math.AreaRhombus(10, 20))
	println("Area of rhombus (float): ", fmt.Sprintf("%f", math.AreaRhombusF(10.0, 20.0)))
}
```

### mathEllipse

```go
func mathEllipse() {
	println("Area of ellipse (integer): ", fmt.Sprintf("%f", math.AreaEllipse(10, 20)))
	println("Area of ellipse (float): ", fmt.Sprintf("%f", math.AreaEllipseF(10.0, 20.0)))
}
```

### mathPolygon

```go
func mathPolygon() {
	println("Area of pentagol by number of sides (integer): ", fmt.Sprintf("%f", math.AreaPolygon(10, 6)))
}
```

### mathTransition

```go
func mathTransition() {
	var distance, speed, time float64 = 100, 10, 10
	println("Speed: ", fmt.Sprintf("%f", math.Speed(distance, time)))
	println("Distance: ", fmt.Sprintf("%f", math.Distance(speed, time)))
	println("Time: ", fmt.Sprintf("%f", math.Time(distance, speed)))
}
```

### MathExample

```go
func MathExample() {
	mathIsPrime()
	mathIsEven()
	mathIsOdd()
	mathIsPerfectSquare()

	mathAbs()
	mathFloor()
	mathCeil()
	mathRound()
	mathMinMax()
	mathOperators()
	mathPower()
	mathFactorial()
	mathRoot()

	mathLog()
	mathPytago()
	mathTrigonometry()
	mathTransition()

	mathRect()
	mathCircle()
	mathTriangle()
	mathSquare()
	mathCube()
	mathSphere()
	mathCylinder()
	mathCone()
	mathTrapezoid()
	mathParallelogram()
	mathRhombus()
	mathEllipse()
	mathPolygon()
}
```

### netOpen

```go
func netOpen() {
	net.Open("https://google.com")
}
```

### netEncode

```go
func netEncode() {
	println("Encode: ", net.Encode("https://google.com"))
}
```

### netDecode

```go
func netDecode() {
	println("Decode: ", net.Decode("https%3A%2F%2Fgoogle.com"))
}
```

### netCheck

```go
func netCheck() {
	ok, err := net.Check("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Response: ", ok)
}
```

### netCheckWithStatusCode

```go
func netCheckWithStatusCode() {
	statusCode, err := net.CheckWithStatusCode("https://google.com")
	if err != nil {
		panic(err)
	}
	println("Status code: ", statusCode)
}
```

### netHeader

```go
func netHeader() {
	header, err := net.Header("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(types.MapAsString(header))
}
```

### netConnectTime

```go
func netConnectTime() {
	connectTime, err := net.ConnectTime("https://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connect time: %fs\n", connectTime)
}
```

### NetExample

```go
func NetExample() {
	netOpen()
	netEncode()
	netDecode()
	netCheck()
	netCheckWithStatusCode()
	netHeader()
	netConnectTime()
}
```

### numRandom

```go
func numRandom() {
	random := number.Random(1, 10)
	println("Random number [1, 10): ", random)
}
```

### numClamp

```go
func numClamp() {
	println("Clamp number: ", number.Clamp(5, 1, 10))
}
```

### numInRange

```go
func numInRange() {
	println("Check number is in range: ", number.InRange(5, 1, 10))
}
```

### NumberExample

```go
func NumberExample() {
	numRandom()
	numClamp()
	numInRange()
}
```

### regexIsMatch

```go
func regexIsMatch() {
	fmt.Println("Match string with regex: ", regex.IsMatch(`[a-z]+\s[a-z]+`, "hello world"))
}
```

### regexMatch

```go
func regexMatch() {
	fmt.Println("Match string with regex: ", regex.Match(`[A-Z]`, "Hello World 123"))
}
```

### regexMatchIndex

```go
func regexMatchIndex() {
	paragraph := "I think Ruth's dog is cuter than your dog!"
	matchIdx := regex.MatchIndex(`[^\w\s']`, paragraph)
	if matchIdx != -1 {
		fmt.Printf("Match with regex (index: %d, value: %s)\n", matchIdx, string(paragraph[matchIdx]))
	} else {
		println("Not found index match regex")
	}
}
```

### RegexExample

```go
func RegexExample() {
	regexIsMatch()
	regexMatch()
	regexMatchIndex()
}
```

### strCapitalize

```go
func strCapitalize() {
	var str = "hello world"
	println(strings.Capitalize(str))
}
```

### strCamelCase

```go
func strCamelCase() {
	var str = "convert_any-string TO-camelCase in Go! 123"
	println("Convert string to Camel Case: ", strings.CamelCase(str))
}
```

### strSnakeCase

```go
func strSnakeCase() {
	var str1 = "hello world"
	var str2 = "ConvertThisString123"
	println("Convert string to Snake Case: ", strings.SnakeCase(str1), strings.SnakeCase(str2))
}
```

### strKebabCase

```go
func strKebabCase() {
	var str = "hello world"
	println("Convert string to Kebab Case: ", strings.KebabCase(str))
}
```

### strIsLetter

```go
func strIsLetter() {
	var str = "hello world"
	println("Check is letter character: ", strings.IsLetter(str[0]))
}
```

### strIsDigit

```go
func strIsDigit() {
	var str = "1hello world"
	println("Check is number character: ", strings.IsDigit(str[0]))
}
```

### strIncludes

```go
func strIncludes() {
	var str = "hello world, this is world"
	println("Check substring in string: ", strings.Includes(str, "world"))
}
```

### strIsLower

```go
func strIsLower() {
	var str = "hELLO WORLD"
	println("Check is lower string: ", strings.IsLower(str[0]))
}
```

### strIsUpper

```go
func strIsUpper() {
	var str = "Hello world"
	println("Check is upper string: ", strings.IsUpper(str[0]))
}
```

### strSplit

```go
func strSplit() {
	var str = "hello world"
	fmt.Println("Split string by separator: ", strings.Split(str, "l"))
}
```

### strWords

```go
func strWords() {
	var str = "hello world"
	println("Split string to array of words: ", strings.Words(str))

}
```

### strReverse

```go
func strReverse() {
	var str = "hello world"
	println("Reverse string: ", strings.Reverse(str))
}
```

### strLower

```go
func strLower() {
	var str = "HELLO WORLD"
	println("Lower string (string): ", strings.Lowers(str))
	println("Lower string (byte): ", strings.Lower(str[0]))
	println("Lower first string: ", strings.LowerFirst(str))
}
```

### strUpper

```go
func strUpper() {
	var str = "hello world"
	println("Upper string (string): ", strings.Uppers(str))
	println("Upper string (byte): ", strings.Upper(str[0]))
	println("Upper first string: ", strings.UpperFirst(str))
}
```

### strRepeat

```go
func strRepeat() {
	var str = "hello world"
	println("Repeat string: ", strings.Repeat(str, 3))
}
```

### strTruncate

```go
func strTruncate() {
	var str = "hello world"
	println("Truncate string (default): ", strings.Truncate(str, 5))
	println("Truncate string (custom): ", strings.Truncate(str, 5, "***"))
}
```

### strReplace

```go
func strReplace() {
	var str = "hello world, this is world"
	println("Replace string: ", strings.Replace(str, "world", "golang"))
}
```

### strTrim

```go
func strTrim() {
	var str = "   hello world, this is world   "
	println("Trim string: ", strings.Trim(str))
	println("Trim left string: ", strings.LTrim(str))
	println("Trim right string: ", strings.RTrim(str))
}
```

### strTrimBlank

```go
func strTrimBlank() {
	println("Trim blank string: ", strings.TrimBlank("   hello world, this is world   "))
	println("Trim left blank string: ", strings.TrimBlank("   hello world, this is world   \t"))
	println("Trim right blank string: ", strings.TrimBlank("   hello world, this is world   \n"))
}
```

### strTrimPrefix

```go
func strTrimPrefix() {
	var str = "   hello world, this is world   "
	println("Trim prefix string: ", strings.TrimPrefix(str, "   "))
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}
```

### strTrimSuffix

```go
func strTrimSuffix() {
	var str = "   hello world, this is world   "
	println("Trim suffix string: ", strings.TrimSuffix(str, "   "))
}
```

### strJoin

```go
func strJoin() {
	var str = []string{"hello", "world"}
	println("Join string: ", strings.Join(str, "-"))
}
```

### strConcat

```go
func strConcat() {
	println("Concat string: ", strings.Concat("hello", "world"))
}
```

### strSubStr

```go
func strSubStr() {
	var str = "hello world, this is world"
	println("Sub string: ", strings.SubStr(str, 0, 5))
	println("Sub string: ", strings.SubStr(str, 0, 1))
	println("Sub string (only start): ", strings.SubStr(str, -5))
	println("Sub string (with negative index): ", strings.SubStr(str, -5, -1))
}
```

### strSlice

```go
func strSlice() {
	var str = "hello world, this is world"
	println("Slice string: ", strings.Slice(str, 0, 5))
	println("Slice string: ", strings.Slice(str, 0, 1))
	println("Slice string (only start): ", strings.Slice(str, -5))
	println("Slice string (not parameters): ", strings.Slice(str))
	println("Slice string (with negative index): ", strings.Slice(str, -5, -1))
}
```

### strSplice

```go
func strSplice() {
	var str = "helloworld, this is world"
	println("Splice string (default not replace): ", strings.Splice(str, 0, 5))
	println("Splice string (with replace): ", strings.Splice(str, 1, 5, "golang"))
	println("Splice string (with replace multiple): ", strings.Splice(str, 1, 5, "golang1", "golang2"))
}
```

### strStartsWith

```go
func strStartsWith() {
	var str = "hello world, this is world"
	println("Starts with: ", strings.StartsWith(str, "hello"))
}
```

### strEndsWith

```go
func strEndsWith() {
	var str = "hello world, this is world"
	println("Ends with: ", strings.EndsWith(str, "world"))
}
```

### strEscape

```go
func strEscape() {
	var str = "This is a <b>bold</b> statement & \"quote\" – © ®"
	println("Escape string: ", strings.Escape(str))
}
```

### strUnescape

```go
func strUnescape() {
	var str = "This is a &lt;b&gt;bold&lt;/b&gt; statement &amp; &quot;quote&quot; – © ®"
	println("Unescape string: ", strings.Unescape(str))
}
```

### strPad

```go
func strPad() {
	var str = "hello world"
	println("Pad-left string: ", strings.PadStart(str, 20, '$'))
	println("Pad-right string: ", strings.PadEnd(str, 20, '@'))
}
```

### strCount

```go
func strCount() {
	var str = "hello world wo wo"
	println("Count words/substr in string (default): ", strings.Count(str))
	println("Count words/substr in string (with char): ", strings.Count(str, "wo"))
}
```

### strLines

```go
func strLines() {
	var str = "hello world\nwo wo"
	println("Count lines of string: ", strings.Lines(str))
}
```

### strIndex

```go
func strIndex() {
	var str = "hello world, this is world"

	f1, l1 := strings.Index(str, "l")
	fmt.Printf("First index start at: %d, end at: %d\n", f1, l1)

	f2, l2 := strings.Index(str, "world")
	fmt.Printf("First index start at: %d, end at: %d\n", f2, l2)

	f3 := strings.FIndex(str, "l")
	fmt.Println("First index start at: ", f3)

	l3 := strings.LIndex(str, "world")
	fmt.Println("Last index start at: ", l3)

	f4, l4 := strings.Index(str, "oo")
	if f4 == -1 || l4 == -1 {
		fmt.Println("Not found")
	}

	if strings.FIndex(str, "oo") == -1 {
		fmt.Println("Not found")
	}

	if strings.LIndex(str, "oo") == -1 {
		fmt.Println("Not found")
	}
}
```

### strRandom

```go
func strRandom() {
	println("Random string: ", strings.Random(10))
}
```

### strAt

```go
func strAt() {
	var str = "hello world"
	println("At string: ", strings.At(str, 1))
	println("At string: ", strings.At(str, -5))
}
```

### strCodePointAt

```go
func strCodePointAt() {
	var str = "hello world"
	println("Code point at string: ", strings.CodePointAt(str, 1))
	println("Code point at string: ", strings.CodePointAt(str, -5))
}
```

### strCodePoint

```go
func strCodePoint() {
	asciiValues := strings.CodePoint("hello world")

	print("Code point string: ")
	for _, asciiValue := range asciiValues {
		fmt.Printf("%d ", asciiValue)
	}
}
```

### strFromCodePointAt

```go
func strFromCodePointAt() {
	println("From code point at string: ", strings.FromCodePointAt(9733))
	println("From code point at string: ", strings.FromCodePointAt(9731))
}
```

### strFromCodePoint

```go
func strFromCodePoint() {
	println("From code point string: ", strings.FromCodePoint(104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100))
}
```

### StringExample

```go
func StringExample() {
	strCapitalize()
	strCamelCase()
	strSnakeCase()
	strKebabCase()

	strIsLetter()
	strIsDigit()
	strIncludes()
	strIsLower()
	strIsUpper()

	strSplit()
	strWords()
	strReverse()
	strLower()
	strUpper()
	strRepeat()
	strTruncate()
	strReplace()
	strTrim()
	strTrimBlank()
	strTrimPrefix()
	strTrimSuffix()
	strJoin()
	strConcat()
	strSubStr()
	strSlice()
	strSplice()
	strStartsWith()
	strEndsWith()
	strEscape()
	strUnescape()
	strPad()

	strCount()
	strLines()
	strIndex()
	strRandom()
	strAt()
	strCodePointAt()
	strCodePoint()
	strFromCodePointAt()
	strFromCodePoint()
}
```

### structMerge

```go
func structMerge() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	address := Address{
		City:    "New York",
		Street:  "123 Main St",
		ZipCode: "10001",
	}

	merged := structs.Merge(person, address)

	fmt.Printf("Struct after merged: %+v\n", merged)

	fmt.Println("Name:", merged.(map[string]interface{})["Name"])

	fmt.Println("City:", merged.(map[string]interface{})["City"])
}
```

### structRemove

```go
func structRemove() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Struct after removed field: %+v\n", structs.Remove(person, "Email"))
}
```

### structAdd

```go
func structAdd() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	newFields := map[string]interface{}{
		"Address": "123 Main St",
		"Phone":   "555-1234",
	}
	result := structs.Add(person, newFields)

	fmt.Printf("Struct after adding fields: %+v\n", result)
}
```

### structSet

```go
func structSet() {
	person := &Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	structs.Set(person, "Name", "Updated Name")

	fmt.Printf("Struct after setting field: %+v\n", person)
}
```

### structGet

```go
func structGet() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	name := structs.Get(person, "Name")
	fmt.Printf("Name: %s\n", name)
}
```

### structClone

```go
func structClone() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "example@gmail.com",
	}

	fmt.Printf("Original: %+v\n", person)

	clone := structs.Clone(person)

	updateClone := clone.(Person)
	updateClone.Name = "Updated Name"
	fmt.Printf("Clone: %+v\n", updateClone)
}
```

### structHas

```go
func structHas() {
	person := Person{
		Name:  "Example",
		Age:   40,
		Email: "",
	}

	has := structs.Has(person, "Email")
	fmt.Printf("Has: %+v\n", has)

	hasEmpty := structs.HasEmpty(person, "Email")
	fmt.Printf("Has empty: %+v\n", hasEmpty)
}
```

### StructExample

```go
func StructExample() {
	structMerge()
	structRemove()
	structAdd()
	structSet()
	structGet()

	structClone()
	structHas()
}
```
