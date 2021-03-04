/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/3 9:37
 */
package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	close(intChan)
	// intChan <- 3 // panic: send on closed channel
	v1, ok1 := <-intChan
	fmt.Println(v1, ok1)
	v2, ok2 := <-intChan
	fmt.Println(v2, ok2)
	fmt.Printf("intChan v1: %v, v2: %v.\n", v1, v2)

	v3, ok := <-intChan
	fmt.Println(v3, ok)
}
