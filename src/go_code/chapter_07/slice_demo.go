/*
 * This is description.
 * 切片:
	1. 需求场景: 当我们用数组保存学生成绩单时, 但是学生的个数是不确定的, 这个时候我们就会需要到切片.
	2. 切片是数组的一个引用, 所以是引用类型. 进行传递, 遵循引用传递的机制.
	3. 切片的长度是可以变化的, 可以理解为一个动态的数组.
	4. 格式: var <slice_name> []<slice_type>
		eg: var students []string
	5. 切片: 底层是一个struct结构体, 如下:
		type slice struct {
			ptr *[3]int
			len int
			cap int
		}
 * @author Chris0710
 * @date 2021/1/17 11:25
*/
package main

import "fmt"

func main() {

	// 1. 切片使用的3种方式:
	// 1.1 切片: 对数组的切割.
	var intArr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// intArr1[3] = 4, intArr1[7] = 8, intArr1[8] = 9
	intSlice := intArr1[3:8]
	fmt.Printf("intSlice=%v, len=%d, cap=%d\n", intSlice, len(intSlice), cap(intSlice)) // intSlice=[4 5 6 7 8], len=5, cap=7(容量值: 属于动态变化.)

	// 1.2 make创建切片. 通过此方式创建的切片对应的数组, 是由make底层进行维护的, 对外不可见, 只能通过slice进行访问.
	// 语法: var <slice_name> []<slice_type> = make([]<slice_type>, len, [cap])
	intSlice2 := make([]int, 5, 10)
	intSlice2[2] = 20
	intSlice2[3] = 30
	fmt.Printf("intSlice2=%v, len=%d, cap=%d\n", intSlice2, len(intSlice2), cap(intSlice2)) // intSlice2=[0 0 20 30 0], len=5, cap=10

	// 1.3 定义切片时, 直接就指定具体的值.
	strSlice := []string{"a", "b", "c"}
	fmt.Printf("strSlice type is %T, strSlice=%v, len=%d, cap=%d\n", strSlice, strSlice, len(strSlice), cap(strSlice))

	// 2. 切片的遍历
	// 2.1 for
	for i := 0; i < len(strSlice); i++ {
		fmt.Printf("slice[%d]=%v\n", i, strSlice[i])
	}
	fmt.Println("----------")
	for i, v := range strSlice {
		fmt.Printf("slice[%d]=%v\n", i, v)
	}

	// 3. 切片注意事项:
	// 3.1 切片初始化的时候, 不能超过len值, 但是后续可以继续动态增长.
	// var slice = arr[0:end] 简写为: var slice = arr[:end]
	// var slice = arr[start:len(arr)] 简写为: var slice = arr[start:]
	// var slice = arr[0:len(arr)] 简写为: var slice = arr[:]
	// 3.2 切片可以继续切片.
	strSlice2 := []string{"a", "b", "c", "d", "e", "f", "g"}
	strSlice3 := strSlice2[0:5]
	fmt.Println("strSlice3:", strSlice3) // a ~ e
	strSlice4 := strSlice3[1:3]
	fmt.Println("strSlice4:", strSlice4) // b ~ c

	// 3.3 append(): 内置函数, 可用于slice的追加元素.
	// append()底层: 1. 创建一个新的数组; 2. 将原数组的值拷贝到新数组中(newArr: 是append底层维护, 程序员不可见); 3. 将slice指针指向新数组; 4. 垃圾回收原数组.
	strSlice5 := append(strSlice4, "1", "2", "3")
	fmt.Println("strSlice5:", strSlice5)

	// 3.4 copy(): 内置函数, 拷贝前后的切片是相互独立的两个数据空间, 互不影响.
	var intSlice4 = []int{1, 2, 3, 4}
	var intSlice5 = make([]int, 10, 10)
	var intSlice6 = make([]int, 1)
	copy(intSlice5, intSlice4)
	copy(intSlice6, intSlice4)
	fmt.Println("intSlice4:", intSlice4) // intSlice4: [1 2 3 4]
	fmt.Println("intSlice5:", intSlice5) // intSlice5: [1 2 3 4 0 0 0 0 0 0]
	fmt.Println("intSlice6:", intSlice6) // intSlice6: [1]
	fmt.Println("-----------------")
	intSlice4[0] = 100
	fmt.Println("intSlice4:", intSlice4) // intSlice4: [100 2 3 4]
	fmt.Println("intSlice5:", intSlice5) // intSlice5: [1 2 3 4 0 0 0 0 0 0]
	fmt.Println("intSlice6:", intSlice6) // intSlice6: [1]

	// 3.5 切片是引用类型, 所以在进行传递时, 遵循引用传递机制.
	var intSlice7 = []int{1, 2, 3, 4, 5, 6, 7}
	intSlice8 := intSlice7[:]
	intSlice9 := intSlice8[0:4]
	intSlice9[0] = 100
	fmt.Println("intSlice7:", intSlice7) // intSlice7: [100 2 3 4 5 6 7]
	fmt.Println("intSlice8:", intSlice8) // intSlice8: [100 2 3 4 5 6 7]
	fmt.Println("intSlice9:", intSlice9) // intSlice9: [100 2 3 4]
	fmt.Println("----------------")

	// 4. string 和 slice
	// 4.1 string底层是一个byte数组, 所以string也是可以进行slice切片操作.
	str4 := "hello, go"
	str4Slice := str4[0:5]
	fmt.Println("str4Slice:", str4Slice)
	str5Slice := str4Slice[2:]
	fmt.Println("str5Slice:", str5Slice)
	fmt.Println("----------------")
	// 4.2 修改string字符串:
	// string本身是不可变的, 不能通过str[0] = 'z'进行修改, 会出现编译报错.
	// 修改方式: string -> byte[](英文) / []rune(英文+中文, 防止乱码) -> 修改 -> 重新转成string.
	str6 := "hello"
	bytes := []byte(str6)
	bytes[0] = 'z'
	fmt.Println("str6:", str6)
	str6 = string(bytes)
	fmt.Println("str6:", str6)

	str7 := "成都"
	runes := []byte(str7)
	runes[0] = 'c'
	fmt.Println("str7:", str7)
	str7 = string(runes)
	fmt.Println("str7:", str7) // str7: c��都

	str8 := "成都"
	runes2 := []rune(str8)
	runes2[0] = 'c'
	fmt.Println("str8:", str8)
	str8 = string(runes2)
	fmt.Println("str8:", str8) // str8: c都

}
