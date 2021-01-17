/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/7 9:28
 */
package main

import "fmt"

func main() {
	// 自增i++的注意事项:
	// 1. i++只能独立使用, 不能进行赋值运算, 不能进行逻辑判断.
	// 2. 格式上只有i++, 而没有++i, 这点与Java不同.
	var i = 10
	i++ // ok
	// j := i++ // 编译报错.
	// fmt.Println(j)
	// if i++ > 0 {
	// 	// 编译报错
	// }

	// ++i // 编译报错
	// --i // 编译报错

	// 位运算符:
	// 1. &: "位与"运算符, 规则: 同为1, 则为1, 否则为0.
	// 2. |: "位或"运算符, 规则: 任意一个为1, 则为1, 否则为0.
	// 3. ^: "异或"运算符, 规则: 当二进位不同时, 则为1, 否则为0.
	// 4. <<: "左移"运算符, 规则: 将"<<"左边的运算数的各二进位全部左移若干位, 高位丢弃, 低位补0. 左移n位, 等价于乘以2的n次方.
	// 5. >>: "右移"运算符, 规则: 将">>"左边的运算数的各二进位全部右移若干位, 右移n位, 等价于除以2的n次方.(规则待完善!!!)
	// 指针运算符:
	// 1. &: 返回变量的存储地址. 如: &a: 返回变量的实际地址.
	// 2. *: 指针变量. 如: *a: a的指针变量.
	var num int = 200
	fmt.Println("num的地址:", &num) // num的地址: 0xc0000140b0

	var prt *int = &num
	fmt.Println("prt指向的值:", *prt) // prt指向的值: 200

	fmt.Println("-----------------")

	// func Scanln(...), func Scanf(...).
	// 需求: 从控制台接收用户信息[姓名, 年龄, 薪水, 是否通过考试.]
	// func Scanln(...)
	var name string
	var age byte
	var salary float32
	var isPass bool
	// fmt.Println("请输入姓名: ")
	// fmt.Scanln(&name) // 当程序执行到fmt.Scanln(&name)此行时, 程序会暂停, 等待用户输入, 并且按下Enter回车键, 以继续执行程序.
	// fmt.Println("请输入年龄: ")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水: ")
	// fmt.Scanln(&salary)
	// fmt.Println("请输入是否通过考试: ")
	// fmt.Scanln(&isPass)
	//
	// fmt.Printf("姓名: %v, 年龄: %v, 薪水: %v, 是否通过考试: %v\n", name, age, salary, isPass)

	// func Scanf(...)
	fmt.Println("请输入: 姓名, 年龄, 薪水, 是否通过考试, 注意: 以空格隔开.")
	fmt.Scanf("%v %v %v %v", &name, &age, &salary, &isPass)
	fmt.Printf("姓名: %v, 年龄: %v, 薪水: %v, 是否通过考试: %v\n", name, age, salary, isPass) // 姓名: jack, 年龄: 22, 薪水: 10, 是否通过考试: false

}
