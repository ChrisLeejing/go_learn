/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/19 16:03
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. 声明map, 注: 声明 -> make分配内存 -> 赋值.
	// var <var_name> map[key_type]value_type
	// var a map[int]int
	// var a map[int]string
	// var a map[string]int
	// var a map[string]string
	// var a map[string]map[string]string
	var books map[int]string
	books = make(map[int]string, 4)
	books[1] = "红楼梦"
	books[2] = "水浒传"
	books[3] = "三国演义"
	fmt.Println("books:", books)
	books[2] = "西游记"
	fmt.Println("books:", books)

	// 2. map crud.
	// map[key] = value, 如果没有, 则增加; 否则, 进行修改.
	// func delete(m map[Type]Type1, key Type); 注: 如果需要删除map中所有的key值, 方案1: 逐个遍历, 进行删除; 方案2: 直接make一个新空间, 如: books = make(map[int]string), 后续等待GC进行垃圾回收.
	// delete(books, 2)
	// fmt.Println("books:", books)
	book1 := books[1]
	fmt.Println("book1:", book1)

	book2, flag2 := books[5]
	book1, flag1 := books[1]
	fmt.Println("flag1:", flag1)
	fmt.Println("flag2:", flag2)
	fmt.Println("book2:", book2)

	// books = make(map[int]string)
	// fmt.Println("books:", books)

	// 3. map遍历
	for i, v := range books {
		fmt.Printf("books: books[%v] = %v\n", i, v)
	}

	var m2 map[int]map[string]string // map: {1: {"a":"A"}, 2:{"b":"B"}}
	m2 = make(map[int]map[string]string, 2)
	m2[1] = make(map[string]string, 2) // you must make it, before you use map.
	m2[1]["a"] = "A"
	m2[1]["b"] = "B"

	m2[2] = make(map[string]string, 2)
	m2[2]["aa"] = "AA"
	m2[2]["bb"] = "BB"

	fmt.Println("m2:", m2) // m2: map[1:map[a:A b:B] 2:map[aa:AA bb:BB]]
	fmt.Println(len(m2))

	// 4. map 切片
	var slice1 []map[string]string
	slice1 = make([]map[string]string, 2)
	if slice1[0] == nil {
		slice1[0] = make(map[string]string, 2) // panic: assignment to entry in nil map, you must make it, before you use map.
		slice1[0]["name"] = "chris"
		slice1[0]["age"] = "12"
	}
	if slice1[1] == nil {
		slice1[1] = make(map[string]string, 2)
		slice1[1]["name"] = "jack"
		slice1[1]["age"] = "10"
	}
	fmt.Println("slice1:", slice1) // slice1: [map[age:12 name:chris] map[age:10 name:jack]]

	s1 := map[string]string{
		"name": "mary",
		"age":  "20",
	}
	slice1 = append(slice1, s1)
	fmt.Println("slice1:", slice1)
	/*
		slice1: [map[age:12 name:chris] map[age:10 name:jack]]
		slice1: [map[age:12 name:chris] map[age:10 name:jack] map[age:20 name:mary]]
	*/

	// 5. map 排序
	// golang中, 目前没有针对于map进行排序的函数, 默认情况下, 是属于自动随机排序. 如果需要排序, 需要进行将key值进行排序后, 再按照key对应的value进行相关排序.
	m1 := make(map[int]int, 5)
	m1[2] = 200
	m1[1] = 100
	m1[4] = 400
	m1[3] = 300
	fmt.Println(m1)

	// 将map的key值放入到切片中
	// 对切片进行排序
	// 遍历切片, 根据对应的key进行输出
	var keys []int

	for i, _ := range m1 {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	fmt.Println("keys:", keys)

	for _, v := range keys {
		for i, j := range m1 {
			if v == i {
				fmt.Printf("按照key进行排序后的结果: m1[%v]=%v\n", i, j)
			}
		}
	}
	/*
		keys: [1 2 3 4]
		按照key进行排序后的结果: m1[1]=100
		按照key进行排序后的结果: m1[2]=200
		按照key进行排序后的结果: m1[3]=300
		按照key进行排序后的结果: m1[4]=400
	*/

	// 6. 注意事项
	// 6.1 map是引用类型, 进行引用传递.
	// 6.2 map容量不足时, 会自动进行扩容.
	var m3 map[int]int
	m3 = make(map[int]int, 5)
	m3[1] = 1000
	m3[2] = 2000
	m3[3] = 3000
	m3[4] = 4000
	m3[5] = 5000
	m3[6] = 6000
	m3[7] = 7000
	fmt.Println("m3修改前:", m3)
	modify(m3)
	fmt.Println("m3修改后:", m3)
	/*
		m3修改前: map[1:1000 2:2000 3:3000 4:4000 5:5000 6:6000 7:7000]
		m: map[1:1001 2:2000 3:3000 4:4000 5:5000 6:6000 7:7000]
		m3修改后: map[1:1001 2:2000 3:3000 4:4000 5:5000 6:6000 7:7000]
	*/

	// 7. map的value值, 可以使用struct结构体, 封装复杂的对象.
	m4 := make(map[int]Student, 5)

	nick := Student{
		Name:  "nick",
		Age:   28,
		Email: "nick@xxx.com",
	}

	jack := Student{
		Name:  "jack",
		Age:   28,
		Email: "jack@xxx.com",
	}
	m4[1] = nick
	m4[2] = jack
	fmt.Println(m4)
	/*
		map[1:{nick 28 nick@xxx.com} 2:{jack 28 jack@xxx.com}]
	*/
}

type Student struct {
	Name  string
	Age   int
	Email string
}

func modify(m map[int]int) {
	m[1] = 1001
	fmt.Println("m:", m)
}
