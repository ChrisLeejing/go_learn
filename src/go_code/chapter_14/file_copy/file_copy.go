/*
 * This is description.
 * 需求: 将E:\Code\Go\go_learn\src\go_code\chapter_14\file_copy\src\golang_logo.jpg图片拷贝到E:\Code\Go\go_learn\src\go_code\chapter_14\file_copy\dst路径下.
   使用: func Copy(dst Writer, src Reader) (written int64, err error)
 * @author Chris0710
 * @date 2021/2/15 15:30
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	src, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("open file err: ", err)
		log.Fatal(err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func CopyFile2(dstFilePath string, srcFilePath string) (written int64, err error) {
	src, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer src.Close()
	reader := bufio.NewReader(src)

	dst, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	writer := bufio.NewWriter(dst)
	writer.Flush()

	return io.Copy(writer, reader)
}

func main() {
	// srcFilePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_copy\\src\\golang_logo_src.jpg"
	// dstFilePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_copy\\dst\\golang_logo_dst.jpg"
	// _, err := CopyFile(dstFilePath, srcFilePath)

	srcFilePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_copy\\src\\golang_logo_src2.jpg"
	dstFilePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_copy\\dst\\golang_logo_dst2.jpg"
	_, err := CopyFile2(dstFilePath, srcFilePath)

	if err != nil {
		fmt.Println("copy file err: ", err)
	}

	fmt.Println("copy file success. ")
}
