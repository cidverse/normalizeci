package envstruct

import (
	"fmt"
	"reflect"
	"strings"
)

func EnvMapToStruct(data interface{}, envMap map[string]string) error {
	return envMapToStruct(data, envMap, "")
}

func envMapToStruct(data interface{}, envMap map[string]string, prefix string) error {
	reflectValue := reflect.ValueOf(data)
	reflectType := reflect.TypeOf(data)

	if reflectValue.Kind() != reflect.Ptr || reflectValue.IsNil() {
		return fmt.Errorf("data must be a non-nil pointer to a struct")
	}

	reflectValue = reflectValue.Elem()
	reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fieldValue := reflectValue.Field(i)
		fieldType := reflectType.Field(i)

		if fieldType.Type.Kind() == reflect.Map {
			envPrefixTag := fieldType.Tag.Get("env-prefix")
			subMap := make(map[string]string)

			for k, v := range envMap {
				if strings.HasPrefix(k, envPrefixTag) {
					subMap[strings.TrimPrefix(k, envPrefixTag)] = v
				}
			}

			fieldValue.Set(reflect.ValueOf(subMap))
		} else if fieldType.Type.Kind() == reflect.Struct {
			err := envMapToStruct(fieldValue.Addr().Interface(), envMap, prefix)
			if err != nil {
				return err
			}
		} else {
			// Get the env tag value and map it to the field value
			envTag := fieldType.Tag.Get("env")
			if envTag != "" {
				key := prefix + envTag
				value, ok := envMap[key]
				if !ok {
					continue
				}

				fieldValue.SetString(value)
			}
		}
	}

	return nil
}
