package number

import (
	"testing"
)

func TestRandom(t *testing.T) {
	min := 1
	max := 10

	result := Random(min, max)

	if result < min || result > max {
		t.Errorf("Random() = %d; want >= %d and <= %d", result, min, max)
	}
}

func BenchmarkRandom(b *testing.B) {
	min := 1
	max := 10

	for i := 0; i < b.N; i++ {
		Random(min, max)
	}
}

func TestClamp(t *testing.T) {
	min := 1
	max := 10

	result := Clamp(0, min, max)

	if result != min {
		t.Errorf("Clamp() = %d; want %d", result, min)
	}

	result = Clamp(11, min, max)

	if result != max {
		t.Errorf("Clamp() = %d; want %d", result, max)
	}

	result = Clamp(5, min, max)

	if result != 5 {
		t.Errorf("Clamp() = %d; want 5", result)
	}
}

func BenchmarkClamp(b *testing.B) {
	min := 1
	max := 10

	for i := 0; i < b.N; i++ {
		Clamp(5, min, max)
	}
}

func TestInRange(t *testing.T) {
	min := 1
	max := 10

	result := InRange(0, min, max)

	if result {
		t.Errorf("InRange() = %t; want false", result)
	}

	result = InRange(11, min, max)

	if result {
		t.Errorf("InRange() = %t; want false", result)
	}

	result = InRange(5, min, max)

	if !result {
		t.Errorf("InRange() = %t; want true", result)
	}
}

func BenchmarkInRange(b *testing.B) {
	min := 1
	max := 10

	for i := 0; i < b.N; i++ {
		InRange(5, min, max)
	}
}
