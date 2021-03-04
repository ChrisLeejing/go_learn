/*
 * This is description.
 * 需求:
	1. 要求统计1-200000的数字中，哪些是素数? 学习了goroutine和channel的知识后，就可以完成了[测试数据: 80000].
	2. 传统的方法，就是使用一个循环，循环的判断各个数是不是素数[ok].
	3. 使用并发/并行的方式，将统计素数的任务分配给多个(4个)goroutine去完成，完成任务时间短.
	4. 协程后，执行的速度，比普通方法提高至少4倍.
	5. 示意图: goroutine_channel02.png
 * @author Chris0710
 * @date 2021/3/3 14:31
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)

	// 开启统计时间.
	start := time.Now().Unix()
	// 1. 开启1个协程, 写入1~8000数据.
	go putNum(intChan)

	// 2. 开启4个协程, 取出数据, 判断是否为素数
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 3. 开启一个匿名协程: 等待标识管道中取出4个标识值.
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// exitChan取完后, 此时可关闭primeChan.
		close(primeChan)
	}()

	// 4. 遍历primeChan, 取出素数.
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数的值: %d\n", v)
	}

	end := time.Now().Unix()
	fmt.Printf("用时: %d\n", end-start)
	fmt.Printf("main主线程退出. ")
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		// time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			// 当从intChan中, 取不出数据时, 退出.
			break
		}
		// 首先假设为素数.
		flag = true
		// 判断是否真的为素数. 备注: j < num, 最好改为j<math.Sqrt(num), 需要int与float64转换.
		for j := 2; j < num; j++ {
			if num%j == 0 {
				flag = false
				break
			}
		}
		// 该数为素数, 写入primeChan中.
		if flag {
			primeChan <- num
		}
	}
	// 此时, 不能关闭primeChan管道(即: close(primeChan)), 原因: 其他协程可能正在同时处理primeChan管道.
	fmt.Println("4个协程中, 有一个primeNum取不到数据了, 退出. ")
	exitChan <- true
}

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}
