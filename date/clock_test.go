package date

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "Single digit hour, minute, and second",
			input:    time.Date(2024, 2, 4, 5, 6, 7, 0, time.UTC),
			expected: "05:06:07",
		},
		{
			name:     "Double digit hour, minute, and second",
			input:    time.Date(2024, 2, 4, 15, 16, 17, 0, time.UTC),
			expected: "15:16:17",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := formatTime(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
