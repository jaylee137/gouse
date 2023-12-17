# Date

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