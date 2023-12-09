package strings

import (
	"testing"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test capitalize", "Test Capitalize"},
		{"test capitalize", "Test Capitalize"},
		{"test capitalize", "Test Capitalize"},
		{"capitalize123", "Capitalize123"},
	}

	for _, test := range tests {
		actual := Capitalize(test.input)
		if actual != test.expected {
			t.Errorf("Capitalize(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test snake", "test_snake"},
		{"test-snake", "test_snake"},
		{"testSnake", "test_snake"},
		{"snake123", "snake_123"},
	}

	for _, test := range tests {
		actual := SnakeCase(test.input)
		if actual != test.expected {
			t.Errorf("SnakeCase(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test kebab", "test-kebab"},
		{"test_kebab", "test-kebab"},
		{"testKebab", "test-kebab"},
		{"kebab123", "kebab-123"},
	}

	for _, test := range tests {
		actual := KebabCase(test.input)
		if actual != test.expected {
			t.Errorf("KebabCase(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test camel", "testCamel"},
		{"test-camel", "testCamel"},
		{"test_camel", "testCamel"},
		{"camel123", "camel123"},
	}

	for _, test := range tests {
		actual := CamelCase(test.input)
		if actual != test.expected {
			t.Errorf("CamelCase(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

// go test -v ./strings/...
