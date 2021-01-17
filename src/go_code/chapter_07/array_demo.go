/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/16 22:11
 */
package main

import "fmt"

// 值传递: 长度是数组的一部分, 形参需要加长度, 如: [3]int
func test(arr [3]int) {
	arr[0] = 100
}

// 指针传递: 长度是数组的一部分, 形参需要加长度, 如: *[3]int
func test2(arr *[3]int) {
	(*arr)[0] = 100
}

func main() {
	// 1. 数组的地址值.
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("arr:", arr) // arr: [10 20 30]
	// arr的地址: 0xc00009e120, arr[0]的地址: 0xc00009e120, arr[1]的地址: 0xc00009e128, arr[2]的地址: 0xc00009e130
	fmt.Printf("arr的地址: %p, arr[0]的地址: %p, arr[1]的地址: %p, arr[2]的地址: %p\n", &arr, &arr[0], &arr[1], &arr[2])

	// 2. 数组的4中初始化方式.
	// 2.1 固定容量赋值
	var intArr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("intArr1:", intArr1)
	var intArray1 = [4]int{1, 2, 3, 4}
	fmt.Println("intArray1:", intArray1)

	// 2.2 可变容量赋值
	var intArr2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("intArr2:", intArr2)

	// 2.3 index-value方式
	var intArr3 = [...]int{2: 200, 1: 100, 0: 0}
	fmt.Println("intArr3:", intArr3) // intArr3: [0 100 200]

	// 2.4 类型推导
	strArr4 := [...]string{1: "go", 2: "java", 3: "kubernetes", 0: "Istio"}
	fmt.Println("strArr4:", strArr4) // [Istio go java kubernetes]
	fmt.Println("1" + strArr4[0] + "2")

	// 3. for-range: array
	books := [...]string{"红楼梦", "三国演义", "水浒传"}

	for i, book := range books {
		fmt.Println("i:", i, "book:", book)
		fmt.Printf("books[%d]=%v\n", i, book)
	}
	for _, book := range books {
		fmt.Println("book:", book)
	}

	// 4. 数组默认是值传递, 进行值拷贝, 数组之间互不影响.
	// 4.1 默认值传递
	intArr4 := [3]int{1, 2, 3}
	test(intArr4)
	fmt.Println("test() intArr4:", intArr4)
	// 4.2 指针引用传递
	test2(&intArr4)
	fmt.Println("test2() intArr4:", intArr4)

}
