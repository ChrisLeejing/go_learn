/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/3 14:59
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	primeChan := make(chan int, 2000)
	start := time.Now().Unix()

	for i := 1; i <= 80000; i++ {
		flag := true
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- i
		}
	}

	end := time.Now().Unix()
	fmt.Println(end - start)

}
