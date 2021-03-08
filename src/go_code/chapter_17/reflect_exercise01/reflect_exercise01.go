/*
 * This is description.
 * 需求: 给一个变量var v float64 = 1.2:
	1. 请使用反射获取到reflect.Value, 然后获取对应的Type, Kind, 值.
	2. 并将reflect.Value转换成interface{}, interface{}转换成float64.
 * @author Chris0710
 * @date 2021/3/8 16:17
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v float64 = 1.2
	value := reflect.ValueOf(v)
	fmt.Printf("value: %v\n", value)

	vType := reflect.TypeOf(v)
	fmt.Printf("vType: %v\n", vType)
	vKind := reflect.ValueOf(v).Kind()
	fmt.Printf("vKind: %v\n", vKind)

	i := value.Interface()
	iValue := i.(float64)
	fmt.Printf("iValue: %v\n", iValue)
}
