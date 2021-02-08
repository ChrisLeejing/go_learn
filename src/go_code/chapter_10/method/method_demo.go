/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/8 15:58
 */
package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

// 给Person类型绑定一个方法, 该test()方法只能通过Person类型的变量进行调用, 不能直接调用, 也不能通过其他类型的变量进行调用.
func (p Person) test() {
	fmt.Println("Person test() method: ", p.Name)
}

// 方法的调用和函数的调用基本一样, 不同的是, 方法的调用会将调用方法的变量作为实参一起传递给方法. 变量类型是值类型, 则进行值拷贝, 变量类型为引用类型, 则进行地址拷贝.
func (p Person) getSum(n1 int, n2 int) int {
	fmt.Println("Person getSum(n1 int, n2 int) method: ", p.Name)
	return n1 + n2
}

type integer int

func (i *integer) print() {
	fmt.Println("修改前 i =", *i)
}

func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	res := fmt.Sprintf("stu Name: %v, stu Age: %v", stu.Name, stu.Age)
	return res
}

// 对于普通函数, 当函数接收为值类型(p Person)时, 不能通过指针类型test1(&p1)进行传递.
func test1(p Person) {
	p.Name = "tom1"
	fmt.Println(p.Name)
}

// 对于普通函数, 当函数接收为指针类型(p *Person)时, 不同通过值类型test2(p1)进行传递.
func test2(p *Person) {
	fmt.Println(p.Name)
}

// 对于方法(如: struct方法), 当接收者为值类型时, 可以通过指针类型进行传递.
func (p Person) test3() {
	p.Name = "tom3"
	fmt.Println("test3() ", p.Name)
}

// 对于方法(如: struct方法), 当接收者为指针类型时, 可以通过值类型进行传递.
func (p *Person) test4() {
	p.Name = "tom4"
	fmt.Println("test4() ", p.Name)
}

func main() {
	p1 := Person{"chris"}
	p1.test()
	// d1 := Dog{"tidy"}
	// d1.test() // dog变量中没有相对应的test()方法, 编译报错, 不能进行调用.

	sum1 := p1.getSum(1, 2)
	fmt.Println("sum1 : ", sum1)
	fmt.Println("--------------")
	/*
		Person getSum(n1 int, n2 int) method:  chris
		sum1 :  3
	*/

	// 注:
	// 1. func (p Person) test() {}: 结构体类型是值类型, 在方法调用过程是通过值传递方式进行调用.
	// 2. func (p *Person) test() {}: 当通过结构体指针方式进行传递时, 可以修改结构体中变量的值.
	// 3. go语言中, 方法作用在指定的数据类型上, 而不是仅仅只作用在struct上, 所以自定义的方法也是可以调用方法的.

	var i integer = 10
	i.print()
	i.change()
	fmt.Println("修改后 i =", i)
	/*
		修改前 i = 10
		修改后 i = 11
	*/

	// 4. 方法的访问范围, 与函数类似, 如果首字母小写, 则只能在当前package内进行访问, 如果首字母大写, 则在package外可以进行访问.
	// 5. 如果一个类型实现了String()方法, 那么当调用fmt.Println()方法时, 会默认调取String()方法.
	s1 := Student{
		Name: "chris",
		Age:  29,
	}
	fmt.Println(s1)  // {chris 29}
	fmt.Println(&s1) // stu Name: chris, stu Age: 29

	// 6. 方法和函数区别:
	// 6.1 调用方式不一样:
	// 	函数调用方式: 函数名(实参列表)
	// 	方法调用方式: 变量.方法名(实参列表)
	p1 = Person{
		Name: "tom",
	}
	test1(p1)
	// test1(&p1) // 编译抛错.
	test2(&p1)
	// test2(p1) // 编译抛错.
	fmt.Println("main p1: ", p1)
	fmt.Println("--------------")

	p1.test3()
	fmt.Println("main() ", p1.Name)
	(&p1).test3() // 编译通过: 形式上传递的是地址, 实质上是进行值拷贝.
	fmt.Println("main() ", p1.Name)
	/*
		test3()  tom3
		main()  tom
		test3()  tom3
		main()  tom
	*/

	fmt.Println("--------------")
	(&p1).test4()
	fmt.Println("main() ", p1.Name)
	p1.test4() // 编译通过: 形式上传递的是值, 实质上是进行地址拷贝. 等价于(&p1).test4()
	fmt.Println("main() ", p1.Name)
	/*
	   test4()  tom4
	   main()  tom4
	   test4()  tom4
	   main()  tom4
	*/

}
