/**
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/11 11:06
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	fmt.Println("-------")

	// another for-loop type.
	j := 0
	for j < 10 {
		fmt.Println("j:", j)
		j++
	}

	fmt.Println("-------")

	// endless loop: for-loop.
	k := 0
	for {
		if k < 10 {
			fmt.Println("k:", k)
		} else {
			fmt.Println("k已经大于或等于10, 退出循环.")
			break
		}
		k++
	}

	// for range: for-loop.
	var str string = "hello,world"
	for i := 0; i < len(str); i++ {
		// 注: 如果str中包含有中文, 则会出现乱码现象.
		fmt.Printf("字符: %c ", str[i])
	}

	fmt.Println()
	fmt.Println("----------")

	// 针对于"中文乱码"现象, 使用[]rune切片来进行完善.
	var str2 string = "hello,world中国"
	var str3 = []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符: %c ", str3[i])
	}

	fmt.Println("")

	// for range: 没有中文乱码的现象.
	for index, value := range str {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}

	fmt.Println("----------")

	for index, value := range str3 {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}

	fmt.Println("----------")
	// while, do ... while: 在golang中, 没有这种语法结构, 可以通过for{} + break(无限循环+break)来实现其对应的效果.
	// 需求: 随机给出一个1~100之间的数, 知道生成这个数为99停止, 看看一共用了多少次?
	// time.Now().Unix(): 返回一个从1970年01月01日 00:00:00 时间到现在的秒数.
	var count int = 0
	rand.Seed(time.Now().UnixNano())
	for {
		num := rand.Intn(100) + 1
		count++
		fmt.Println("num:", num)
		if num == 90 {
			break
		}

	}
	fmt.Println("生成99时, 总共运行的次数:", count)

	fmt.Println("----------------")

	// 使用label标签, 跳出指定的外层循环.
label1:
	for i := 0; i < 4; i++ {
		fmt.Println("i:", i)
		// label2:
		for j := 0; j < 10; j++ {
			if j == 5 {
				break label1
			}
			fmt.Println("j:", j)
		}
	}
	/*
		i: 0
		j: 0
		j: 1
		j: 2
		j: 3
		j: 4

	*/

	// continue: 结束本次循环, 继续执行下一次循环.
	for i := 0; i < 4; i++ {
		fmt.Println("i =", i)
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j:", j)
		}
		fmt.Println("---------")
	}

	// goto语句: 可以指定跳转到任意位置. 注: Go程序设计语言一般不推荐使用, 以免造成程序上逻辑混乱.
	var num3 int = 30
	if num3 > 20 {
		goto label3

	}
	fmt.Println("hello 1")
	fmt.Println("hello 2")
	fmt.Println("hello 3")
label3:
	fmt.Println("hello 4")
	fmt.Println("hello 5")
	fmt.Println("hello 6")

	fmt.Println("-------------")

	// return: 退出函数, 如果该函数是main(), 那么将终止程序.
	for i := 0; i < 5; i++ {
		if i == 2 {
			return // fmt.Println("hello......") 此行代码将不会执行.
		}
		fmt.Println("i =", i)
	}

	fmt.Println("hello......")
}
