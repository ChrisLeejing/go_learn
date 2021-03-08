/*
 * This is description.
 * @author Chris0710
 * @date 2021/3/8 16:17
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. 看代码, 查看是否正确, 为什么?
	// var str string = "tom"
	// fs := reflect.ValueOf(str)
	// fs.SetString("jack") // panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	// fmt.Printf("str: %v\n", str)

	// 2. 修正代码如下:
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("str: %v\n", str) // str: jack
}
