/**
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/11 10:28
 */
package main

import (
	"fmt"
)

func main() {
	// var dayNum byte
	// fmt.Println("请输入数字: 1, 2, 3, 4, 5, 6, 7, 你输入的数字为: ")
	// fmt.Scanln(&dayNum)
	// switch dayNum { // 注:
	// case 1, 8: // 注: case可以写多个值进行匹配, 用逗号隔开.
	// 	fmt.Println("星期一")
	// case 2:
	// 	fmt.Println("星期二")
	// case 3:
	// 	fmt.Println("星期三")
	// case 4:
	// 	fmt.Println("星期四")
	// case 5:
	// 	fmt.Println("星期五")
	// default:
	// 	fmt.Println("周六或周日")
	// }
	//
	// switch { // 注: 省略switch后的表达式, 等价于if...else判断.
	// case dayNum == 5:
	// 	fmt.Println("星期五")
	// default:
	// 	fmt.Println("周六或周日")
	// }
	//
	// switch dayNum := 4; { // 注: switch后的表达式可以进行赋值运算, 但是不推荐.
	// case dayNum == 5:
	// 	fmt.Println("星期五")
	// default:
	// 	fmt.Println("周六或周日")
	// }

	fmt.Println("------------")

	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 默认只穿透一层: case 20, 即: 忽略case 20的判断条件, 且只能写在case块的最后一句.
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("not ok.")
	}

	fmt.Println("------------")

	// switch 用于type-switch中, 判断interface变量中实际指向的变量类型.
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型为: %T\n", i)
	case int:
		fmt.Println("x的类型为int类型.")
	case float64:
		fmt.Println("x的类型为float64类型.") // x的类型为float64类型.
	case bool, string:
		fmt.Println("x的类型为bool类型或string类型.")
	default:
		fmt.Println("x为未知类型.")
	}

	/*
		总结:
			1. 对于判断的类型不多, 且数据类型是整数, 浮点数, 字符, 字符串等, 适合用switch进行判断, 简洁高效.
			2. 对于区间判断, 或是bool判断, 建议使用if语句.
	*/
}
