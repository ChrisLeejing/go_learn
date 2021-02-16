/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/15 9:16
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. 打开和关闭file文件(file对象, file指针, file句柄)
	file, err := os.Open("E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_demo\\test.txt")
	if err != nil {
		fmt.Println("open the file error: ", err)
	}
	fmt.Println("file: ", file) // file:  &{0xc0000cc780}

	err = file.Close()
	if err != nil {
		fmt.Println("close the file error: ", err)
	}

}
