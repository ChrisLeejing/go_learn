/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/3/16 9:12
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 客户端开始连接.
	fmt.Println("客户端开始连接: ")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	defer conn.Close()

	if err != nil {
		// handle error
		fmt.Println("client conn err: ", err)
		return
	}
	// 终端进行输入数据.
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client readString err: ", err)
		}
		if strings.Trim(line, "\r\n") == "exit" {
			fmt.Println("客户端关闭. ")
			break
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("write msg err: ", err)
		}

	}

}
