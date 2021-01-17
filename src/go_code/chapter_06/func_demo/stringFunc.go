/**
 * This is description.
 * 字符串常用的系统函数.
 * @author Chris0710
 * @date 2021/1/13 10:37
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1. 按字节数统计字符串长度: len(), 注: golang使用utf-8编码, 一个汉字占用3个字节.
	var str string = "hello北"
	fmt.Println("str length is ", len(str)) // str length is  8
	fmt.Println("----------")

	// 2. 字符串遍历, 同时处理中文问题. str := []rune(str)
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c ", str2[i]) // h e l l o 北
	}
	fmt.Println()
	fmt.Println("----------")

	// 3. 字符串转整数
	str3 := "123"
	res3, err := strconv.Atoi(str3)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("res3 =%v, type is %T\n", res3, res3) // res3 =123, type is int
	}
	fmt.Println("----------")

	// 4. 整数转字符串
	n4 := 100
	res4 := strconv.Itoa(n4)
	fmt.Printf("res4 %v, type is %T\n", res4, res4) // res4 100, type is string
	fmt.Println("----------")

	// 5. 字符串转byte数组, []byte. var bytes = []byte(str)
	str5 := "hello, go"
	bytes := []byte(str5)
	fmt.Printf("bytes =%v, type is %T\n", bytes, bytes) // bytes =[104 101 108 108 111 44 32 103 111], type is []uint8
	fmt.Println("----------")

	// 6. []byte数组转字符串
	bytes6 := []byte{97, 98, 99}
	str6 := string(bytes6)
	fmt.Printf("str6 =%v, type is %T\n", str6, str6) // str6 =abc, type is string
	fmt.Println("----------")

	// 7. 十进制数转二进制, 八进制, 十六进制数
	n7 := 108
	formatInt2 := strconv.FormatInt(int64(n7), 2)
	fmt.Println("formatInt2 =", formatInt2) // formatInt2 = 1101100
	formatInt8 := strconv.FormatInt(int64(n7), 8)
	fmt.Println("formatInt8 =", formatInt8) // formatInt8 = 154
	formatInt16 := strconv.FormatInt(int64(n7), 16)
	fmt.Println("formatInt16 =", formatInt16) // formatInt16 = 6c
	fmt.Println("----------")

	// 8. 查看字符串是否包含某个子字符串.
	str8 := "hello"
	fmt.Println("是否包含子字符串ll: ", strings.Contains(str8, "ll")) // 是否包含子字符串ll:  true
	fmt.Println("是否包含子字符串oo: ", strings.Contains(str8, "oo")) // 是否包含子字符串oo:  false
	fmt.Println("----------")

	// 9. 统计字符串有多少个子字符串.
	str9 := "hello"
	fmt.Println("hello 中含有ll的个数: ", strings.Count(str9, "ll")) // hello 中含有ll的个数:  1
	fmt.Println("hello 中含有l的个数: ", strings.Count(str9, "l"))   // hello 中含有l的个数:  2
	fmt.Println("hello 中含有oo的个数: ", strings.Count(str9, "oo")) // hello 中含有oo的个数:  0
	fmt.Println("----------")

	// 10. 不区分大小写, 比较字符串是否相等.
	str10 := "hello"
	string10 := "Hello"
	fmt.Println(strings.EqualFold(str10, string10)) // true
	fmt.Println(str10 == string10)                  // false
	fmt.Println("----------")

	// 11. 返回子字符串在字符串中第一次出现的index.
	str11 := "hello"
	fmt.Println("l第一次出现的索引: ", strings.Index(str11, "l")) // l第一次出现的索引:  2
	fmt.Println("----------")

	// 12. 返回子字符串在字符串中最后一次出现的index.
	str12 := "hello"
	fmt.Println("l最后一次出现的索引: ", strings.LastIndex(str12, "l")) // l最后一次出现的索引:  3
	fmt.Println("----------")

	// 13. 将指定的子字符串替换为另一个子字符串. 当n=-1时, 全部替换, 等价于strings.ReplaceAll()
	str13 := "go go go hello"
	string13 := strings.Replace(str13, "go", "golang", 2)
	fmt.Println("string13: ", string13) // string13:  golang golang go hello
	fmt.Println("----------")

	// 14. 分割字符串
	str14 := "hello, go, hello, world, hello, kubernetes"
	strArr := strings.Split(str14, ", ") // strArr:  [hello go hello world hello kubernetes]
	fmt.Println("strArr: ", strArr)
	fmt.Println("----------")

	// 15. 大小写转换
	str15 := "Hello, world"
	lower := strings.ToLower(str15)
	upper := strings.ToUpper(str15)
	fmt.Println("lower: ", lower) // lower:  hello, world
	fmt.Println("upper: ", upper) // upper:  HELLO, WORLD
	fmt.Println("----------")

	// 16. 字符串去两边空格
	str16 := "          Hello, go  "
	string16 := strings.TrimSpace(str16)
	fmt.Println("string16: ", string16) // string16:  Hello, go
	fmt.Println("----------")

	// 17. 字符串去两边指定的字符
	str17 := "           !!!$$Hello, go  $$  !!!   !!!"
	string17 := strings.Trim(str17, " !")
	fmt.Println("string17: ", string17) // string17:  $$Hello, go  $$
	fmt.Println("----------")

	// 18. 字符串去左边指定的字符
	str18 := "           !!!$$Hello, go  $$  !!!   !!!"
	string18 := strings.TrimLeft(str18, "! ")
	fmt.Println("string18: ", string18) // string18:  $$Hello, go  $$  !!!   !!!
	fmt.Println("----------")

	// 19. 字符串去右边指定的字符
	str19 := "           !!!$$Hello, go  $$  !!!   !!!"
	string19 := strings.TrimRight(str19, "! ")
	fmt.Println("string19: ", string19) // string19:             !!!$$Hello, go  $$
	fmt.Println("----------")

	// 20. 判断字符串是否匹配前缀
	str20 := "http://www.hunlish.com"
	fmt.Println("字符串是否以http为前缀: ", strings.HasPrefix(str20, "http")) // 字符串是否以http为前缀:  true
	fmt.Println("----------")

	// 21. 判断字符串是否匹配后缀
	str21 := "http://www.hunlish.com"
	fmt.Println("字符串是否以com为后缀: ", strings.HasSuffix(str21, "com")) // 字符串是否以com为后缀:  true
	fmt.Println("----------")

}
