/*
 * This is description.
 * 需求: 请完成goroutine和channel协同工作的案例, 具体要求如下:
	1. 开启一个writeData协程, 向管道intChan中写入50个整数.
	2. 开启一个readData协程, 从管道intChan中读取writeData写入的整数.
	3. 注意: writeData和readData操作的是同一个管道.
	4. 主线程需要等待writeData和readData协程同时都结束工作时, 才退出管道.
   扩展: channel阻塞问题的提出, 即: 如果注释掉go readData(intChan, exitChan), 程序会发生什么?
	1. 如果只是向管道中写入数据, 而没有读取数据, 就会出现阻塞现象, 抛错: fatal error: all goroutines are asleep - deadlock!
	2. 原因: 因为intChan容量为10, 而writeData(...)会向intChan中写入50个数据, 所以抛错点在: intChan <- i.
	3. 如果读管道和写管道的频率不一致, 其实是没有关系的, 编译器判断, 只要有读取操作, 程序就会正常运行.
 * @author Chris0710
 * @date 2021/3/3 10:00
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)

	go readData(intChan, exitChan)
	for {
		v, ok := <-exitChan
		fmt.Printf("v: %v, ok: %v\n", v, ok)
		// ok: true -> 从exitChan管道中取到值(代码执行到exitChan <- true) -> !OK = false -> 主线程等待.
		// ok: false -> 从管道中没有取到值, 说明exitChan管道关闭(代码执行到close(exitChan)), 即: readData(...)读取数据, 执行结束. -> !OK = true -> 主线程退出.
		// 主线程退出的标识: 必须让readData(...)执行完毕, 且关闭exitChan管道.
		// 这里!ok和ok是一样的结果
		// 因为!ok指的取到后再取一次取不到(exitChan管道关闭), ok是直接取到exitChan管道中的true值(更直接)
		if !ok {
			break
		}
	}
	fmt.Println("主线程结束.")
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		// 写法1:
		// if v, ok := <-intChan; ok {
		// 	fmt.Printf("readData: %d\n", v)
		// } else {
		// 	break
		// }
		// 写法2:
		v, ok := <-intChan
		if !ok {
			break
		}
		// 尽管读取数据频率和写入数据频率不一致, 但是毫无影响, Go认为只要在读取数据就行.
		time.Sleep(time.Second)
		fmt.Printf("readData: %d\n", v)
	}
	exitChan <- true
	// 关闭管道
	close(exitChan)
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData: %d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	// 关闭管道
	close(intChan)
}
