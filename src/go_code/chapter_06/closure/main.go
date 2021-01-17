/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/12 23:17
 */
package main

import (
	"fmt"
	"strings"
)

// 闭包累加器
func AddUpper() func(int) int {
	// 闭包: 闭包是一个函数和与之相关引用的整体(line 11 ~ line 21).
	var n int = 1
	var str string = "hello"
	return func(i int) int {
		n = n + i
		str += string(36)
		fmt.Println("str =", str)
		return n
	}
}

// 需求: 编写一个程序, 要求:
// 1. 编写一个函数 makeSuffix(suffix string), 可以接收一个文件后缀名(.jpg), 并返回一个闭包.
// 2. 调用闭包, 可以传入一个文件名, 如果该文件名没有指定后缀(.jpg), 则返回文件名.jpg, 如已经有了.jpg后缀, 则返回原文件名.
// 3. 要求使用闭包方式完成.
// 4. strings.HasSuffix(...), 该函数可以判断某个字符串是否有指定的后缀.
func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}

func main() {
	addUpper := AddUpper()
	fmt.Println(addUpper(1))  // 2	str = hello$
	fmt.Println(addUpper(10)) // 12 str = hello$$
	fmt.Println(addUpper(20)) // 32 str = hello$$$

	suffixFunc := makeSuffix(".jpg")
	fmt.Println(suffixFunc("1.jpg")) // 1.jpg
	fmt.Println(suffixFunc("2"))     // 2.jpg
	fmt.Println(suffixFunc("3.png")) // 3.png.jpg

}
