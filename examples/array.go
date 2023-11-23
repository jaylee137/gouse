package examples

import (
	"fmt"
	"github.com/thuongtruong1009/gouse/array"
)

func arrEqual() {
	fmt.Println("--- Compare equal ---")
	fmt.Println("[int]: ", array.Equal(1, 1))
	fmt.Println("[uint]: ", array.Equal(uint(1), uint(1)))
	fmt.Println("[float]: ", array.Equal(1.2, 1.1))
	fmt.Println("[string]: ", array.Equal("1", "0"))
	fmt.Println("[rune]: ", array.Equal('a', 'a'))
	fmt.Println("[bool]: ", array.Equal(true, true))
	fmt.Println("[complex]: ", array.Equal(1 + 2i, 1 + 2i))
	fmt.Println("[struct]: ", array.Equal(struct{a int}{1}, struct{a int}{1}))
}

func arrIncludes(){
	fmt.Println("--- Check element is exist in array ---")
	fmt.Println("[int]: ", array.Includes([]int{1, -2, 3}, 1))
	fmt.Println("[uint]: ", array.Includes([]uint{1, 2, 3}, 1))
	fmt.Println("[float]: ", array.Includes([]float64{1.2, 2.3, 3.4}, 1.2))
	fmt.Println("[string]: ", array.Includes([]string{"1", "2", "3"}, "0"))
	fmt.Println("[rune]: ", array.Includes([]rune{'a', 'b', 'c'}, 'a'))
	fmt.Println("[complex]: ", array.Includes([]complex128{1 + 2i, 2 + 3i}, 1 + 2i))
	fmt.Println("[struct]: ", array.Includes([]struct{a int}{struct{a int}{1}, struct{a int}{2}}, struct{a int}{3}))
	fmt.Println("[interface]: ", array.Includes([]interface{}{1, "2", true}, 1))
}

func arrMost() {
	fmt.Println("--- Most frequency in array ---")
	fmt.Println("[int]: ", array.Most([]int{1, -2, 3, 2, 2, 1, 2, 3}))
	fmt.Println("[uint]: ", array.Most([]uint{1, 2, 3, 2, 2, 1, 2, 3}))
	fmt.Println("[float]: ", array.Most([]float64{1.2, 2.3, 3.4, 2.3, 2.3, 1.2, 2.3, 3.4}))
	fmt.Println("[string]: ", array.Most([]string{"1", "2", "3", "2", "2", "1", "2", "3"}))
	fmt.Println("[rune]: ", string(array.Most([]rune{'a', 'b', 'c', 'b', 'b', 'a', 'b', 'c'})))
	fmt.Println("[complex]: ", array.Most([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 2 + 3i, 2 + 3i, 1 + 2i, 2 + 3i, 3 + 4i}))
	fmt.Println("[struct]: ", array.Most([]struct{a int}{struct{a int}{1}, struct{a int}{2}, struct{a int}{3}, struct{a int}{2}, struct{a int}{2}, struct{a int}{1}, struct{a int}{2}, struct{a int}{3}}))
	fmt.Println("[interface]: ", array.Most([]interface{}{1, "2", true, "2", "2", 1, "2", true}))
}

func arrSum() {
	fmt.Println("--- Sum elements in array ---")
	fmt.Println("[int]: ", array.Sum([]int{1, -2, 3}))
	fmt.Println("[uint]: ", array.Sum([]uint{1, 2, 3}))
	fmt.Println("[float]: ", array.Sum([]float64{1.2, 2.3, 3.4}))
	fmt.Println("[rune]: ", array.Sum([]rune{'a', 'b', 'c'}))
	fmt.Println("[complex]: ", array.Sum([]complex128{1 + 2i, 2 + 3i, 3 + 4i}))
}

func arrChunk() {
	fmt.Println("--- Chunk array ---")
	fmt.Println("[int]: ", array.Chunk([]int{1, -2, 3, -4, 5, 6}, 3))
	fmt.Println("[uint]: ", array.Chunk([]uint{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println("[float]: ", array.Chunk([]float64{1.2, 2.3, 3.4, 4.5, 5.6, 6.7}, 3))
	fmt.Println("[string]: ", array.Chunk([]string{"1", "2", "3", "4", "5", "6"}, 3))
	fmt.Println("[rune]: ", array.Chunk([]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3))
	fmt.Println("[complex]: ", array.Chunk([]complex128{1 + 2i, 2 + 3i, 3 + 4i, 4 + 5i, 5 + 6i, 6 + 7i}, 3))
	fmt.Println("[struct]: ", array.Chunk([]struct{a int}{struct{a int}{1}, struct{a int}{2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, 3))
	fmt.Println("[interface]: ", array.Chunk([]interface{}{1, "2", true, "4", "5", 6}, 3))
}

func arrDiff() {
	fmt.Println("--- Difference array ---")
	fmt.Println("[int]: ", array.Diff([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println("[uint]: ", array.Diff([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	fmt.Println("[float]: ", array.Diff([]float64{1.2, 2.3, 3.4}, []float64{4.5, 5.6, 6.7}))
	fmt.Println("[string]: ", array.Diff([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	fmt.Println("[rune]: ", array.Diff([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	fmt.Println("[complex]: ", array.Diff([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	fmt.Println("[struct]: ", array.Diff([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, []struct{a int}{struct{a int}{1}, struct{a int}{2}}))
	fmt.Println("[interface]: ", array.Diff([]interface{}{1, "2", false}, []interface{}{3, "3", true}))
}

func arrDrop() {
	fmt.Println("--- Drop elements in array (default n = 1) ---")
	fmt.Println("[int] with default: ", array.Drop([]int{1, -2, 3, -4, 5, 6}))
	fmt.Println("[int]: ", array.Drop([]int{1, -2, 3, -4, 5, 6}, 2))
	fmt.Println("[uint]: ", array.Drop([]uint{1, 2, 3, 4, 5, 7}, 2))
	fmt.Println("[float]: ", array.Drop([]float64{1.2, 2.3, 3.4}, 2))
	fmt.Println("[string]: ", array.Drop([]string{"1", "4", "5", "6"}, 2))
	fmt.Println("[rune]: ", array.Drop([]rune{'a', 'b', 'd', 'e', 'f'}, 2))
	fmt.Println("[complex]: ", array.Drop([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 2))
	fmt.Println("[struct]: ", array.Drop([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, 2))
	fmt.Println("[interface]: ", array.Drop([]interface{}{1, "2", false}, 2))
}

func arrIndex() {
	fmt.Println("--- Index of element in array ---")
	fmt.Println("[int]: ", array.Index([]int{1, -2, 3, -4, 5, 6}, 3))
	fmt.Println("[uint]: ", array.Index([]uint{1, 2, 3, 4, 5, 7}, 3))
	fmt.Println("[float]: ", array.Index([]float64{1.2, 2.3, 3.4}, 3.4))
	fmt.Println("[string]: ", array.Index([]string{"1", "4", "5", "6"}, "5"))
	fmt.Println("[rune]: ", array.Index([]rune{'a', 'b', 'd', 'e', 'f'}, 'e'))
	fmt.Println("[complex]: ", array.Index([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, 5 + 6i))
	fmt.Println("[struct]: ", array.Index([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, struct{a int}{3}))
	fmt.Println("[interface]: ", array.Index([]interface{}{1, "2", false}, false))
}

func arrMerge() {
	fmt.Println("--- Merge arrays ---")
	fmt.Println("[int]: ", array.Merge([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, []int{1, -2, 3, -4, 5, 6}))
	fmt.Println("[uint]: ", array.Merge([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	fmt.Println("[float]: ", array.Merge([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	fmt.Println("[string]: ", array.Merge([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	fmt.Println("[rune]: ", array.Merge([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	fmt.Println("[complex]: ", array.Merge([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	fmt.Println("[struct]: ", array.Merge([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, []struct{a int}{struct{a int}{1}, struct{a int}{2}}))
	fmt.Println("[interface]: ", array.Merge([]interface{}{1, "2", false}, []interface{}{3, "3", true}))
}

func arrMin() {
	fmt.Println("--- Min element in array ---")
	fmt.Println("[int]: ", array.Min([]int{1, -2, 3}))
	fmt.Println("[uint]: ", array.Min([]uint{1, 2, 3}))
	fmt.Println("[string]: ", array.Min([]string{"z", "d", "m"}))
	fmt.Println("[rune]: ", string(array.Min([]rune{'z', 'd', 'm'})))
	fmt.Println("[float]: ", array.Min([]float64{1.2, 2.3, 3.4}))
}

func arrMax() {
	fmt.Println("--- Max element in array ---")
	fmt.Println("[int]: ", array.Max([]int{1, -2, 3}))
	fmt.Println("[uint]: ", array.Max([]uint{1, 2, 3}))
	fmt.Println("[string]: ", array.Max([]string{"z", "d", "m"}))
	fmt.Println("[rune]: ", string(array.Max([]rune{'z', 'd', 'm'})))
	fmt.Println("[float]: ", array.Max([]float64{1.2, 2.3, 3.4}))
}

func arrIntersect() {
	fmt.Println("--- Intersection arrays ---")
	fmt.Println("[int]: ", array.Intersect([]int{1, -2, 3, -4, 5, 6}, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println("[uint]: ", array.Intersect([]uint{1, 2, 3, 4, 5, 7}, []uint{1, 2, 3, 4, 5, 6}))
	fmt.Println("[float]: ", array.Intersect([]float64{1.2, 2.3, 3.4}, []float64{1.2, 4.5, 5.6, 6.7}))
	fmt.Println("[string]: ", array.Intersect([]string{"1", "4", "5", "6"}, []string{"1", "2", "3", "6"}))
	fmt.Println("[rune]: ", array.Intersect([]rune{'a', 'b', 'd', 'e', 'f'}, []rune{'a', 'b', 'c', 'f'}))
	fmt.Println("[complex]: ", array.Intersect([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, []complex128{1 + 2i, 2 + 3i, 6 + 7i}))
	fmt.Println("[struct]: ", array.Intersect([]struct{a int}{struct{a int}{1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, []struct{a int}{struct{a int}{1}, struct{a int}{2}}))
	fmt.Println("[interface]: ", array.Intersect([]interface{}{1, "2", true}, []interface{}{3, "3", true}))
}

func arrCbKey() {
	fmt.Println("--- Find key of element pass condition in callback function ---")
	fmt.Println("[int]: ", array.CbKey([]int{1, -2, 3, -4, 5, 6}, func (v int) bool {
		return v == 3
	}))

	fmt.Println("[uint]: ", array.CbKey([]uint{1, 2, 3, 4, 5, 7}, func (v uint) bool {
		return v == 3
	}))

	fmt.Println("[float]: ", array.CbKey([]float64{1.2, 2.3, 3.4}, func (v float64) bool {
		return v == 3.4
	}))

	fmt.Println("[string]: ", array.CbKey([]string{"1", "4", "5", "6"}, func (v string) bool {
		return v == "5"
	}))

	fmt.Println("[rune]: ", array.CbKey([]rune{'a', 'b', 'd', 'e', 'f'}, func (v rune) bool {
		return v == 'e'
	}))

	fmt.Println("[complex]: ", array.CbKey([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func (v complex128) bool {
		return v == 5 + 6i
	}))

	fmt.Println("[struct]: ", array.CbKey([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, func (v struct{a int}) bool {
		return v.a == 3
	}))

	fmt.Println("[interface]: ", array.CbKey([]interface{}{1, "2", false}, func (v interface{}) bool {
		return v == false
	}))
}

func arrCbIndex() {
	fmt.Println("--- Find index of element pass condition in callback function ---")
	fmt.Println("[int]: ", array.CbIndex([]int{1, -2, 3, -4, 5, 6}, func (v int) bool {
		return v == 3
	}))

	fmt.Println("[uint]: ", array.CbIndex([]uint{1, 2, 3, 4, 5, 7}, func (v uint) bool {
		return v == 3
	}))

	fmt.Println("[float]: ", array.CbIndex([]float64{1.2, 2.3, 3.4}, func (v float64) bool {
		return v == 3.4
	}))

	fmt.Println("[string]: ", array.CbIndex([]string{"1", "4", "5", "6"}, func (v string) bool {
		return v == "5"
	}))

	fmt.Println("[rune]: ", array.CbIndex([]rune{'a', 'b', 'd', 'e', 'f'}, func (v rune) bool {
		return v == 'e'
	}))

	fmt.Println("[complex]: ", array.CbIndex([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func (v complex128) bool {
		return v == 5 + 6i
	}))

	fmt.Println("[struct]: ", array.CbIndex([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, func (v struct{a int}) bool {
		return v.a == 3
	}))

	fmt.Println("[interface]: ", array.CbIndex([]interface{}{1, "2", false}, func (v interface{}) bool {
		return v == false
	}))
}

func arrCbFilter() {
	fmt.Println("--- Filter elements in array by pass condition in callback function---")
	fmt.Println("[int]: ", array.CbFilter([]int{1, -2, 3, -4, 5, 6}, func (v int) bool {
		return v > 2
	}))

	fmt.Println("[uint]: ", array.CbFilter([]uint{1, 2, 3, 4, 5, 7}, func (v uint) bool {
		return v > 2
	}))

	fmt.Println("[float]: ", array.CbFilter([]float64{1.2, 2.3, 3.4}, func (v float64) bool {
		return v > 2
	}))

	fmt.Println("[string]: ", array.CbFilter([]string{"1", "4", "5", "6"}, func (v string) bool {
		return v > "3"
	}))

	fmt.Println("[rune]: ", array.CbFilter([]rune{'a', 'b', 'd', 'e', 'f'}, func (v rune) bool {
		return v > 'c'
	}))

	fmt.Println("[complex]: ", array.CbFilter([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func (v complex128) bool {
		return real(v) > 3
	}))

	fmt.Println("[struct]: ", array.CbFilter([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, func (v struct{a int}) bool {
		return v.a > 0
	}))

	fmt.Println("[interface]: ", array.CbFilter([]interface{}{1, "2", false}, func (v interface{}) bool {
		return v == false
	}))
}

func arrCbFind() {
	fmt.Println("--- Find element in array by pass condition in callback function---")
	fmt.Println("[int]: ", array.CbFind([]int{1, -2, 3, -4, 5, 6}, func (v int) bool {
		return v == 3
	}))

	fmt.Println("[uint]: ", array.CbFind([]uint{1, 2, 3, 4, 5, 7}, func (v uint) bool {
		return v == 3
	}))

	fmt.Println("[float]: ", array.CbFind([]float64{1.2, 2.3, 3.4}, func (v float64) bool {
		return v == 3.4
	}))

	fmt.Println("[string]: ", array.CbFind([]string{"1", "4", "5", "6"}, func (v string) bool {
		return v == "5"
	}))

	fmt.Println("[rune]: ", array.CbFind([]rune{'a', 'b', 'd', 'e', 'f'}, func (v rune) bool {
		return v == 'e'
	}))

	fmt.Println("[complex]: ", array.CbFind([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func (v complex128) bool {
		return v == 5 + 6i
	}))

	fmt.Println("[struct]: ", array.CbFind([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, func (v struct{a int}) bool {
		return v.a == 3
	}))

	fmt.Println("[interface]: ", array.CbFind([]interface{}{1, "2", false}, func (v interface{}) bool {
		return v == false
	}))
}

func arrCbFor() {
	fmt.Println("--- Loop array then handler with callback function ---")
	fmt.Println("[int]: ")
	array.CbFor([]int{1, -2, 3, -4, 5, 6}, func (v int) {
		fmt.Println(v)
	})

	fmt.Println("[uint]: ")
	array.CbFor([]uint{1, 2, 3, 4, 5, 7}, func (v uint) {
		fmt.Println(v)
	})

	fmt.Println("[float]: ")
	array.CbFor([]float64{1.2, 2.3, 3.4}, func (v float64) {
		fmt.Println(v)
	})

	fmt.Println("[string]: ")
	array.CbFor([]string{"1", "4", "5", "6"}, func (v string) {
		fmt.Println(v)
	})

	fmt.Println("[rune]: ")
	array.CbFor([]rune{'a', 'b', 'd', 'e', 'f'}, func (v rune) {
		fmt.Println(v)
	})

	fmt.Println("[complex]: ")
	array.CbFor([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func (v complex128) {
		fmt.Println(v)
	})

	fmt.Println("[struct]: ")
	array.CbFor([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, func (v struct{a int}) {
		fmt.Println(v)
	})

	fmt.Println("[interface]: ")
	array.CbFor([]interface{}{1, "2", false}, func (v interface{}) {
		fmt.Println(v)
	})
}

func arrCbMap() {
	fmt.Println("--- Map array then handler with callback function ---")
	fmt.Println("[int]: ", array.CbMap([]int{1, -2, 3, -4, 5, 6}, func (v int) int {
		return v * 2
	}))

	fmt.Println("[uint]: ", array.CbMap([]uint{1, 2, 3, 4, 5, 7}, func (v uint) uint {
		return v * 2
	}))

	fmt.Println("[float]: ", array.CbMap([]float64{1.2, 2.3, 3.4}, func (v float64) float64 {
		return v * 2
	}))

	fmt.Println("[string]: ", array.CbMap([]string{"1", "4", "5", "6"}, func (v string) string {
		return v + "1"
	}))

	fmt.Println("[rune]: ", array.CbMap([]rune{'a', 'b', 'd', 'e', 'f'}, func (v rune) rune {
		return v + 1
	}))

	fmt.Println("[complex]: ", array.CbMap([]complex128{1 + 2i, 2 + 3i, 5 + 6i, 6 + 7i}, func (v complex128) complex128 {
		return v * 2
	}))

	fmt.Println("[struct]: ", array.CbMap([]struct{a int}{struct{a int}{-1}, struct{a int}{-2}, struct{a int}{3}, struct{a int}{4}, struct{a int}{5}, struct{a int}{6}}, func (v struct{a int}) struct{a int} {
		return struct{a int}{v.a * 2}
	}))

	fmt.Println("[interface]: ", array.CbMap([]interface{}{1, "2", false}, func (v interface{}) interface{} {
		return v
	}))
}

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

	arrIntersect()
	arrMin()
	arrMax()

	arrCbKey()
	arrCbIndex()
	arrCbFilter()
	arrCbFind()
	arrCbFor()
	arrCbMap()
}