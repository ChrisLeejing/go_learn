/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/1/6 14:54
 */
package book_demo // 这里属于写一个库, 所以包名可以自定义, 如果需要将文件编译为一个可执行的文件, 则包名为: package main, 这是一个语法规定.

var Book1 string = "《红楼梦》"

func SumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
