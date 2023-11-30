package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/function"
)

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

func funcLock() {
	tmpOneOne := function.LockOneInOneOut(func(i interface{}) interface{} {
		return i
	})
	fmt.Println(tmpOneOne(1))

	tmpOneTwo := function.LockOneInTwoOut(func(i interface{}) (interface{}, interface{}) {
		return i, i
	})
	fmt.Println(tmpOneTwo(1))

	tmpTwoOne := function.LockTwoInOneOut(func(i1, i2 interface{}) interface{} {
		return i1
	})
	fmt.Println(tmpTwoOne(1, 2))

	tmpTwoTwo := function.LockTwoInTwoOut(func(i1, i2 interface{}) (interface{}, interface{}) {
		return i1, i2
	})
	fmt.Println(tmpTwoTwo(1, 2))
}

func FunctionExample() {
	// funcDelay()
	// funcRetry()
	// funcTimes()
	// funcInterval()
	funcLock()
}
