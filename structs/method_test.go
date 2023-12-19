package structs

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

type Job struct {
	Title string
}

type Developer struct {
	Name    string
	Age     int
	Address string
	Title   string
}

func TestMergeStructs(t *testing.T) {
	tests := []struct {
		name   string
		input  []interface{}
		output interface{}
	}{
		{
			name:   "Merge Person and Job",
			input:  []interface{}{Person{Name: "John", Age: 25, Address: "123 Main St"}, Job{Title: "Developer"}},
			output: map[string]interface{}{"Name": "John", "Age": 25, "Address": "123 Main St", "Title": "Developer"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Merge(test.input...)
			fmt.Println(result)
			fmt.Println(result.(map[string]interface{})["Name"])
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected %+v, but got %+v", test.output, result)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    interface{}
		fieldsToRemove []string
		expectedResult interface{}
	}{
		{
			name:           "Remove Address field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldsToRemove: []string{"Address"},
			expectedResult: struct {
				Name string
				Age  int
			}{Name: "John", Age: 25},
		},
		{
			name:           "Remove multiple fields",
			inputStruct:    Person{Name: "Jane", Age: 30, Address: "456 Side St"},
			fieldsToRemove: []string{"Age", "Address"},
			expectedResult: struct{ Name string }{Name: "Jane"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Remove(test.inputStruct, test.fieldsToRemove...)
			expectedResult := reflect.ValueOf(test.expectedResult).Interface()

			if !reflect.DeepEqual(result, expectedResult) {
				t.Errorf("Expected %+v, but got %+v", expectedResult, result)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    interface{}
		newFields      map[string]interface{}
		expectedResult interface{}
	}{
		{
			name:           "Add Email field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			newFields:      map[string]interface{}{"Email": "john@example.com"},
			expectedResult: map[string]interface{}{"Name": "John", "Age": 25, "Address": "123 Main St", "Email": "john@example.com"},
		},
		{
			name:           "Add multiple fields",
			inputStruct:    Person{Name: "Jane", Age: 30, Address: "456 Side St"},
			newFields:      map[string]interface{}{"Email": "jane@example.com", "Phone": "555-1234"},
			expectedResult: map[string]interface{}{"Name": "Jane", "Age": 30, "Address": "456 Side St", "Email": "jane@example.com", "Phone": "555-1234"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Add(test.inputStruct, test.newFields)

			// Convert result and expected result to maps for comparison
			resultMap := make(map[string]interface{})
			for _, field := range reflect.VisibleFields(reflect.TypeOf(result)) {
				resultMap[field.Name] = reflect.ValueOf(result).FieldByName(field.Name).Interface()
			}

			expectedResultMap := test.expectedResult.(map[string]interface{})

			if !reflect.DeepEqual(resultMap, expectedResultMap) {
				t.Errorf("Expected %+v, but got %+v", expectedResultMap, resultMap)
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    interface{}
		fieldName      string
		value          interface{}
		expectedResult interface{}
	}{
		{
			name:           "Set Name field",
			inputStruct:    &Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			value:          "Jane",
			expectedResult: &Person{Name: "Jane", Age: 25, Address: "123 Main St"},
		},
		{
			name:           "Set Age field",
			inputStruct:    &Person{Name: "Bob", Age: 30, Address: "456 Side St"},
			fieldName:      "Age",
			value:          35,
			expectedResult: &Person{Name: "Bob", Age: 35, Address: "456 Side St"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Set(test.inputStruct, test.fieldName, test.value)
			result := test.inputStruct

			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Errorf("Expected %+v, but got %+v", test.expectedResult, result)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name           string
		inputStruct    interface{}
		fieldName      string
		expectedResult interface{}
	}{
		{
			name:           "Get Name field",
			inputStruct:    Person{Name: "John", Age: 25, Address: "123 Main St"},
			fieldName:      "Name",
			expectedResult: "John",
		},
		{
			name:           "Get Age field",
			inputStruct:    Person{Name: "Bob", Age: 30, Address: "456 Side St"},
			fieldName:      "Age",
			expectedResult: 30,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Get(test.inputStruct, test.fieldName)

			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Errorf("Expected %+v, but got %+v", test.expectedResult, result)
			}
		})
	}
}
