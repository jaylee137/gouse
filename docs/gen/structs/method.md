# Method


### Merge
```go
import (
	"fmt"

	"reflect"
)
```

```go
func Merge(structs ...interface{}) interface{} {
	if len(structs) == 0 {
		return nil
	}

	fieldValues := make(map[string]interface{})

	// Iterate through the input structs
	for _, structInstance := range structs {
		structValue := reflect.ValueOf(structInstance)

		// Iterate through the fields of the struct
		for i := 0; i < structValue.NumField(); i++ {
			field := structValue.Type().Field(i)
			fieldName := field.Name

			// Store field value in the map
			fieldValues[fieldName] = structValue.Field(i).Interface()
		}
	}

	// Create a slice of reflect.StructField for the new struct
	var newFields []reflect.StructField

	// Iterate through the fields of the new struct
	for fieldName, value := range fieldValues {
		// Add fields to the new struct dynamically
		newFields = append(newFields, reflect.StructField{
			Name: fieldName,
			Type: reflect.TypeOf(value),
		})
	}

	// Create a new struct type
	newStructType := reflect.StructOf(newFields)

	// Create a new instance of the struct
	newStructValue := reflect.New(newStructType).Elem()

	// Set field values in the new struct
	for _, field := range newFields {
		fieldName := field.Name
		value := fieldValues[fieldName]
		newStructValue.FieldByName(fieldName).Set(reflect.ValueOf(value))
	}

	return newStructValue.Interface()
}
```

### Remove
```go
import (
	"fmt"

	"reflect"
)
```

```go
func Remove(structInstance interface{}, fields ...string) interface{} {
	structValue := reflect.ValueOf(structInstance)
	structType := structValue.Type()

	// Create a map to store excluded field names
	excludedFields := make(map[string]bool)
	for _, field := range fields {
		excludedFields[field] = true
	}

	// Create a new list of struct fields excluding the specified fields
	var newFields []reflect.StructField
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structType.Field(i).Name
		if !excludedFields[fieldName] {
			newFields = append(newFields, structType.Field(i))
		}
	}

	newStructType := reflect.StructOf(newFields)
	newStructValue := reflect.New(newStructType).Elem()

	// Copy values from the original struct to the new struct
	for _, field := range newFields {
		fieldName := field.Name
		value := structValue.FieldByName(fieldName).Interface()
		newStructValue.FieldByName(fieldName).Set(reflect.ValueOf(value))
	}

	return newStructValue.Interface()
}
```

### Add
```go
import (
	"fmt"

	"reflect"
)
```

```go
func Add(structInstance interface{}, fields map[string]interface{}) interface{} {
	structType := reflect.TypeOf(structInstance)

	// Create a map to store existing field values
	existingValues := make(map[string]interface{})
	for i := 0; i < structType.NumField(); i++ {
		fieldName := structType.Field(i).Name
		existingValues[fieldName] = reflect.ValueOf(structInstance).Field(i).Interface()
	}

	// Add new fields to the map
	for fieldName, value := range fields {
		existingValues[fieldName] = value
	}

	// Create a new struct type with the combined fields
	newStructType := reflect.StructOf(
		func() []reflect.StructField {
			var fields []reflect.StructField
			for fieldName, value := range existingValues {
				fields = append(fields, reflect.StructField{
					Name: fieldName,
					Type: reflect.TypeOf(value),
				})
			}
			return fields
		}(),
	)

	// Create a new struct instance of the new type
	newStructValue := reflect.New(newStructType).Elem()

	// Set field values for the new instance
	for i := 0; i < newStructType.NumField(); i++ {
		fieldName := newStructType.Field(i).Name
		fieldValue := existingValues[fieldName]
		newStructValue.Field(i).Set(reflect.ValueOf(fieldValue))
	}

	return newStructValue.Interface()
}
```

### Set
```go
import (
	"fmt"

	"reflect"
)
```

```go
func Set(structInstance interface{}, fieldName string, value interface{}) {
	// Get the reflect.Value of the struct instance
	structValue := reflect.ValueOf(structInstance)

	// Check if the structInstance is a pointer
	if structValue.Kind() != reflect.Ptr || structValue.IsNil() {
		fmt.Println("Struct instance must be a non-nil pointer.")
		return
	}

	// Dereference the pointer to get the actual struct value
	structValue = structValue.Elem()

	// Get the field by name
	field := structValue.FieldByName(fieldName)

	// Check if the field is valid and exported
	if !field.IsValid() || !field.CanSet() {
		fmt.Printf("Field %s is unexported or not found.\n", fieldName)
		return
	}

	// Check if the types match before setting the value
	if reflect.ValueOf(value).Type().AssignableTo(field.Type()) {
		field.Set(reflect.ValueOf(value))
	} else {
		fmt.Printf("Field %s type mismatch: expected %s, got %s\n", fieldName, field.Type(), reflect.ValueOf(value).Type())
	}
}
```

### Get
```go
import (
	"fmt"

	"reflect"
)
```

```go
func Get(structInstance interface{}, fieldName string) interface{} {
	structValue := reflect.ValueOf(structInstance)

	// Get the field value
	fieldValue := structValue.FieldByName(fieldName)

	return fieldValue.Interface()
}
```
