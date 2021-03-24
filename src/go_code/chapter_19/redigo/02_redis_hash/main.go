/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/22 15:12
 */
package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		// handle error
		fmt.Println("redis.Dial err:", err)
		return
	}
	fmt.Println("Redis连接成功.")
	defer c.Close()

	_, err = c.Do("HSET", "user01", "name", "john")
	if err != nil {
		fmt.Println("HSET err:", err)
		return
	}
	_, err = c.Do("HSET", "user01", "age", 18)
	if err != nil {
		fmt.Println("HSET err:", err)
		return
	}

	name, err := redis.String(c.Do("HGET", "user01", "name"))
	if err != nil {
		fmt.Println("HGET err:", err)
		return
	}

	age, err := redis.Int(c.Do("HGET", "user01", "age"))
	if err != nil {
		fmt.Println("HGET err:", err)
		return
	}

	fmt.Printf("name =%#v, age=%#v\n", name, age)
	/*
		Redis连接成功.
		name ="john", age=18
	*/

	fmt.Println("--------------------")

	_, err = c.Do("HMSET", "user02", "name", "john2", "age", 20)
	if err != nil {
		fmt.Println("HMSET err:", err)
		return
	}
	res, err := redis.Strings(c.Do("HMGET", "user02", "name", "age"))
	if err != nil {
		fmt.Println("HMGET err:", err)
		return
	}
	for i, v := range res {
		fmt.Printf("user02[%#v]=[%#v]\n", i, v)
	}
	/*
	   user02[0]=["john2"]
	   user02[1]=["20"]
	*/
}
