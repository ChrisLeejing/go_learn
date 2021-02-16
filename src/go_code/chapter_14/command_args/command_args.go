/*
 * This is description.
 * 需求: 获取到命令行输入的所有参数.
 * @author Chris0710
 * @date 2021/2/16 14:31
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("命令行的参数个数: %v\n", len(args))
	for i, v := range args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	/*
		E:\Code\Go\go_learn\src\go_code\chapter_14\command_args>command_args.exe chris 30 true
		命令行的参数个数: 4
		args[0]=command_args.exe
		args[1]=chris
		args[2]=30
		args[3]=true
	*/
}
