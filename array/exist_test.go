package array

import "testing"

func TestExist(t *testing.T) {
	if !IsExist[int]([]int{1, 2, 3}, 1) {
		t.Errorf("Exist[int]([]int{1, 2, 3}, 1) = false, want true")
	}
}
