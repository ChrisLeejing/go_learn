/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/6 14:15
 */
package main

import (
	"fmt"
	// "go_code/chapter_03/param_01/book_demo"
	book "go_code/chapter_03/param_01/book_demo" // 包名过长, 支持给包名取别名, 但是原包名则失效, 需要使用别名来访问该包中的函数和变量.
)

func main() {
	// 	获取变量的地址
	var i1 int = 10
	fmt.Println("i1 =", i1)     // i1 = 10
	fmt.Println("i1的地址值:", &i1) // i1的地址值: 0xc0000140b0

	// 1. i1Pstr是一个指针变量. 2. i1Pstr的类型是*int. 3. i1Pstr本身的值是&i1. 4. 获取i1Pstr指针所指向的具体指: *i1Pstr.
	var i1Pstr *int = &i1
	fmt.Println("i1的指针地址:", i1Pstr)       // i1的指针地址: 0xc0000140b0
	fmt.Println("i1Pstr指向的具体值:", *i1Pstr) // i1Pstr指向的具体值: 10

	var i2 int = 12
	fmt.Println("i2的地址值:", &i2)
	var i2Pstr *int = &i2
	fmt.Println("i2Pstr指针修改前, 所指向的具体指:", *i2Pstr) // i2Pstr指针修改前, 所指向的具体指: 12
	*i2Pstr = 11
	fmt.Println("i2Pstr指针修改后, 所指向的具体指:", *i2Pstr) // i2Pstr指针修改后, 所指向的具体指: 11

	fmt.Println("---------------")

	// pointer exercise
	var a int = 300
	var b int = 400
	var prt *int = &a                               // *prt -> 300
	*prt = 100                                      // *prt -> 100, a = 100
	prt = &b                                        // *prt -> 400
	*prt = 200                                      // *prt -> 200, b = 200
	fmt.Printf("a=%v, b=%v, *prt=%v\n", a, b, *prt) // a=100, b=200, *prt=200

	/*
		1. 值类型: 值类型都有对应的指针类型, int -> *int, float32 -> *float32, 等等.
		2. 值类型: 包括基本数据类型int系列, float系列, bool, string, 数组, 结构体struct.
		3. 引用类型: pointer指针, slice切片, chan管道, map, interface, 等等.
	*/

	// 输出book_demo package中的变量Book1.
	// fmt.Println("输出book_demo package中的变量Book1:", book_demo.Book1)
	fmt.Println("输出book_demo package中的变量Book1:", book.Book1)
	// sum, sub := book_demo.SumAndSub(10, 20)
	sum, sub := book.SumAndSub(10, 20)
	fmt.Printf("sum = %v, sub = %v", sum, sub)
}
