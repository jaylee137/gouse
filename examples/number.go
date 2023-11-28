package examples

import "github.com/thuongtruong1009/gouse/number"

func numRandom() {
	println("Random number: ", number.Random(1, 10))
}

func numClamp() {
	println("Clamp number: ", number.Clamp(5, 1, 10))
}

func numInRange() {
	println("Check number is in range: ", number.InRange(5, 1, 10))
}

func NumberExample() {
	numRandom()
	numClamp()
	numInRange()
}
