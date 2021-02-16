/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/15 11:52
 */
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 2. 读文件操作示例.
	// 2.2 ioutil一次性将文件读取到内存中: 适合于文件较小的情况.
	// 注: 因为文件的open和close操作, 源码如下:
	// f, err := os.Open(filename)
	//	if err != nil {
	//		return nil, err
	//	}
	//	defer f.Close()
	//	被封装到ioutil.ReadFile(file)中, 所以我们不需要显示的进行open和close文件.
	file := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\file_demo\\test.txt"
	// func ReadFile(filename string) ([]byte, error)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err: ", err)
	}
	fmt.Println(content)
	fmt.Println("--------")
	fmt.Println(string(content))
	/*
	   [104 101 108 108 111 44 32 103 111 108 97 110 103 46 13 10 104 101 108 108 111 44 32 119 111 114 108 100 46 13 10]
	   --------
	   hello, golang.
	   hello, world.
	*/
}
