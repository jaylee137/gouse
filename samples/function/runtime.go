package function

import (
	"fmt"
	"time"

	"github.com/thuongtruong1009/gouse/function"
)

func funcRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task completed.")
	}

	duration := function.RunTime(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}
