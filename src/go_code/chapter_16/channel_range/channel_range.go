/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/3 9:47
 */
package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i + 1
	}
	// 遍历channel之前必须先关闭channel, 否则将抛出: fatal error: all goroutines are asleep - deadlock!
	close(intChan)
	// 备注: 遍历channel暂不支持使用普通的for循环.
	for v := range intChan {
		fmt.Printf("v: %d\n", v)
	}
}
