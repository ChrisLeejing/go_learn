/*
 * This is description.
 方式2: 使用面向对象编程实现家庭记账软件.
   		功能1: 完成主菜单, 并且按4可以进行退出.
   		功能2: 可以显示明细和登记收入的功能.
		功能3: 可以显示登记支出的功能.
		功能4: 在输入4退出程序之前, 给出提示: "你确定退出吗? (y/n)", 必须正确输入y/n, 否则循环输入指令, 直到输入y/n为止.
		功能5: 当没有任何收支情况时, 请提示: "目前, 还没有收支明细, 请记录一笔吧"

 * @author Chris0710
 * @date 2021/2/13 14:30
*/
package main

import (
	"fmt"
	"go_code/chapter_12/family_account/object_oriented/module"
)

func main() {
	fmt.Println("面向过程      ------------->       面向对象")
	account := module.NewAccount(0.0, 0.0, 0.0, "", nil)
	// n: 用于接收用户在控制台输入的值.
	var n string
	// flag: 用于退出for循环.
	flag := true
	detailFlag := false
	for {
		fmt.Println("--------------------家庭收支记账软件--------------------")
		fmt.Printf("\t\t\t1 收支明细\n")
		fmt.Printf("\t\t\t2 登记收入\n")
		fmt.Printf("\t\t\t3 登记支出\n")
		fmt.Printf("\t\t\t4 退    出\n")
		fmt.Printf("请选择(1-4): ")
		fmt.Scanln(&n)

		switch n {
		case "1":
			account.ShowDetails(detailFlag)
		case "2":
			account.ShowIncome()
			detailFlag = true
		case "3":
			account.ShowOutcome()
			detailFlag = true
		case "4":
			account.Exit(&flag)
		default:
			fmt.Println("请输入正确的选项.")
		}

		if !flag {
			// 退出程序
			break
		}
	}

}
