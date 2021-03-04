/*
 * This is description.
 * channel使用细节和注意事项:
	1. channel管道可以声明为只读或只写.
	2. channel只读, 只写的最佳实践案例.
 * @author Chris0710
 * @date 2021/3/4 10:46
*/
package main

import "fmt"

func main() {
	// 1.1 默认情况下, 可读可写.
	// var intChan1 chan int

	// 1.2 声明为只写
	var intChan2 chan<- int
	intChan2 = make(chan<- int, 2)
	intChan2 <- 1
	// <-intChan2 // Invalid operation: <-intChan2 (receive from send-only type chan<- int)

	// 1.3 声明为只读
	// var intChan3 <-chan int
	// s := <-intChan3
	// intChan3 <- 3 // Invalid operation: intChan3 <- 3 (send to receive-only type <-chan int)
	// fmt.Println(s)

	// 2. channel只读, 只写的最佳实践案例.
	var ch chan int
	ch = make(chan int, 10)
	var exitChan chan struct{}
	exitChan = make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("main 结束.")
}

// ch <-chan int: 只读
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}

// ch chan<- int: 只写
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
