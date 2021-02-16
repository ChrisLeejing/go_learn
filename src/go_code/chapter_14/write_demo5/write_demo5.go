/*
 * This is description.
 * 需求: 将5-1.txt的文件写入到5-2.txt文件中, 注: 5-1.txt和5-2.txt都已存在.
   使用ioutil.ReadFile和ioutil.WriteFile进行读写文件.
 * @author Chris0710
 * @date 2021/2/15 14:44
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath1 := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\write_demo5\\5-file.txt"
	filePath2 := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\write_demo5\\5-2.txt"
	fileBytes1, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("read file 5-file.txt err: ", err)
		return
	}
	fmt.Println(string(fileBytes1))

	// If the file does not exist, WriteFile creates it with permissions perm
	// (before umask); otherwise WriteFile truncates it before writing.
	err = ioutil.WriteFile(filePath2, fileBytes1, 0666)
	if err != nil {
		fmt.Println("write file err: ", err)
	}
	fmt.Println("write file 5-2.txt success. ")
}
