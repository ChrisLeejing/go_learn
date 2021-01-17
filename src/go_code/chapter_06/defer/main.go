/*
 * This is description.
 * defer: 我们在编写程序过程中, 经常会创建资源(数据库连接, 文件句柄, 锁等), 为了使程序执行完成后, 及时的释放资源,Go语言提供了defer(延时机制).
		  当程序执行到defer时, defer语句暂时不会被执行, 而是被压入到defer栈中, 当执行完毕后, 再以先进后出的规则从defer栈中取出, 并执行.
 * @author Chris0710
 * @date 2021/1/13 0:02
*/
package main

import (
	"database/sql"
	"fmt"
	"os"
)

func sum(n1 int, n2 int) int {
	defer fmt.Println("n1 =", n1)
	defer fmt.Println("n2 =", n2)

	// 在defer将语句放入栈中时, 同时也会将相应的值拷贝入栈.
	n1++
	n2++

	res := n1 + n2
	fmt.Println("res = ", res)
	return res
}

// 注: 在defer语句后, 程序可以继续使用已创建的资源, 当函数执行完毕后, 系统会自动从defer栈中取出被被压入的语句进行执行, 关闭相应的资源.
// defer的最佳实践1: 关闭文件资源
func test1() {
	f, _ := os.Open("test.txt")
	defer f.Close()
	// 其他逻辑代码
}

// defer的最佳实践2: 释放数据库连接
func test2() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 其他逻辑代码
}

func main() {
	sum := sum(1, 2)
	fmt.Println("main sum =", sum)
	/*
		res =  3
		n2 = 2
		n1 = 1
		main sum = 3
		---------------------
		res =  5
		n2 = 2
		n1 = 1
		main sum = 5
	*/
}
