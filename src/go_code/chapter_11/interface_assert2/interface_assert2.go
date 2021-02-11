/*
 * This is description.
 * 断言最佳实践2: 写一个函数, 循环判断出参数的类型是什么(bool, int, int32, int64, float32, float64, string, Student, *Student)?
 * @author Chris0710
 * @date 2021/2/11 12:42
 */
package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func JudgeType(items ...interface{}) {
	for i, item := range items {
		switch item.(type) {
		case int, int32, int64:
			fmt.Printf("第%v个参数的类型是int, 值: %v\n", i, item)
		case float32:
			fmt.Printf("第%v个参数的类型是float32, 值: %v\n", i, item)
		case float64:
			fmt.Printf("第%v个参数的类型是float64, 值: %v\n", i, item)
		case string:
			fmt.Printf("第%v个参数的类型是string, 值: %v\n", i, item)
		case bool:
			fmt.Printf("第%v个参数的类型是bool, 值: %v\n", i, item)
		case Student:
			fmt.Printf("第%v个参数的类型是Student, 值: %v\n", i, item)
		case *Student:
			fmt.Printf("第%v个参数的类型是*Student, 值: %v\n", i, item)
		default:
			fmt.Printf("第%v个参数的类型是其他类型, 值: %v\n", i, item)
		}
	}

}
func main() {
	var a1 int = 1
	var a2 int32 = 2
	var a3 int64 = 3
	var a4 float32 = 1.0
	var a5 = 3.14
	var a6 = "hello, go"
	var a7 = false
	var a8 = [3]int{1, 2, 3}
	JudgeType(a1, a2, a3, a4, a5, a6, a7, a8)
	fmt.Println("---------")
	student1 := Student{"chris"}
	student2 := &Student{"jack"}
	JudgeType(student1, student2)
}
