/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/16 9:16
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务端开始进行监听.
	fmt.Println("服务端开始进行监听: ")

	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	defer listener.Close()

	if err != nil {
		// handle error
		fmt.Println("server listener err: ", err)
		return
	}
	for {
		// 循环接收客户端输入的值.
		conn, err := listener.Accept()
		if err != nil {
			// handle error
			fmt.Println("server accept err: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// 创建一个切片, 存储客户端发送的数据.
		bytes := make([]byte, 1024)
		// 从conn连接中, 读取数据.
		len, err := conn.Read(bytes)
		if err != nil {
			fmt.Printf("读取客户端[%s]数据出错. \n", conn.RemoteAddr().String())
			return
		}
		// 服务端接收到客户端的数据, 并显示输出.
		fmt.Printf("服务端接收到的数据: %s\n", string(bytes[:len]))
	}

}
