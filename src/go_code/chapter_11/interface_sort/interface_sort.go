/*
 * This is description.
 * 需求: 将Student的切片按照Score/Name的大小降序(升序)排列.
 * @author Chris0710
 * @date 2021/2/9 17:25
 */
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1. 声明Student结构体.
type Student struct {
	Name  string
	Age   int
	Score float64
}

// 2. 声明Student结构体切片类型.
type StudentSlice []Student

// 3. 实现func Sort(data Interface)接口.
func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i, j int) bool {
	// return s[i].Score < s[j].Score // 针对于Score进行升序排列.
	return s[i].Name < s[j].Name // 针对于Name进行升序排列.
}

func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	var students StudentSlice
	for i := 0; i < 10; i++ {
		s := Student{
			// 注: fmt.Sprintf用于string与int类型拼接.
			Name:  fmt.Sprintf("chris_%d", rand.Intn(100)),
			Age:   10,
			Score: float64(rand.Intn(100)),
		}
		students = append(students, s)
	}

	for _, s := range students {
		fmt.Println("排序前: ", s)
	}

	fmt.Println("--------------")

	sort.Sort(students)

	for _, s := range students {
		fmt.Println("升序排序后: ", s)
	}

	fmt.Println("--------------")
	sort.Sort(sort.Reverse(students))

	for _, s := range students {
		fmt.Println("降序排序后: ", s)
	}

}
