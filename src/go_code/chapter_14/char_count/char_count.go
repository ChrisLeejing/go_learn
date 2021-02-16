/*
 * This is description.
 * 需求: 统计一个文件中, 英文, 数字, 空格, 其他字符的个数.
 * @author Chris0710
 * @date 2021/2/16 11:52
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type CharCount struct {
	CharNum   int
	NumberNum int
	SpaceNum  int
	OtherNum  int
}

func main() {
	FilePath := "E:\\Code\\Go\\go_learn\\src\\go_code\\chapter_14\\char_count\\1.txt"
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal("open file err: ", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// []rune: 为了兼容中文字符.
		runes := []rune(content)
		for _, v := range runes {
			switch { // 此用法类似于if...else if...else
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.CharNum++
			case v >= '0' && v <= '9':
				count.NumberNum++
			case v == ' ' || v == '\t':
				count.SpaceNum++
			default:
				count.OtherNum++
			}
		}
	}

	fmt.Printf("字符的个数: %v, 数字的个数: %v, 空格的个数: %v, 其他字符个数: %v.", count.CharNum, count.NumberNum, count.SpaceNum, count.OtherNum)
}
