package array

import "testing"

func TestExist(t *testing.T) {
	if !Includes[int]([]int{1, 2, 3}, 1) {
		t.Errorf("Exist[int]([]int{1, 2, 3}, 1) = false, want true")
	}
}

func TestEqual(t *testing.T) {
	if !Includes[int]([]int{1, 2, 3}, 1) {
		t.Errorf("Equal[int](1, 1) = false, want true")
	}
}