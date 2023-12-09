package strings

import (
	"testing"
)

func TestIsLetter(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'a', true},
		{'A', true},
		{'z', true},
		{'Z', true},
		{'1', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsLetter(test.input)
		if actual != test.expected {
			t.Errorf("IsLetter(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'1', true},
		{'9', true},
		{'a', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsDigit(test.input)
		if actual != test.expected {
			t.Errorf("IsDigit(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"test", "es", true},
		{"test", "test", true},
		{"test", "st", true},
		{"test", "t", true},
		{"test", "e", true},
		{"test", "s", true},
		{"test", "z", false},
		{"test", "tt", false},
		{"test", "ttt", false},
		{"test", "tttt", false},
		{"test", "ttttt", false},
		{"test", "tttttt", false},
	}

	for _, test := range tests {
		actual := Includes(test.input, test.substr)
		if actual != test.expected {
			t.Errorf("Includes(%q, %q): expected %t, actual %t", test.input, test.substr, test.expected, actual)
		}
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"test", "es", false},
		{"test", "test", true},
		{"test", "st", false},
		{"test", "t", true},
		{"test", "e", false},
		{"test", "s", false},
		{"test", "z", false},
		{"test", "tt", false},
		{"test", "ttt", false},
		{"test", "tttt", false},
		{"test", "ttttt", false},
		{"test", "tttttt", false},
	}

	for _, test := range tests {
		actual := StartsWith(test.input, test.substr)
		if actual != test.expected {
			t.Errorf("StartsWith(%q, %q): expected %t, actual %t", test.input, test.substr, test.expected, actual)
		}
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"test", "es", false},
		{"test", "test", true},
		{"test", "st", true},
		{"test", "t", true},
		{"test", "e", false},
		{"test", "s", false},
		{"test", "z", false},
		{"test", "tt", false},
		{"test", "ttt", false},
		{"test", "tttt", false},
		{"test", "ttttt", false},
		{"test", "tttttt", false},
	}

	for _, test := range tests {
		actual := EndsWith(test.input, test.substr)
		if actual != test.expected {
			t.Errorf("EndsWith(%q, %q): expected %t, actual %t", test.input, test.substr, test.expected, actual)
		}
	}
}

func TestIsLower(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'a', true},
		{'A', false},
		{'z', true},
		{'Z', false},
		{'1', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsLower(test.input)
		if actual != test.expected {
			t.Errorf("IsLower(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsUpper(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'a', false},
		{'A', true},
		{'z', false},
		{'Z', true},
		{'1', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsUpper(test.input)
		if actual != test.expected {
			t.Errorf("IsUpper(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}
