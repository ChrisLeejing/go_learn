/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/15 11:06
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 2. 读文件操作示例.
	file, err := os.Open("E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_demo\\test.txt")
	if err != nil {
		fmt.Println("open the file error: ", err)
	}
	fmt.Println("file: ", file) // file:  &{0xc0000cc780}
	defer file.Close()          // 函数退出后, 及时关闭file句柄, 否则可能会出现内存泄漏.

	// 2.1 bufio缓冲方式读取文件内容:
	// NewReader returns a new Reader whose buffer has the default size.
	/*
		const (
			defaultBufSize = 4096
		)
	*/
	reader := bufio.NewReader(file)
	for {
		// delim: \n 表示读取到换行符, 则结束本次循环.
		readString, err := reader.ReadString('\n')
		// io.EOF: 表示读取到文件末尾.
		if err == io.EOF {
			break
		}
		// 输出文件读取内容.
		fmt.Println(readString)
	}
	fmt.Println("文件读取结束. ")

}
