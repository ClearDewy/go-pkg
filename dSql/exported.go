/**
 * @ Author: ClearDewy
 * @ Desc:
 **/
package dSql

import "fmt"

type String string
type Bool bool
type Int int

func (s *String) Scan(value interface{}) error {
	if value == nil {
		*s = ""
	} else if bv, ok := value.([]byte); ok {
		// 如果value是[]byte类型，则将其转换为string
		*s = String(bv)
	} else {
		// 您可以根据需要在这里添加更多的类型检查
		return fmt.Errorf("cannot scan type %T into String", value)
	}
	return nil
}
func (b *Bool) Scan(value interface{}) error {
	if value == nil {
		*b = false
	} else if iv, ok := value.(int64); ok {
		// 如果value是int64类型，则将其转换为Bool
		*b = iv != 0
	} else {
		return fmt.Errorf("cannot scan type %T into Bool", value)
	}
	return nil
}
func (i *Int) Scan(value interface{}) error {
	if value == nil {
		*i = 0
	} else if iv, ok := value.(int64); ok {
		// 如果value是int64类型，则将其转换为Int
		*i = Int(iv)
	} else {
		return fmt.Errorf("cannot scan type %T into Int", value)
	}
	return nil
}
