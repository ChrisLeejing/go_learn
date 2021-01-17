/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/5 9:57
 */
package main

import "fmt"

func sumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sum, sub := sumAndSub(1, 2)
	fmt.Println("sum:", sum, "sub:", sub)

	_, sub = sumAndSub(1, 3)
	fmt.Println("sub:", sub)
	/*
		sum: 3 sub: -1
		sub: -2
	*/
}
