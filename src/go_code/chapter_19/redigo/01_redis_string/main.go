/*
 * This is description.
 * 参考: https://pkg.go.dev/github.com/gomodule/redigo/redis
 * @author Chris0710
 * @date 2021/3/21 16:20
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

	_, err = c.Do("SET", "hello", "world")
	if err != nil {
		fmt.Println("Redis SET err:", err)
		return
	}
	s, err := redis.String(c.Do("GET", "hello"))
	if err != nil {
		fmt.Println("Redis GET err:", err)
		return
	}
	fmt.Printf("%#v\n", s)
	/*
		Redis连接成功.
		"world"
	*/

	fmt.Println("-------------------------")

	// 批量操作: SET/GET
	// MSET date "2012.3.30" time "11:00 a.m." weather "sunny"
	_, err = c.Do("MSET", "date", "2012.3.30", "time", "11:00 a.m.", "weather", "sunny")
	if err != nil {
		fmt.Println("MSET err:", err)
		return
	}

	res, err := redis.Strings(c.Do("MGET", "date", "time", "weather"))
	if err != nil {
		fmt.Println("MGET err:", err)
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}

	// 过期时间: EXPIRE
	// # 设置过期时间为 10 秒
	_, err = c.Do("EXPIRE", "hello", 10)
	if err != nil {
		fmt.Println("EXPIRE err:", err)
		return
	}

}
