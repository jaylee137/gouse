package arrs

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/array"
)

func ArrEqual() {
	fmt.Println("Compare equal: ", arrs.Equal(1, 1))
}

func ArrExist(){
	fmt.Println("Exist with generic type: ", arrs.IsExist([]int{1, 2, 3}, 1))
}

func ArrMinMax() {
	fmt.Println("Min with generic type: ", arrs.Min([]int{1, 2, 3}))

	fmt.Println("Max with generic type: ", arrs.Max([]int{1, 2, 3}))
}

func ArrOps() {
	fmt.Println("Most frequency with generic type: ", arrs.Most([]int{1, 2, 3, 2, 2, 1, 2, 3}))

	// fmt.Println("Sum with generic type: ", arrs.Sum([]int{1, 2, 3}))
}

func Example() {
	ArrEqual()
	ArrExist()
	ArrMinMax()
	ArrOperations()
}