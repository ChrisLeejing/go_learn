/*
 * This is description.
 * 时间和日期相关的函数
 * @author Chris0710
 * @date 2021/1/13 22:32
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("now is %v, type is %T\n", now, now) // now is 2021-01-13 22:36:29.902365 +0800 CST m=+0.004006501, type is time.Time

	// 2. 获取年, 月, 日, 时, 分, 秒.
	fmt.Println("year: ", now.Year())
	fmt.Println("month: ", now.Month())
	fmt.Println("month: ", int(now.Month()))
	fmt.Println("day: ", now.Day())
	fmt.Println("hour: ", now.Hour())
	fmt.Println("minute: ", now.Minute())
	fmt.Println("second: ", now.Second())

	// 3. 格式化日期:
	// 3.1 fmt.Printf, fmt.Sprintf：当前时间: 2021-1-13 22:46:28
	fmt.Printf("当前时间: %d-%d-%d %d:%d:%d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	date := fmt.Sprintf("当前时间: %d-%d-%d %d:%d:%d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(date)

	// 3.2 使用time.Format(), 备注: 此时间为固定时间, 不可更改.
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	// 时间常量
	// 4.1 每隔1秒打印一个数字, 直到打印到100.
	// 4.2 每隔0.1秒(100毫秒)打印一个数字, 直到打印到100.
	i := 0
	for {
		i++
		// time.Sleep(time.Second)
		// time.Sleep(time.Second / 10) // ok
		// time.Sleep(0.1 * time.Second) // not ok, 编译报错.
		// time.Sleep(100 * time.Millisecond) // ok
		fmt.Println("i =", i)
		if i == 100 {
			break
		}
	}

	// 5. 时间戳: 当前时间戳, Unix: 1610550645, UnixNano: 1610550645322430300
	fmt.Printf("当前时间戳, Unix: %d, UnixNano: %d\n", now.Unix(), now.UnixNano())

}
