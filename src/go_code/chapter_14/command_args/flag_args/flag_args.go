/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/16 14:44
 */
package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	// &user: 接收用户命令行中输入的 -u 后面的参数值
	// "u": 就是 -u 指定参数
	// "": 默认值
	// "用户名,默认为空": 备注说明
	flag.StringVar(&user, "u", "", "用户名, 默认为空.")
	flag.StringVar(&pwd, "pwd", "", "密码, 默认为空.")
	flag.StringVar(&host, "h", "localhost", "主机名, 默认localhost.")
	flag.IntVar(&port, "port", 3306, "端口号, 默认: 3306.")

	// 转换操作.
	flag.Parse()

	fmt.Printf("user: %v, pwd: %v, host: %v, port: %v", user, pwd, host, port)
	/*
	   E:\Code\Go\go_learn\src\go_code\chapter_14\command_args\flag_args>flag_args.exe -u root -pwd 123456 -h 127.0.0.1
	   user: root, pwd: 123456, host: 127.0.0.1, port: 3306
	*/
}
