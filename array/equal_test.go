package array

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !Equal[int](1, 1) {
		t.Errorf("Equal[int](1, 1) = false, want true")
	}
}