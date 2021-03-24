/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/22 16:18
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

	// LPUSH mylist a b c
	_, err = c.Do("LPUSH", "mylist", "a", "b", "c")
	if err != nil {
		fmt.Println("LPUSH err:", err)
		return
	}
	res, err := redis.Strings(c.Do("LRANGE", "mylist", 0, -1))
	if err != nil {
		fmt.Println("LRANGE err:", err)
		return
	}

	fmt.Println(res)
}
