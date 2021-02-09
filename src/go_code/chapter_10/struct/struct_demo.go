/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/7 22:09
 */
package main

import (
	"encoding/json"
	"fmt"
	"go_code/chapter_10/factory"
)

// 1. 结构体创建时, 字段默认值: 布尔类型(false), 数值(0), 字符串(""); 引用类型: 指针(nil), map(nil), slice(nil), 其中nil表示还没有分配空间, 需要先make后, 再进行使用.
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针
	slice  []int             // 切片
	m1     map[string]string // map
}

// 2. 结构体是值拷贝, 不同结构体之间字段相互独立.
type Book struct {
	Name  string
	price float64
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

type A struct {
	Num int
}
type B struct {
	Num int
}

type Student struct {
	Name string
	Age  int
}

type integer int

type Hero struct {
	Name  string `json:"name"` // `json: "name"` : 即为struct中的tag.
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1") // ok1
	}

	if p1.slice == nil {
		fmt.Println("ok2") // ok2
	}

	if p1.m1 == nil {
		fmt.Println("ok3") // ok3
	}

	p1.slice = make([]int, 5)
	p1.slice[0] = 100 // 使用前必须make, 否则抛错panic: runtime error: index out of range [0] with length 0
	p1.m1 = make(map[string]string)
	p1.m1["k1"] = "value1" // 使用前必须make, 否则抛错panic: runtime error: index out of range [0] with length 0
	fmt.Println(p1)

	b1 := Book{"三国演义", 30}
	b2 := b1
	b2.Name = "红楼梦"
	var b3 = Book{"水浒传", 33}
	fmt.Println("b1 =", b1) // b1 = {三国演义 30}
	fmt.Println("b2 =", b2) // b2 = {红楼梦 30}
	fmt.Println("b3 =", b3)

	// 3. 创建结构体, 以及对结构体进行赋值.
	// 3.1 方式1: 直接声明: var p1 Person
	// 3.2 方式2: 类型推导: b1 := Book{"三国演义", 30} 或者 var b3 = Book{"水浒传", 33}
	// 3.3 方式3: 结构化指针&: var b4 *Book = new(Book)
	var b4 *Book = new(Book)
	// 指针标准写法: (*b4).Name, 由于Go语言底层做了处理, 也可以写成b4.Name.
	b4.Name = "活着"
	(*b4).Name = "平凡的世界"
	fmt.Println("b4 =", b4)

	// 3.4 方式4: 结构化指针&: var b5 *Book = &Book{}
	var b5 *Book = &Book{}
	b5.Name = "围城"
	(*b5).Name = "蒋勋说红楼梦"   // go编译器底层对b5.Name进行了(*b5).Name的转换处理.
	fmt.Println("b5 =", b5) // b5 = &{蒋勋说红楼梦 0}

	// 4. struct类型的内存分配地址分析.
	var b6 Book
	b6.Name = "大学之路"
	b6.price = 40
	var b7 = &b6
	fmt.Println("b7 price: ", b7.price)    // b7 price:  40
	fmt.Println("b7 price: ", (*b7).price) // b7 price:  40

	b7.Name = "文明之光"
	fmt.Printf("b7 name: %v, b6 name: %v\n", b7.Name, b6.Name)    // b7 name: 文明之光, b6 name: 文明之光
	fmt.Printf("b7 name: %v, b6 name: %v\n", (*b7).Name, b6.Name) // b7 name: 文明之光, b6 name: 文明之光

	fmt.Printf("b6的地址值: %p\n", &b6)               // b6的地址值: 0xc000004620
	fmt.Printf("b7的地址值: %p, b7的值: %p\n", &b7, b7) // b7的地址值: 0xc000006030, b7的值: 0xc000004620

	// 5. 结构体中所有字段在内存中是连续的.
	// r1 有4个int值, 打印地址值, 在内存中是连续的, 如下:
	// r1.leftUp.x 的地址值: 0xc00000c460, r1.leftUp.y 的地址值: 0xc00000c468, r1.rightDown.x 的地址值: 0xc00000c470, r1.rightDown.y 的地址值: 0xc00000c478
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("r1.leftUp.x 的地址值: %p, r1.leftUp.y 的地址值: %p, r1.rightDown.x 的地址值: %p, r1.rightDown.y 的地址值: %p\n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// r2有2个*Point类型, 这2个*Point本身所指向的地址值是连续的, 但是他们所指向的地址值并不一定连续(与系统分配内存地址有关).
	// r2.leftUp本身的地址值: 0xc00003a270, r2.rightDown本身的地址值: 0xc00003a278
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp本身的地址值: %p, r2.rightDown本身的地址值: %p\n", &r2.leftUp, &r2.rightDown)
	// r2.leftUp所指向的地址值: 0xc000014100, r2.rightDown所指向的地址值: 0xc000014110
	fmt.Printf("r2.leftUp所指向的地址值: %p, r2.rightDown所指向的地址值: %p\n", r2.leftUp, r2.rightDown)

	// 6. 结构体中类型转换时, 必须需要有相同的字段(字段名称, 字段个数, 字段类型).
	var a A
	var b B
	b = B(a)
	fmt.Printf("a = %v, b =%v\n", a, b)

	// 7. 结构体给type进行重定义时, 类似于取别名, Go会认为是新的数据类型, 但是相互之间是可以转换的.
	type Stu Student
	var s1 Student
	var s2 Stu
	s2 = Stu(s1) // 这里需要强转.
	fmt.Println(s1, s2)

	var i integer = 10
	var j int
	j = int(i) // 这里需要强转.
	fmt.Println(i, j)

	// 8. struct字段中, 可以写一个tag, tag可以使用反射机制进行获取, 使用场景: 序列化, 反序列化.
	h1 := Hero{"风清扬", 60, "独孤九剑"}
	h1Json, err := json.Marshal(h1)
	if err != nil {
		fmt.Println("json处理错误: ", err)
	}
	fmt.Println("jsonString: ", string(h1Json)) // jsonString:  {"name":"风清扬","age":60,"skill":"独孤九剑"}, 注: 字段中的大写变为了小写.

	// 9. 创建结构体时, 指定字段值.
	// 方式一: 返回结构体
	book9 := Book{"Kubernetes in Action", 40.99}
	var book10 = Book{"Spring in Action", 50.90}
	fmt.Println(book9, book10)

	book11 := Book{
		Name:  "spring boot in action",
		price: 41.90,
	}
	var book12 = Book{
		Name:  "Spring microservice in action",
		price: 55.90,
	}

	fmt.Println(book11, book12)

	// 方式二: 返回结构体指针.
	var book13 *Book = &Book{"文明之光", 30.99}
	book14 := &Book{"格局", 40.99}
	book15 := &Book{
		Name:  "蒋勋说红楼梦",
		price: 19.99,
	}
	var book16 *Book = &Book{
		Name:  "爱你就像爱生命",
		price: 20.99,
	}
	fmt.Println(*book13, *book14, *book15, *book16)

	// 10. 工厂模式
	// 情况1: 当结构体的变量首字母为大写时, 通过<package_name>.<struct_name>可以进行调用.
	// p10 := factory.People{
	// 	Name:  "tom",
	// 	Score: 80.90,
	// }
	// fmt.Println("p10 =", p10)

	// 情况2: 当结构体的变量首字母为小写时, 可以通过创建工厂方法进行调用.
	// p11 := factory.NewPeople("chris", 100)
	// fmt.Println("p11 = ", *p11)
	// fmt.Println("people name: ", p11.Name, ", people score: ", p11.Score)

	// 情况3: 当结构体的变量首字母为小写时, 且其中的Name字段也是小写.
	p12 := factory.NewPeople("bob", 100)
	fmt.Println("p12 = ", *p12)
	// fmt.Println("people name: ", p12.name, ", people score: ", p12.Score) // 编译报错.
	fmt.Println("people name: ", p12.GetName(), ", people score: ", p12.Score)

}
