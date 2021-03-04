/*
 * This is description.
 * goroutine可以使用recover来解决协程中出现的panic(导致当前协程崩溃的问题):
	说明: 如果我们启动了一个协程, 该协程抛出了panic, 如果没有捕获这个panic问题, 则整个程序会崩溃. 此时我们可以使用recover来捕获panic, 进行解决, 即使这个协程出了问题, 但是主进程将继续正常执行.
 * @author Chris0710
 * @date 2021/3/4 14:05
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	go test()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("main(). ")
	}
	/*
		test()抛错:  assignment to entry in nil map
		main().
		hello ~
		hello ~
		main().
		main().
		hello ~
		hello ~
		main().
		hello ~
		main().
	*/
}

// 使用defer + recover解决panic
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()抛错: ", err)
		}
	}()

	// 使用map前, 不进行make分配空间, 模拟一个错误(panic: assignment to entry in nil map).
	var m map[int]string
	m[0] = "hello"

}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello ~")
	}
}
