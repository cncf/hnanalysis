package test

import (
	"fmt"
)

// Desc - returns string "variable_type:variable_value"
func Desc(v interface{}) string {
	return fmt.Sprintf("%T:%v", v, v)
}
