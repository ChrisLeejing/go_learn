/*
 * This is description.
 * goruntine协程入门:
	需求:
	1. 在主线程中开启一个goruntine, 在该协程中, 循环10次, 每隔1秒输出一句"hello, world".
	2. 在主线程中, 循环10次, 也每隔1秒输出一句"hello, world".
	3. 主线程和协程同时运行.
	总结:
	1. 主线程是一个物理线程, 直接作用在CPU上, 重量级, 比较耗费资源.
	2. go协程是一种从主线程中分出来的一种逻辑态, 可以理解为一种轻量级的线程, 耗费资源量小.
	3. go语言中, 可以很轻松的开启上万个协程, 相较于其他语言基于线程的并发机制(比如: Java), 耗费资源量小, 也体现出go在并发中的优势.
 * @author Chris0710
 * @date 2021/2/26 15:49
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello, world. ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() // 开启一个goruntine.

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello, world. ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	/*
		main() hello, world.  0
		test() hello, world.  0
		test() hello, world.  1
		main() hello, world.  1
		test() hello, world.  2
		main() hello, world.  2
		main() hello, world.  3
		test() hello, world.  3
		test() hello, world.  4
		main() hello, world.  4
		main() hello, world.  5
		test() hello, world.  5
		test() hello, world.  6
		main() hello, world.  6
		main() hello, world.  7
		test() hello, world.  7
		test() hello, world.  8
		main() hello, world.  8
		main() hello, world.  9
		test() hello, world.  9
	*/
}
