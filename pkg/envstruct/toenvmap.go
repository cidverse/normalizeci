package envstruct

import (
	"reflect"
	"strings"
)

func StructToEnvMap(data interface{}) map[string]string {
	return structToEnvMap(data, "")
}

func structToEnvMap(data interface{}, prefix string) map[string]string {
	result := make(map[string]string)

	reflectValue := reflect.ValueOf(data)
	reflectType := reflect.TypeOf(data)

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Field(i)
		fieldType := reflectType.Field(i)

		if fieldType.Type.Kind() == reflect.Map {
			envPrefixTag := fieldType.Tag.Get("env-prefix")
			subMap := field.Interface().(map[string]string)

			for k, v := range subMap {
				result[envPrefixTag+strings.ToUpper(k)] = v
			}
		} else if fieldType.Type.Kind() == reflect.Struct {
			subMap := structToEnvMap(field.Interface(), "")
			for k, v := range subMap {
				result[prefix+k] = v
			}
		} else {
			// Get the env tag value and map it to the field value
			envTag := fieldType.Tag.Get("env")
			if envTag != "" {
				result[prefix+envTag] = field.String()
			}
		}
	}

	return result
}
