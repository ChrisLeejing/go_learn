/*
 * This is description.
 * 演示常见的channel使用场景.
 * @author Chris0710
 * @date 2021/3/1 17:08
 */
package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 1. 使用intChan进行存3个int值, 并进行取出, 同时演示存入过多的数据, 或者取出过多的数据.
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	// intChan <- 4 // fatal error: all goroutines are asleep - deadlock!
	n1 := <-intChan
	n2 := <-intChan
	n3 := <-intChan
	// n4 := <-intChan // fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(n4)
	fmt.Printf("从intChan取出的值, n1 = %v, n2 = %v, n3 = %v\n", n1, n2, n3)

	// 2. 创建一个mapChan, 可以存放10个map[string]string, 演示存入和取出操作.
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["name1"] = "chris"
	m1["name2"] = "jack"
	mapChan <- m1
	m2 := make(map[string]string, 5)
	m2["city1"] = "chengdu"
	m2["city2"] = "beijing"
	mapChan <- m2
	res1 := <-mapChan
	res2 := <-mapChan
	fmt.Printf("从mapChan中取出的值, res1 = %v, res2 = %v\n", res1, res2)

	// 3. 创建一个catChan, 用于存放10个Cat类型的结构体变量, 演示写入和读取的操作.
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	catChan <- Cat{"tom1", 10}
	catChan <- Cat{"tom2", 20}
	cat1 := <-catChan
	cat2 := <-catChan
	fmt.Printf("cat1: %v, cat2: %v\n", cat1, cat2)

	// 4. 创建一个catChan2, 用于存放10个*Cat类型的结构体变量, 演示写入和读取的操作.
	// var catChan2 chan *Cat
	// catChan2 = make(chan *Cat, 10)
	// var catChan2 = make(chan *Cat, 10)
	catChan2 := make(chan *Cat, 10)
	catChan2 <- &Cat{"tom11", 30}
	catChan2 <- &Cat{"tom22", 40}

	cat11 := <-catChan2
	cat22 := <-catChan2
	fmt.Printf("cat11: %v, cat22: %v\n", cat11, cat22)

	// 5. 创建一个allChan, 用于存放所有数据类型的变量, 演示写入和读取的操作.
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	allChan <- Cat{"tom33", 30}
	allChan <- 10
	allChan <- "chris"
	cat5 := <-allChan
	v2 := <-allChan
	v3 := <-allChan
	fmt.Printf("cat5: %v, v2: %v, v3: %v\n", cat5, v2, v3)

	// 6. 在示例5的基础上, 判断: cat6.Name 会输出什么?
	allChan <- Cat{"tom6", 30}
	cat6 := <-allChan
	fmt.Printf("cat6: %T, cat6: %v\n", cat6, cat6) // cat6: main.Cat, cat6: {tom6 30}
	// fmt.Println(cat6.Name) // 编译不通过.
	// 使用类型断言
	cat66 := cat6.(Cat)
	fmt.Printf("cat66: %T, cat66: %v\n", cat66, cat66)
	fmt.Printf("cat66 Name: %v, cat66 Age: %v\n", cat66.Name, cat66.Age) // cat66 Name: tom6, cat66 Age: 30
}
