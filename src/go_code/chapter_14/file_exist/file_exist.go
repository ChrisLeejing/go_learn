/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/15 15:07
 */
package main

import (
	"fmt"
	"os"
)

func PathExists(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if err == nil { // 没有错误, 说明文件存在
		fmt.Println("fileInfo: ", fileInfo)
		return true, nil
	}
	if os.IsNotExist(err) { // 有错误, 说明文件不存在
		return false, nil
	}
	return false, err
}

func main() {
	filePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_exist\\file.txt"
	exists, err := PathExists(filePath)
	fmt.Printf("exists: %v\n err: %v", exists, err)
}
