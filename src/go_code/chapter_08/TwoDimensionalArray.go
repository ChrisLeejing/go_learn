/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/19 15:38
 */
package main

import (
	"fmt"
)

func main() {
	// 1. 赋值方式
	// 1.1 先声明, 后赋值.
	var arr1 [2][3]int // [[0, 0, 0], [0, 0, 0]]
	arr1[1][1] = 1
	fmt.Println("arr1的值:", arr1)
	fmt.Printf("arr1[0]的地址:%p\n", &arr1[0])
	fmt.Printf("arr1[1]的地址:%p\n", &arr1[1])

	fmt.Printf("arr1[0][0]的地址:%p\n", &arr1[0][0])
	fmt.Printf("arr1[1][0]的地址:%p\n", &arr1[1][0])
	/*
		arr1的值: [[0 0 0] [0 1 0]]
		arr1[0]的地址:0xc00000a3c0
		arr1[1]的地址:0xc00000a3d8
		arr1[0][0]的地址:0xc00000a3c0
		arr1[1][0]的地址:0xc00000a3d8
	*/

	// 1.2 直接初始化.
	var arr2 [2][3]string = [2][3]string{{"1", "2", "3"}, {"a", "b", "c"}}
	fmt.Println("arr2:", arr2)

	// 2. 遍历二维数组
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Printf("arr[%v][%v] = %v\t", i, j, arr2[i][j])
		}
		fmt.Println()
	}

	for i, v1 := range arr2 {
		for j, v2 := range v1 {
			fmt.Printf("arr[%v][%v] = %v\t", i, j, v2)
		}
		fmt.Println()
	}
}
