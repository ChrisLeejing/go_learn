/**
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/11 9:59
 */
package main

import "fmt"

func main() {
	var age int

	fmt.Println("请输入您的年龄: ")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你已经成年, 你的年龄为：", age)
	} else {
		fmt.Println("你还未成年, 你的年龄为：", age)
	}

	fmt.Println("-------------")

	if age := 20; age > 18 {
		fmt.Println("你已经成年.")
	}

}
