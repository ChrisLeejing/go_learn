/*
 * This is description.
 * 需求: 在3.txt文件中, 追加5句"hello, go"
 * @author Chris0710
 * @date 2021/2/15 14:19
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\write_demo3\\3.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file 3.txt err:", err)
		return
	}
	defer file.Close()
	str := "hello, go\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("write file 3.txt err: ", err)
		}
	}

	writer.Flush()
	fmt.Println("write file 3.txt success.")
}
