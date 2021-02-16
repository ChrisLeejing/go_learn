/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/16 15:16
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
}

func testMarshalStruct(user User) []byte {

	data, err := json.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
func testMarshalMap() {
	// 定义一个map
	var user map[string]interface{}
	// 使用map,需要make
	user = make(map[string]interface{})
	user["name"] = "chris"
	user["age"] = 30
	user["address"] = "chengdu"

	// 将user这个map进行序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("user map 序列化后=%v\n", string(data))

}

// 演示对切片进行序列化, 切片: []map[string]interface{}
func testMarshalSlice() {
	var slice []map[string]interface{}
	var user1 map[string]interface{}
	// 使用map前, 需要先make
	user1 = make(map[string]interface{})
	user1["name"] = "jack"
	user1["age"] = "7"
	user1["address"] = "北京"
	slice = append(slice, user1)

	var user2 map[string]interface{}
	// 使用map前, 需要通过make方法分配确定的内存地址
	user2 = make(map[string]interface{})
	user2["name"] = "tom"
	user2["age"] = "20"
	user2["address"] = [2]string{"成都", "硅谷"}
	slice = append(slice, user2)

	// 将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

// 对基本数据类型序列化, 对基本数据类型进行序列化意义不大
func testMarshalFloat64() {
	var num1 float64 = 2345.67
	// 对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func testUnmarshalStruct(str string) {
	var user User
	data := []byte(str)
	err := json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user: ", user)
}

// 演示将json字符串，反序列化成map
func testUnmarshalMap() {
	str := "{\"address\":\"chengdu\",\"age\":30,\"name\":\"chris\"}"
	// 定义一个map
	var user map[string]interface{}

	// 注意：反序列化map, 不需要make, 因为make分配内存操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &user)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 user=%v\n", user)
}

// 演示将json字符串，反序列化成切片
func testUnmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"成都\",\"杭州\"],\"age\":\"20\",\"name\":\"tom\"}]"

	// 定义一个slice
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	user := User{
		Name: "tom",
		Age:  18,
	}
	fmt.Println("-----------序列化-----------")
	data := testMarshalStruct(user)
	fmt.Println("data: ", string(data))

	testMarshalMap()
	testMarshalSlice()
	testMarshalFloat64()

	fmt.Println("-----------反序列化-----------")

	str := "{\"Name\":\"tom\",\"Age\":18}"
	testUnmarshalStruct(str)
	testUnmarshalMap()
	testUnmarshalSlice()
}
