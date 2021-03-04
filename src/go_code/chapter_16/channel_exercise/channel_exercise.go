/*
 * This is description.
 * 需求:
	1. 创建一个Person结构体[Name, Age, Address].
	2. 使用random随机创建10个Person放入channel中.
	3. 对channel进行遍历, 进行打印输出到控制台中.
 * @author Chris0710
 * @date 2021/3/2 10:17
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var personChan = make(chan Person, 10)
	for i := 1; i <= 10; i++ {
		person := Person{"chris" + strconv.Itoa(rand.Intn(100)), 30, "成都"}
		personChan <- person
	}

	// for person := range personChan {
	// 	fmt.Printf("person: %v\n", person)
	// }
	for i := 1; i <= 10; i++ {
		person := <-personChan
		fmt.Printf("person: %v\n", person)
	}
}
