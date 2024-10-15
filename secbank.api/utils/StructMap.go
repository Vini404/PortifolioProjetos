package utils

import (
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	// Result map to hold the struct fields and values
	result := make(map[string]interface{})

	// Get the value of the struct
	value := reflect.ValueOf(obj)

	// Get the type of the struct
	structType := reflect.TypeOf(obj)

	// Iterate over the struct fields
	for i := 0; i < value.NumField(); i++ {
		// Get the field and its corresponding type
		field := structType.Field(i)
		fieldValue := value.Field(i).Interface()

		// Get the db tag value
		dbTag := field.Tag.Get("db")

		// Use the db tag as the key in the result map
		if dbTag != "" {
			result[dbTag] = fieldValue
		}
	}

	return result
}

func StructToMapWithoutID(entity interface{}, idField string) (map[string]interface{}, error) {
	entityValue := reflect.ValueOf(entity)
	entityType := reflect.TypeOf(entity)

	// Ensure the input is a struct
	if entityType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %s", entityType.Kind().String())
	}

	result := make(map[string]interface{})

	// Iterate over the struct fields
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldValue := entityValue.Field(i)

		// Skip the ID field (if it exists and is set to 0)
		if field.Tag.Get("db") == idField && fieldValue.Int() == 0 {
			continue
		}

		// Add the field to the result map if it's not ID
		result[field.Tag.Get("db")] = fieldValue.Interface()
	}

	return result, nil
}
