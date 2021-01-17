/*
 * This is description.
 * error:
	1. Go语言不支持Java的try...catch...finally的错误处理方式.
	2. defer + recover: 处理异常.
	3. 自定义异常: errors.New + panic 函数.
 * @author Chris0710
 * @date 2021/1/13 23:39
*/
package main

import (
	"errors"
	"fmt"
	"strings"
)

func test() {
	defer func() {
		// recover()是属于内置函数, 可以捕获到异常.
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
			// 处理发生异常的逻辑, 如: 发送邮件.
			fmt.Println("程序发生异常了, 请注意.")
		}
	}()
	num1 := 1
	num2 := 0

	res := num1 / num2 // panic: runtime error: integer divide by zero
	fmt.Println("res =", res)
}

func main() {
	// test()
	// for {
	// 	fmt.Println("main()...")
	// 	time.Sleep(time.Second)
	// }

	test2()
	fmt.Println("main()...")
}

func test2() {
	err := readConf("application.yaml")
	if err != nil {
		// 如果程序出现错误, 将终止程序, panic(err)将输出程序错误信息.
		panic(err) // panic: 文件后缀名错误
	}
	fmt.Println("test2()继续执行.") // 此行代码不会被执行.
}

func readConf(s string) error {
	if !strings.HasSuffix(s, "yml") {
		return errors.New("文件后缀名错误")
	}
	return nil
}
