/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/18 10:52
 */
package main

func AddUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
