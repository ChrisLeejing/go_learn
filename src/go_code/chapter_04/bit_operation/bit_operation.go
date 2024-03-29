/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/8 9:30
 */
package main

import "fmt"

func main() {
	/*
		原码, 反码, 补码
		对于有符号的而言, 用下面6句话解释"原码, 反码, 补码"
		1. 二进制的最高位是符号位, 0: 正数, 1: 负数.
			1 = [0000 0001], -1 = [1000 0001]
		2. 正数的原码, 反码, 补码都一样.
		3. 负数的反码: 它的原码符号位保持不变, 其他位取反(1->0, 0->1).
		4. 负数的补码 = 它的反码 + 1
			1: 原码[0000 0001], 反码[0000 0001], 补码[0000 0001]
		   -1: 原码[1000 0001], 反码[1111 1110], 补码[1111 1111]
		5. 0的反码, 补码都为0.
		6. 在计算机运算时, 都是以补码方式进行运算的.
			1 + 1 = 1 + 1
			1 - 1 = 1 + (-1)
	*/

	/*
		2: 0000 0010		2: 0000 0010      2: 0000 0010       -2: 1111 1110[补码] <- 1111 1101[补码] <- 1000 0010[原码]
		3: 0000 0011		3: 0000 0011      3: 0000 0011        2: 0000 0010
		&: ---------		|: ---------      ^: ---------        ^: ---------
		   0000 0010 = 2	   0000 0011 = 3     0000 0001 = 1       1111 1100[补码] -> 1111 1011[反码] -> 1000 0100[原码] = -4
	*/
	fmt.Println(2 & 3)  // 2
	fmt.Println(2 | 3)  // 3
	fmt.Println(2 ^ 3)  // 1
	fmt.Println(-2 ^ 2) // -4

	/*
		Golang有两个移位运算符, >>(右移), <<(左移), 规则如下:
		1. >>: 低位溢出, 符号位不变, 并用符号位补溢出的高位.
		2. <<: 符号位不变, 低位补0.
	*/

	/*
		a: 0000 0001(补码) >> 2 => 00 0000 00 => 0000 0000, 结果为0.
		b: 1000 0001(原码) => 1111 1110(反码) => 1111 1111(补码).
			即: -1 >> 2 => 1111 1111(补码) >> 2 => 11 1111 11 => 1111 1111(补码) => 1111 1110(反码) => 1000 0001(原码), 结果为-1.
		c: 0000 0001(补码) << 2 => 00 0001 00(补码) => 00 0001 00(反码) => 00 0001 00(原码), 结果为4.
		d: 1000 0001(原码) => 1111 1110(反码) => 1111 1111(补码).
			即: -1 << 2 => 1111 1111(补码) << 2 => 11 1111 00(补码) => 1111 1011(反码) => 1000 0100(原码), 结果为-4.
	*/
	var a int = 1 >> 2
	var b int = -1 >> 2
	var c int = 1 << 2
	var d int = -1 << 2
	fmt.Println("a =", a) // a = 0
	fmt.Println("b =", b) // b = -1
	fmt.Println("c =", c) // c = 4
	fmt.Println("d =", d) // d = -4

}
