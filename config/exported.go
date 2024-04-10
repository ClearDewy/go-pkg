/**
 * @ Author: ClearDewy
 * @ Desc:
 **/
package config

import (
	"os"
	"reflect"
)

func LoadEnvDefault(v interface{}) error {
	elem := reflect.ValueOf(v).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		typeField := elem.Type().Field(i)
		var err error

		envName := typeField.Tag.Get("env")
		if envName == "" {
			envName = camelCaseToEnvVar(typeField.Name)
		}
		if envVal, ok := os.LookupEnv(envName); field.CanSet() {
			if ok {
				err = setFieldWithValue(field, envVal)
			} else {
				if tagValue := typeField.Tag.Get("default"); tagValue != "" {
					err = setFieldWithValue(field, tagValue)
				}
			}
		} else if err != nil {
			return err
		}
	}
	return nil
}
