/*
 * This is description.
 * 需求:
	1. 使用反射来遍历Struct结构体的字段, 调用结构体的方法, 并获取结构体标签的值.
	2. 使用反射的方式来获取结构体tag的标签, 遍历字段的值, 修改字段的值, 调用结构体方法(要求: 使用传递地址的方式完成.)
   参考: Golang的反射reflect深入理解和示例: https://studygolang.com/articles/12348
 * @author Chris0710
 * @date 2021/3/8 17:26
*/
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"stuName"`
	Age   int     `json:"stuAge"`
	Score float32 `json:"分数"`
	Sex   string  `json:"stuSex"`
}

func main() {
	s := Student{
		Name:  "chris",
		Age:   12,
		Score: 99.5,
		Sex:   "男",
	}
	s2 := Student{
		Name:  "mary",
		Age:   18,
		Score: 80.5,
		Sex:   "女",
	}

	TestStruct(s)
	fmt.Println("============================================================")
	TestStruct2(&s2)

}

// 方法1: 返回两个数的和.
func (s Student) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 方法2: 接收4个值, 赋值给Student.
func (s Student) SetStudent(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
	fmt.Println("调用SetStudent(name string, age int, score float32, sex string)")
}

// 方法3: 打印student的值.
func (s Student) Print() {
	fmt.Println("start........")
	fmt.Printf("student: %v\n", s)
	fmt.Println("end..........")
}

// 设置新的名字: name
func (s *Student) SetName(name string) string {
	oldName := s.Name
	s.Name = name
	return oldName
}

// 设置新的年龄: age+1
func (s *Student) AddAge() bool {
	s.Age++
	return true
}

// 反射遍历Struct.
func TestStruct(i interface{}) {
	// 获取reflect.Type类型.
	iType := reflect.TypeOf(i)
	fmt.Printf("iType: %v\n", iType) // iType: main.Student
	// 获取reflect.Value类型.
	iValue := reflect.ValueOf(i)
	fmt.Printf("iValue: %v\n", iValue) // iValue: {chris 12 99.5 男}
	// 获取reflect.Value.Kind类别.
	iKind := iValue.Kind()
	fmt.Printf("iKind: %v\n", iKind) // iKind: struct

	// 如果传入的不是struct kind, 直接退出.
	if iKind != reflect.Struct {
		fmt.Println("expect struct kind. ")
		return
	}
	numField := iValue.NumField()
	fmt.Printf("numField: %v\n", numField)

	for i := 0; i < numField; i++ {
		fmt.Printf("Field num: %d, Field value: %v\n", i, iValue.Field(i))
		tagValue := iType.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("Field num: %d, Field value: %v, tagValue: %v\n", i, iValue.Field(i), tagValue)
		}
	}
	fmt.Println("------------------------------")
	numMethod := iValue.NumMethod()
	for i := 0; i < numMethod; i++ {
		// 方法的排序默认: 按照函数名的ASCII码进行排序.
		method := iType.Method(i)
		fmt.Printf("method index: %d, method name: %s\n", i, method.Name)
		if i == 0 {
			// 将调用方法: GetSum(n1, n2 int)
			// iValue.Method(i).Call([]int{100, 200}): 此错误写法.
			// 正确方式如下:
			var params []reflect.Value
			params = append(params, reflect.ValueOf(100))
			params = append(params, reflect.ValueOf(200))
			values := iValue.Method(i).Call(params)
			fmt.Printf("反射调取method name: %s, 结果: %d\n", method.Name, values[0].Int())
		} else if i == 1 {
			// 将调用方法: Print()
			iValue.Method(i).Call(nil)
		} else if i == 2 {
			// 将调用方法: SetStudent(name string, age int, score float32, sex string)
			var params []reflect.Value
			params = append(params, reflect.ValueOf("jack"))
			params = append(params, reflect.ValueOf(20))
			params = append(params, reflect.ValueOf(float32(99.5)))
			params = append(params, reflect.ValueOf("男"))
			iValue.Method(i).Call(params)
		}
		fmt.Println("------------------------------")
	}
}

// 需求2: 使用反射的方式来获取结构体tag的标签, 遍历字段的值, 修改字段的值, 调用结构体方法(要求: 使用传递地址的方式完成.)
func TestStruct2(i interface{}) {
	var setNameStr = "SetName"
	var addAgeStr = "AddAge"

	// 1. 获取到结构体类型变量的反射类型.
	iValue := reflect.ValueOf(i)
	fmt.Printf("修改前iValue: %v\n", iValue)
	// 2. 获取确切的方法名.
	// 带参数函数: func (s *Student) SetName(name string) string {}
	methodName := iValue.MethodByName(setNameStr)
	newName := []reflect.Value{reflect.ValueOf("Mary2")}
	oldName := methodName.Call(newName)
	fmt.Printf("oldName: %v\n", oldName)
	fmt.Printf("修改name后, iValue: %v\n", iValue)

	// 不带参数函数: func (s *Student) AddAge() bool {}
	methodAge := iValue.MethodByName(addAgeStr)
	methodAge.Call(nil) // Call()内部有in = make([]Value, n+1)初始化.
	// methodAge.Call(make([]reflect.Value, 0))
	fmt.Printf("修改age后, iValue: %v\n", iValue)
	/*
	   修改前iValue: &{mary 18 80.5 女}
	   oldName: [mary]
	   修改name后, iValue: &{Mary2 18 80.5 女}
	   修改age后, iValue: &{Mary2 19 80.5 女}
	*/
}
