/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/1 11:03
 */
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前操作系统的CPU数量.
	numCPU := runtime.NumCPU()
	// 设置程序允许执行的最大CPU数量.
	// 1.8 之前需要设置, 使得CPU利用率更高.
	// 1.8 之后默认采用多核执行, 不需要显示设置.
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("numCPU: ", numCPU)
}
