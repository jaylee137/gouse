package samples

import "github.com/thuongtruong1009/gouse/number"

func numRandom() {
	random := number.Random(1, 10)
	println("Random number [1, 10): ", random)
}

func numClamp() {
	println("Clamp number: ", number.Clamp(5, 1, 10))
}

func numInRange() {
	println("Check number is in range: ", number.InRange(5, 1, 10))
}
