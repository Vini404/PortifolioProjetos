package utils

import "reflect"

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
