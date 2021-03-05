/*
 * This is description.
 * 需求:
	1. 请编写一个案例, 用于演示(基本数据类型, interface{}, reflect.Value)进行反射的基本操作.
	2. 请编写一个案例, 用于演示(结构体类型, interface{}, reflect.Value)进行反射的基本操作.
 * @author Chris0710
 * @date 2021/3/5 16:39
*/
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	// 需求1:
	var num = 1
	ReflectTest01(num)
	fmt.Println("-----------------")

	// 需求2:
	s := Student{
		Name: "chris",
		Age:  30,
	}
	m := Monster{
		Name: "牛魔王",
		Age:  300,
	}

	ReflectTest02(s)
	fmt.Println("-----------------")
	ReflectTest02(m)
}

func ReflectTest01(i interface{}) {
	// 通过反射获取传入变量的type, kind, value.
	// 1. 获取到reflect.Type.
	rType := reflect.TypeOf(i)
	fmt.Println("rType: ", rType) // rType:  int

	// 2. 获取到reflect.Value.
	rValue := reflect.ValueOf(i)
	n2 := rValue.Int() + 2
	fmt.Println("n2: ", n2)
	fmt.Printf("rValue Type: %T, rValue value: %v\n", rValue, rValue)
	/*
		n2:  3
		rValue Type: reflect.Value, rValue value: 1
	*/

	// 3. 将reflect.Value转换为interface{}
	num2 := rValue.Interface()
	// 4. 将interface{}类型通过具体的类型断言, 转换成具体的类型.
	res2 := num2.(int)
	fmt.Printf("res2 Type: %T, res2 value: %v\n", res2, res2) // res2 Type: int, res2 value: 1
}

func ReflectTest02(i interface{}) {
	// 1. 获取到reflect.Type.
	rType := reflect.TypeOf(i)
	fmt.Println("rType: ", rType)

	// 2. 获取到reflect.Value.
	rValue := reflect.ValueOf(i)
	fmt.Printf("rValue Type: %T, rValue value: %v\n", rValue, rValue)

	// 3. 获取变量对应的kind.
	tKind := rType.Kind()
	vKind := rValue.Kind()
	fmt.Printf("tKind: %v, vKind: %v\n", tKind, vKind)

	// 4. 将reflect.Value转换为interface{}
	v1 := rValue.Interface()
	fmt.Printf("v1 Type: %T, v1 value: %v\n", v1, v1)
	s, ok := v1.(Student)
	if ok {
		fmt.Printf("s.Name = %v\n", s.Name)
	}
}
