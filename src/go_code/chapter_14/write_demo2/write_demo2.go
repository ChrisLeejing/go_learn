/*
 * This is description.
 * 需求: 将2.txt中的文件覆盖掉, 重新写入10行"hello, world. "
 * @author Chris0710
 * @date 2021/2/15 14:03
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\write_demo2\\2.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	// 关闭file句柄.
	defer file.Close()
	str := "hello, world.\n"
	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("write file 2.txt err: ", err)
		}
	}
	writer.Flush()
	fmt.Println("write file 2.txt success.")

}
