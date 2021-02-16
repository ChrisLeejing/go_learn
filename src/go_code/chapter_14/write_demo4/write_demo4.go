/*
 * This is description.
 * 需求: 打开已存在的4.txt文件, 将原因的文件读出并显示终端上, 并且追加5行"hello, go"到文件中.
 * @author Chris0710
 * @date 2021/2/15 14:28
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\write_demo4\\4.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(content)
	}

	str := "hello, go\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("write file 4.txt err: ", err)
		}
	}

	writer.Flush()
	fmt.Println("write file 4.txt success. ")
}
