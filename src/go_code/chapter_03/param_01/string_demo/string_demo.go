/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/6 9:11
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name = "chris\nlee"
	fmt.Println(name)
	fmt.Println("-------------")
	var str = `
package main

import "fmt"

func main()  {
	fmt.Println("Hello, world!")

}
`
	fmt.Printf(str)
	fmt.Println("-------------")

	var str2 = "hello" + "world"
	str2 += "!!!"
	fmt.Println("str2 =", str2)

	fmt.Println("-------------")

	var str3 = "hello" + "world" + "hello" +
		"world" + "hello" +
		"world" + "hello" + "world" +
		"hello" + "world"
	fmt.Println("str3 =", str3)

	fmt.Println("-------------")

	var a int
	var b int32
	var c float32
	var d float64
	var isMarried bool
	var str4 string
	fmt.Printf("a=%v, b=%v, c=%v, d=%v, isMarried=%v, str4=%v\n", a, b, c, d, isMarried, str4)

	fmt.Println("-------------")

	var n int = 29
	var n1 float32 = float32(n)
	var n2 float64 = float64(n)
	var n3 int8 = int8(n)
	fmt.Printf("n1=%v, %T, \nn2=%v, %T, \nn3=%v, %T\n", n1, n1, n2, n2, n3, n3)

	fmt.Println("-------------")

	var n4 int = 1000
	var n5 int8 = int8(n4)
	fmt.Printf("n5=%v, %T\n", n5, n5) // n5=-24, int8. the number(1000) will be calculated with overflowed, because int8 is [-128, 127].

	fmt.Println("-------------")

	// 基本数据类型转string类型.
	// case 1: fmt.Sprintf(format string, a ...interface{}) string {...}
	var a2 int = 1
	var b2 int32 = 2
	var c2 float32 = 2.3
	var d2 float64 = 1.4
	var isMarried2 bool = true
	var myChar byte = 'a'
	var str6 string
	str6 = fmt.Sprintf("%d", a2)
	fmt.Printf("a2=%T, %q\n", str6, str6)

	str6 = fmt.Sprintf("%d", b2)
	fmt.Printf("b2=%T, %q\n", str6, str6)

	str6 = fmt.Sprintf("%f", c2)
	fmt.Printf("c2=%T, %q\n", str6, str6)

	str6 = fmt.Sprintf("%f", d2)
	fmt.Printf("d2=%T, %q\n", str6, str6)

	str6 = fmt.Sprintf("%t", isMarried2)
	fmt.Printf("isMarried2=%T, %q\n", str6, str6)

	str6 = fmt.Sprintf("%c", myChar)
	fmt.Printf("myChar=%T, %q\n", str6, str6)

	fmt.Println("-------------")

	// case 2: strconv
	// func FormatBool(b bool) string
	// func FormatComplex(c complex128, fmt byte, prec, bitSize int) string
	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// func FormatInt(i int64, base int) string
	// func FormatUint(i uint64, base int) string
	// func Itoa(i int) string
	var num3 int = 200
	var num4 float64 = 3.1415
	var num5 int64 = 12345
	var b3 bool = true
	num3Str := strconv.FormatInt(int64(num3), 10)
	num3Str16 := strconv.FormatInt(int64(num3), 16)
	fmt.Println("num3Str:", num3Str)
	fmt.Println("num3Str16:", num3Str16) // num3Str16: c8 = 12 * 16 + 8 = 200
	fmt.Printf("num3Str type is %T, num3Str =%q\n", num3Str, num3Str)

	// num4Str := strconv.FormatFloat(num4, 'f', -1, 64)
	// fmt.Printf("num4Str type is %T, num4Str =%q\n", num4Str, num4Str) // num4Str type is string, num4Str ="3.1415"

	// num4Str := strconv.FormatFloat(num4, 'f', 3, 64)
	// fmt.Printf("num4Str type is %T, num4Str =%q\n", num4Str, num4Str) // num4Str type is string, num4Str ="3.142"

	num4Str := strconv.FormatFloat(num4, 'f', 2, 64)
	fmt.Printf("num4Str type is %T, num4Str =%q\n", num4Str, num4Str) // num4Str type is string, num4Str ="3.14"

	b3Str := strconv.FormatBool(b3)
	fmt.Printf("b3Str type is %T, b3Str =%q\n", b3Str, b3Str) // b3Str type is string, b3Str ="true"

	fmt.Println("-------------")

	num5Str := strconv.Itoa(int(num5))
	fmt.Printf("num5Str type is %T, num5Str =%q\n", num5Str, num5Str) // num5Str type is string, num5Str ="12345"

	// string类型转基本数据类型.
	// case 1: strconv
	// func ParseBool(str string) (bool, error)
	// func ParseComplex(s string, bitSize int) (complex128, error)
	// func ParseFloat(s string, bitSize int) (float64, error)
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// func ParseUint(s string, base int, bitSize int) (uint64, error)
	var str5 string = "true"
	var b5 bool
	b5, err := strconv.ParseBool(str5)
	fmt.Printf("b5 type is %T, value is %v, error is %v\n", b5, b5, err) // b5 type is bool, value is true, error is <nil>
	// b5, _ = strconv.ParseBool(str5)
	// fmt.Printf("b5 type is %T, value is %v\n", b5, b5) // b5 type is bool, value is true
	fmt.Println("-------------")
	var str7 string = "1234678"
	var i6 int
	i6Int64, _ := strconv.ParseInt(str7, 10, 64)
	i6 = int(i6Int64)
	fmt.Printf("i6Int64 type is %T, value is %v\n", i6Int64, i6Int64)
	fmt.Printf("i6 type is %T, value is %v\n", i6, i6)

	fmt.Println("-------------")

	var str8 string = "123.478"
	var f8 float64
	f8, _ = strconv.ParseFloat(str8, 64)
	fmt.Printf("f8 type is %T, value is %v\n", f8, f8)

	fmt.Println("-------------")

	// 需要确保string转换的是有效的数据类型, 否则直接转换为0进行处理.
	// eg: "hello" -> 0
	var str9 string = "hello"
	var i9 int64
	i9, err9 := strconv.ParseInt(str9, 10, 64)
	// i9 type is int64, value is 0, err isstrconv.ParseInt: parsing "hello": invalid syntax
	fmt.Printf("i9 type is %T, value is %v, err is%v\n", i9, i9, err9)

}
