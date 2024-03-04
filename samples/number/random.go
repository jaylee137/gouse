package number

import "github.com/thuongtruong1009/gouse/number"

func SampleNumRandom() {
	random := number.Random(1, 10)
	println("Random number [1, 10): ", random)
}
