/*
 * This is description.
 * 参考:
	* Redis connection pool with golang: https://medium.com/rahasak/redis-connection-pool-with-golang-ee75f9fcf0e0
	* golang开发中 redis连接池的使用: https://studygolang.com/articles/26061
 * @author Chris0710
 * @date 2021/3/22 16:33
*/
package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数.
		MaxActive:   0,   // 与Redis的最大连接数, 0: 连接没有限制.
		IdleTimeout: 100, // 最大空闲时间.
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379") // 初始化连接哪个IP的Redis.
		},
	}
}

func main() {
	// 取出一个连接.
	conn1 := pool.Get()
	defer conn1.Close()

	_, err := conn1.Do("SET", "redis_pool01", "01")
	if err != nil {
		fmt.Println("redis pool01 SET err:", err)
		return
	}
	s, err := redis.String(conn1.Do("GET", "redis_pool01"))
	if err != nil {
		fmt.Println("redis pool01 GET err:", err)
		return
	}
	fmt.Println("redis_pool01:", s)

	fmt.Println("----------------")

	// pool.Close() // 如果需要继续获取连接, 则连接池不能关闭. 否则: redis pool02 SET err: redigo: get on closed pool.
	conn2 := pool.Get()
	_, err = conn2.Do("SET", "redis_pool02", "02")
	if err != nil {
		fmt.Println("redis pool02 SET err:", err)
		return
	}
	s2, err := redis.String(conn2.Do("GET", "redis_pool02"))
	if err != nil {
		fmt.Println("redis pool02 GET err:", err)
		return
	}
	fmt.Println("redis_pool02:", s2)
}
