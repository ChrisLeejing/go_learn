/*
 * This is description.
 * 需求：现在要计算 1-200 的各个数的阶乘, 并且把各个数的阶乘放入到map中, 最后显示出来.
	要求使用goroutine完成.
	思路:
	1. 编写一个函数统计传入值的阶乘, 并将结果值放入到一个全局定义的map中.
	解决fatal error: concurrent map writes问题, 也即是不同的goroutine之间如何通讯的问题:
	1. 通过加入全局互斥锁机制.
	2. 引出channel通讯机制.

 * @author Chris0710
 * @date 2021/3/1 14:38
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// 使用map前, 需要make分配内存地址空间
	m    = make(map[int]int, 10)
	lock sync.RWMutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	lock.Lock()   // 加锁
	m[n] = res    // golang里map是线程不安全的, 如果在并发读写时, 会出现fatal error: concurrent map writes.
	lock.Unlock() // 解锁
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range m {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
