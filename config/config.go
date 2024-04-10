/**
 * @ Author: ClearDewy
 * @ Desc:
 **/
package config

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func camelCaseToEnvVar(name string) string {
	// 正则表达式匹配大写字母
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	words := re.FindAllString(name, -1)
	for i := range words {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, "_")
}
func setFieldWithValue(field reflect.Value, value string) error {
	if field.CanSet() {
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Int, reflect.Uint:
			// 首先检查字段类型是否为 time.Duration
			if field.Type() == reflect.TypeOf(time.Duration(0)) {
				if duration, err := time.ParseDuration(value); err == nil {
					field.Set(reflect.ValueOf(duration))
				} else {
					return err
				}
			} else {
				// 处理普通的整数
				if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
					field.SetInt(intValue)
				} else {
					return err
				}
			}
		case reflect.Bool:
			if boolValue, err := strconv.ParseBool(value); err == nil {
				field.SetBool(boolValue)
			} else {
				return err
			}

		default:
			field.Set(reflect.ValueOf(value))
		}
	}
	return nil
}
