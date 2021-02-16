/*
 * This is description.
 * 需求: 创建1.txt文件, 并向其中写入5行"hello, golang."
 * @author Chris0710
 * @date 2021/2/15 12:03
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// func OpenFile(name string, flag int, perm FileMode) (filePath *File, err error)
	filePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\write_demo\\2.txt"
	// 0666: r=4, w=2, x=1, 0666: rw_rw_rw_, 即: 所有人都拥有读写操作.
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open filePath err: ", err)
		return
	}
	defer file.Close()

	str := "hello, golang.\n"
	// 使用带缓存*Writer进行写入, 默认值: defaultBufSize = 4096
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		// writer.WriteString: return length, err.
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("write file file.txt err: ", err)
		}

	}
	// bufio.NewWriter写入文件时, 使用缓存写入到内存中, 需要对writer进行刷新, 才能将缓存中的数据真正的写入磁盘. 否则, 磁盘中可能没有数据.
	writer.Flush()
	fmt.Println("write file file.txt success. ")

}
