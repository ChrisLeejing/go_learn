/*
 * This is description.
 * channel_demo 通讯机制:
	1. channel是一种带有数据类型的(string, int), 线程安全的(sync), 先进先出的(FIFO)的队列数据结构.
	2. channel是引用数据类型, 必须make初始化之后, 才能进行使用.
	3. 示例: 如下所示(L12 ~ L19)
	小结:
	1. channel只能存放指定数据类型的数据.
	2. 当向channel中放入数据的数量超过channel的容量, 或者从channel中取数据的数量超过channel的容量时, 这两种情况下, 都会抛错: deadlock.
	3. 向channel中取出数据后, channel长度减少, 此时可以继续向channel中放入数据.
 * @author Chris0710
 * @date 2021/3/1 15:42
*/
package main

import "fmt"

var intChan chan int
var strChan chan string
var mapChan chan map[int]string
var perChan chan Person
var perChan2 chan *Person

type Person struct {
}

func main() {
	// 1. 演示一下channel的用法.
	var intChan1 chan int
	intChan1 = make(chan int, 3)
	fmt.Printf("intChan1的值: %v, intChan1的地址: %p\n", intChan1, &intChan1) // intChan1的值: 0xc000018100, intChan1的地址: 0xc000006028
	// 向管道channel中写入数据.
	intChan1 <- 1
	num := 2
	intChan1 <- num
	intChan1 <- 3
	// intChan1 <- 4 // 当放入数据的数量大于channel管道的容量时, 抛错: fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("intChan1的len=%d, cap=%d\n", len(intChan1), cap(intChan1))

	// 从管道channel中读数据.
	var n1 int
	n1 = <-intChan1
	fmt.Printf("n1 = %d, intChan1的len=%d, cap=%d\n", n1, len(intChan1), cap(intChan1))

	n2 := <-intChan1
	n3 := <-intChan1
	// n4 := <-intChan1 // 当读数据时, 超过了管道中的数量时, 抛错: fatal error: all goroutines are asleep - deadlock!
	// fmt.Printf("n2 = %d, n3 = %d, n4 = %d\n", n2, n3, n4)
	fmt.Printf("n2 = %d, n3 = %d\n", n2, n3)
}
