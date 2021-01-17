/*
 * This is description.
 * parameter transmission: 函数的两种参数传递方式:
	1. 值传递: 值参数类型默认传递方式是值传递.
	2. 引用传递: 引用类型参数默认传递方式是引用传递.
	3. 值传递和引用传递其实都是传递参数的拷贝, 不同的是: 值传递是进行值的拷贝, 而引用传递是进行地址的拷贝, 通常情况下, 地址拷贝的效率要高于值拷贝.
	4. 值类型: 基本数据类型int系列, float系列, bool, string, 数组, 结构体struct.
	5. 引用类型: 指针, slice切片, map, channel管道, interface等.
 * @author Chris0710
 * @date 2021/1/13 10:14
*/
package main

import "fmt"

var num int = 10

// Name := "chris" //编译报错: Name := "chris" <=> var Name string; Name = "chris", 赋值语句不能出现在函数体外.

func main() {
	fmt.Println("num =", num)
}
