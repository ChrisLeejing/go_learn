/*
 * This is description.
 * channel使用细节和注意事项:
	1. 使用select解决channel阻塞问题.
 * @author Chris0710
 * @date 2021/3/4 11:30
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	var strChan chan string
	strChan = make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法在遍历channel时, 如果没有关闭channel, 会因为阻塞而发生deadlock.
	// 问题: 在实际开发过程中, 我们可能不确定具体什么时候关闭channel, 这时, 我们可以通过select来解决.
label1:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("intChan 取出的数据: %d\n", v)
			time.Sleep(time.Second)
		case v := <-strChan:
			fmt.Printf("strChan 取出的数据: %s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("所有channel都取不到数据了, 此时, 可以加入常规的业务逻辑. ")
			time.Sleep(time.Second)
			// return
			break label1
		}
	}
}
