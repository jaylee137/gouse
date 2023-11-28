package array

import (
	"testing"
)

func TestMin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	minExpected := 1
	if Min(arr) != minExpected {
		t.Errorf("Expected %d, got %d", minExpected, Min(arr))
	}

	maxExpected := 5
	if Max(arr) != maxExpected {
		t.Errorf("Expected %d, got %d", maxExpected, Max(arr))
	}
}
