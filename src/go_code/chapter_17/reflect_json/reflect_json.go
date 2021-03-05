/*
 * This is description.
 * reflect发射引出场景: 利用json tag标签的序列化引出反射.
	场景1:
	* 不确定接口调用的那个函数, 可根据传入参数在运行时, 确定调用的具体接口, 这种需要对函数或方法进行反射. 如: 下面这种桥接模式:
		func bridge(funcPtr interface{}, args ...interface{}){
			// 具体逻辑
		}
	场景2:
	* 对结构体序列化时, 如果结构体有指定的tag时, 也会使用到反射生成对应的tag字符串.
		如: 本案例中的Monster:
		type Monster struct {
			Name   string  `json:"monsterName"`
			Age    int     `json:"monsterAge"`
			Salary float64 `json:"monsterSalary"`
			Sex    string  `json:"monsterSex"`
		}
 * @author Chris0710
 * @date 2021/3/5 15:21
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name   string  `json:"monsterName"`
	Age    int     `json:"monsterAge"`
	Salary float64 `json:"monsterSalary"`
	Sex    string  `json:"monsterSex"`
}

func main() {
	monster := Monster{
		Name:   "玉兔精",
		Age:    200,
		Salary: 300.00,
		Sex:    "female",
	}
	m, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json m err: ", err)
	}
	data := string(m)
	fmt.Println("data: ", data)
	// data:  {"monsterName":"玉兔精","monsterAge":200,"monsterSalary":300,"monsterSex":"female"}
}
