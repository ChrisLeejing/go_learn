/*
 * This is description.
 * How to use variable in go.
 * @author Chris0710
 * @date 2021/1/5 10:06
 */
package main

import (
	"fmt"
	"unsafe"
)

// define global variable.
var age2 int = 10

// while declare the variable, assign the variable with init value, the type of variable can be omit.
var name2 = "cc"

var (
	age3         = 20
	name3 string = "dd~"
)

func main() {
	var i int
	fmt.Println("i = ", i)

	var j int
	j = 5
	fmt.Println("j = ", j)

	k := 2.0
	fmt.Println("k = ", k)

	m := 2.1
	fmt.Println("m = ", m)
	fmt.Printf("The type of m is %T, the value is %f\n", m, m)

	// n := "chris" <=> var n string; n = "chris"
	n := "chris"
	fmt.Println("n = ", n)
	/*
		i =  0
		j =  5
		k =  2
		m =  2.1
		The type of m is float64, the value is 2.100000
		n =  chris
	*/

	// multiple variable declare.
	// case 1:
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// case 2:
	var age, name, gender = 29, "chris", true
	fmt.Println("age=", age, "name=", name, "gender=", gender)

	// case 3:
	price, book := 30.09, "A Dream of Red Mansions"
	fmt.Println("price=", price, "book=", book)
	/*
		n1= 0 n2= 0 n3= 0
		age= 29 name= chris gender= true
		price= 30.09 book= A Dream of Red Mansions
	*/

	// print the global variables.
	fmt.Println("age2=", age2)
	fmt.Println("name2=", name2)
	fmt.Println("age3=", age3)
	fmt.Println("name3=", name3)
	/*
		age2= 10
		name2= cc
		age3= 20
		name3= dd
	*/

	// print the type of variable.
	var m2 = 100
	fmt.Printf("The type of m is %T. \n", m2)

	var m3 float64 = 200
	fmt.Printf("The type of m3 is %T, the size of m3's bytes is %d\n", m3, unsafe.Sizeof(m3))

	fmt.Println("----------------------")
	// 2 types of float number.
	num1 := 3.14
	fmt.Println("num1 = ", num1)
	num2 := .14
	fmt.Println("num2 = ", num2)

	num3 := 3.14e2
	fmt.Println("num3 = ", num3)
	num4 := -3.14e2
	fmt.Println("num4 = ", num4)
	num5 := 3.14e-2
	fmt.Println("num5 = ", num5)
	/*
		num1 =  3.14
		num2 =  0.14
		num3 =  314
		num4 =  -314
		num5 =  0.0314
	*/
	fmt.Println("-----------------")
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1 = ", c1, "c2 = ", c2)    // c1 =  97 c2 =  48
	fmt.Printf("c1 = %c, c2 = %c\n", c1, c2) // c1 = a, c2 = 0

	//var c3 byte = '北' // .\var_use.go:115:6: constant 21271 overflows byte
	//fmt.Printf("c3 = %c, %d", c3, c3)

	var c4 int = '北'
	fmt.Printf("c4 = %c, ASICC number is %d.\n", c4, c4) // c4 = 北, ASICC number is 21271
	var c5 = 21271
	fmt.Printf("the byte of number 21271 is %c.\n", c5) // the byte of number 21271 is 北.

	var c6 = 1 + 'a'
	fmt.Println("c6 = ", c6)                         // c6 =  98
	fmt.Printf("the byte of number 98 is %c.\n", c6) // the byte of number 98 is b.

	fmt.Println("-------------")
	var b1 = true
	var b2 = false
	fmt.Println("b1 =", b1, ", b2 = ", b2)
	fmt.Println("b1 size is", unsafe.Sizeof(b1))
}
