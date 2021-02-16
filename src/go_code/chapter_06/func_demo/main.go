/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/11 22:46
 */
package main

import "fmt"

// 值传递方式.
func getAdd(n int) int {
	n++
	fmt.Println("getAdd()..., n =", n)
	return n // 当只有一个返回值时, (int)可以写成int.
}

// 指针传递方式.
func getAdd2(n *int) *int {
	*n++
	fmt.Println("getAdd()..., n =", n)
	fmt.Println("getAdd()..., *n =", *n)
	return n // 当只有一个返回值时, (int)可以写成int.
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// 支持对函数多个返回值进行命名处理.
func getSumAndSub2(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFunc(funcVar func(int, int) int, n1 int, n2 int) int {
	return funcVar(n1, n2)
}

// 自定义函数类型
type myFuncType func(int, int) int

// 函数作为形参变量进行调用.
func myFunc2(funcTypeVar myFuncType, n1 int, n2 int) int {
	return funcTypeVar(n1, n2)
}

// 支持0~n个可变参数, args是slice切片, 可以通过args[i]将切片中的值取出来.
func sum1(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// 支持1~n个可变参数, args是slice切片, 可以通过args[i]将切片中的值取出来.
func sum2(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func sum3(n1, n2 float32) float32 {
	return n1 + n2
}

// 需求: 请编写一个函数swap(n1 *int, n2 *int), 用于交换n1与n2的值.
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

// init函数: 每一个源文件都包含一个init函数, 该函数会在main()函数调用之前, 被Go语言框架进行调用.

func init() {
	fmt.Println("This ia a init() function.")
	/*
		This ia a init() function.
		This ia a main() function.
		...
	*/
}

/*
	1. 如果一个文件中, 同时包含有全局变量, init(), main(), 则执行的顺序为: 全局变量 -> init() -> main().
		执行顺序1: global variable num int = test()...
		执行顺序2: This ia a init() function.
		执行顺序3: This ia a main() function.
	2. init()作用: 在程序开始之前, 进行一些初始化的工作.
	3. 执行顺序:
		场景:
		read_demo.go								util.go
		import (						// 执行顺序1: 全局变量       <------
			util						var b string
		)
		// 执行顺序3: 全局变量			// 执行顺序2: init()函数
     	var a string	<-------			func init() {
											...
		// 执行顺序4: init()函数			}
		func init() {
			util.b
			...
		}

		// 执行顺序5: main()主函数
		func main() {
			...
		}
*/
var age int = test()

func test() int {
	fmt.Println("global variable num int = test()...")
	return 20
}

// 全局匿名函数
var globalAnonymousFunc = func(n1 int, n2 int) int {
	return n1 - n2
}

func main() {
	fmt.Println("This ia a main() function.")
	result := getAdd(10)
	fmt.Println("result =", result)

	sum, sub := getSumAndSub(10, 20)
	fmt.Printf("sum =%d, sub = %d\n", sum, sub)

	_, sub2 := getSumAndSub(10, 20) // _: 表示占位省略.
	fmt.Printf("sub = %d\n", sub2)

	// 基本数据类型和数组都是值传递, 即: 进行值拷贝. 在函数内修改, 并不会影响到原来的值.
	n := 12
	getAdd(n)
	fmt.Println("main()..., n =", n)
	fmt.Println("-----------")

	/*
		getAdd()..., n = 13
		main()..., n = 12
	*/

	// 如果希望函数内的值改变, 使其函数外的值也进行变化, 可以传递变量的地址&, 函数内已指针的方式传递变量.
	getAdd2(&n)
	fmt.Println("main()..., n =", n)
	/*
		getAdd()..., n = 0xc0000140f8
		getAdd()..., *n = 13
		main()..., n = 13
	*/

	fmt.Println("------------")

	// 8) 在Go中, 函数也是一种数据类型, 可以赋值给一个变量, 则该变量就是一个函数类型的变量, 可以通过该变量对函数进行调用.
	a := getSum
	fmt.Printf("a的类型为: %T, getSum()的类型为: %T\n", a, getSum)

	res := a(10, 40)
	fmt.Println("res =", res)
	/*
		a的类型为: func(int, int) int, getSum()的类型为: func(int, int) int
		res = 50
	*/

	fmt.Println("------------")

	// 函数作为形参变量进行调用.
	res2 := myFunc(getSum, 100, 200)
	fmt.Println("res2 =", res2) // res2 = 300
	/*
			func myFunc(funcVar func(int, int) int, n1 int, n2 int) int {
				return funcVar(n1, n2)
			}
		分析: myFunc(getSum, 100, 200) -> return getSum(100, 200) -> return 100 + 200 -> return 300.
	*/

	fmt.Println("-----------")
	// Go语言中, 支持自定义数据类型.
	type myInt int // 相当于给int取了一个别名: myInt, 虽然myInt也是int类型, 但是Go会认为myInt和int是两个不同的数据类型.
	var n1 int = 20
	var n2 myInt = myInt(n1) // Go会认为myInt和int是两个不同的数据类型, 所以此处需要转换类型.
	fmt.Println("n2 =", n2)

	fmt.Println("-----------")
	func2Res := myFunc2(getSum, 100, 100)
	fmt.Println("func2Res =", func2Res) // func2Res = 200

	fmt.Println("-----------")
	// 支持对函数多个返回值进行命名处理.
	sum, sub = getSumAndSub2(11, 22)
	fmt.Printf("sum =%d, sub =%d\n", sum, sub) // sum =33, sub =-11

	fmt.Println("-----------")
	// Go 支持可变参数
	sum1 := sum1(1, 2, 3, 4, 5)
	fmt.Println("sum1 =", sum1) // sum1 = 15

	sum2 := sum2(1, 2, 3, 4, 5, 6)
	fmt.Println("sum2 =", sum2) // sum2 = 21

	fmt.Println("-----------")
	sum3 := sum3(1, 2)
	fmt.Println("sum3 =", sum3) // sum3 = 3

	// var n3, n4 int = 1, 2
	// sum3res := sum3(n3, n4) // 此时会编译出错, 类型不匹配.
	// fmt.Println("sum3res =", sum3res)

	fmt.Println("-----------")

	var n5 int = 1
	var n6 int = 2
	fmt.Printf("交换前: n5 =%d, n6 =%d\n", n5, n6)
	swap(&n5, &n6) // 传入变量的指针地址.
	fmt.Printf("交换后: n5 =%d, n6 =%d\n", n5, n6)

	// init函数: 详见文件头部func init(){}

	fmt.Println("-----------")
	// 匿名函数1:
	res7 := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 2)
	fmt.Println("匿名函数调用1, res7 =", res7)

	// 匿名函数2: 将一个函数赋值给一个变量(函数变量), 则可以通过这个变量来进行调用匿名函数.
	anonymousFun := func(n1 int, n2 int) int {
		return n1 * n2
	}
	res8 := anonymousFun(2, 3)
	fmt.Println("匿名函数调用2, res8 =", res8)

	// 全局匿名函数: 将一个匿名函数赋值给一个全局变量, 则该全局变量为一个全局匿名函数, 可以在程序中有效.
	anonymousFuncRes := globalAnonymousFunc(3, 4)
	fmt.Println("全局匿名函数, anonymousFuncRes =", anonymousFuncRes) // 全局匿名函数, anonymousFuncRes = -1

}
