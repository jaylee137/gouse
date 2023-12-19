package structs

import (
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    interface{}
		expectedResult interface{}
	}{
		{
			name:           "Clone Person struct",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			expectedResult: Person{Name: "John", Age: 25, Address: "123 Main St"},
		},
		{
			name:           "Clone with different values",
			inputStruct:    Person{Name: "Alice", Age: 30, Address: "456 Side St"},
			expectedResult: Person{Name: "Alice", Age: 30, Address: "456 Side St"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Clone(test.inputStruct)
			expectedResult := test.expectedResult

			if !reflect.DeepEqual(result, expectedResult) {
				t.Errorf("Expected %+v, but got %+v", expectedResult, result)
			}
		})
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    interface{}
		fieldName      string
		expectedResult bool
	}{
		{
			name:           "Has Name field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			expectedResult: true,
		},
		{
			name:           "Does not have Email field",
			inputStruct:    Person{Name: "Alice", Age: 30, Address: "456 Side St"},
			fieldName:      "Email",
			expectedResult: false,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Has(test.inputStruct, test.fieldName)
			expectedResult := test.expectedResult

			if result != expectedResult {
				t.Errorf("Expected %v, but got %v", expectedResult, result)
			}
		})
	}
}

func TestHasEmpty(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    interface{}
		fieldName      string
		expectedResult bool
	}{
		{
			name:           "Name field is not empty",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			expectedResult: false,
		},
		{
			name:           "Email field is empty",
			inputStruct:    Person{Name: "Alice", Age: 30, Address: "456 Side St"},
			fieldName:      "Email",
			expectedResult: true,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HasEmpty(test.inputStruct, test.fieldName)
			expectedResult := test.expectedResult

			if result != expectedResult {
				t.Errorf("Expected %v, but got %v", expectedResult, result)
			}
		})
	}
}
