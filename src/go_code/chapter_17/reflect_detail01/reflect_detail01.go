/*
 * This is description.
 * 反射的注意事项和细节:
 * @author Chris0710
 * @date 2021/3/8 10:21
 */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. reflect.Value.Kind: 获取变量的类别, 返回的是一个常量Const(Bool, Int, Struct...).

	// 2. Type和Kind的区别: Type: 类型; Kind: 类别. 二者可能相同, 可能不同.
	// var num int = 1		Type: Int, Kind: Int. 二者相同.
	// var stu Student		Type: pkg.Student, Kind: Struct. 二者不同.

	// 3. 通过反射, 可以让变量在interface{}, reflect.Value之间转化, 详见: reflect_demo01.go
	// 变量 <-> interface{} <-> reflect.Value.

	// 4. 使用反射来获取变量的值, 类型需要匹配. 比如: x为int, 则需要reflect.Value(x).Int(). 而不能使用其他, 否则抛出panic.

	// 5. 通过反射来修改变量的值, 当使用SetXxx()来修改变量的值时, 需要使用到reflect.Value.Elem()来达到通过修改对应的指针进行修改变量的值.
	var num int = 1
	testInt(&num)
	fmt.Printf("num: %v\n", num) // num: 2

	// 6. 如何理解reflect.Value.Elem()?
	var n int = 100
	value := reflect.ValueOf(&n) // var value *int = &n
	value.Elem().SetInt(200)     // *value = 200
	fmt.Printf("n修改后的值: %v\n", n)
}

func testInt(i interface{}) {
	iType := reflect.TypeOf(i)
	fmt.Printf("iType: %v\n", iType)
	// reflect.ValueOf(i).SetInt(2) // panic: reflect: reflect.flag.mustBeAssignable using unaddressable value.
	reflect.ValueOf(i).Elem().SetInt(2)
}
