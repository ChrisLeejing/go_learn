/*
 * This is description.
 * 内置函数: https://golang.org/pkg/builtin/
 * @author Chris0710
 * @date 2021/1/13 23:12
 */
package main

import "fmt"

func main() {
	// 1. func len(v Type) int: 用于求长度, 比如: string, slice, channel_demo, array, map.
	str1 := "hello"
	fmt.Println("str1 的长度: ", len(str1))

	// 2. func new(Type) *Type: 用于分配内存, 主要用于值类型, 比如: int, float32, struct...返回的是指针.
	num1 := 100
	fmt.Printf("num1的类型: %T, 值: %v, 地址: %v\n", num1, num1, &num1) // num1的类型: int, 值: 100, 地址: 0xc0000140b8

	num2 := new(int)
	*num2 = 200
	fmt.Printf("num2的类型: %T, 值: %v, 地址: %v, num2指针所指向的值: %v\n", num2, num2, &num2, *num2) // num2的类型: *int, 值: 0xc0000140f0, 地址: 0xc000006030, num2指针所指向的值: 200

	// 3. func make(t Type, size ...IntegerType) Type: 用于分配内存, 主要用于引用类型, 比如: channel_demo, map, slice.
}
