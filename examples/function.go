package examples

import "github.com/thuongtruong1009/gouse/function"

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

func funcRetry() {
	function.Retry(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}

func funcTimes() {
	function.Times(func() {
		println("Times")
	}, 3)
}

func funcInterval() {
	function.Interval(func() {
		println("Interval")
	}, 1)
}

func FunctionExample() {
	funcDelay()
	funcRetry()
	funcTimes()
	funcInterval()
}
