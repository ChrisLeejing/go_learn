/*
 * This is description.
 * 方式1: 使用面向过程编程实现家庭记账软件.
   		功能1: 完成主菜单, 并且按4可以进行退出.
   		功能2: 可以显示明细和登记收入的功能.
		功能3: 可以显示登记支出的功能.
		功能4: 在输入4退出程序之前, 给出提示: "你确定退出吗? (y/n)", 必须正确输入y/n, 否则循环输入指令, 直到输入y/n为止.
		功能5: 当没有任何收支情况时, 请提示: "目前, 还没有收支明细, 请记录一笔吧"

 * @author Chris0710
 * @date 2021/2/13 9:18
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// n: 用于接收用户在控制台输入的值.
	var n string
	// flag: 用于退出for循环.
	flag := true
	detailFlag := false
	exitChar := ""
	balance := 0.0
	income := 0.0
	outcome := 0.0
	note := ""
	var stringSlice []string
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
			fmt.Printf("--------------------当前收支明细记录--------------------\n\n")
			if !detailFlag {
				fmt.Println("提示: 目前, 您还没有任何收支明细, 请记录一笔吧.")
			} else {
				fmt.Printf("收支\t收支金额\t账户余额\t说明\n")
				// 遍历[]string
				for _, v := range stringSlice {
					fmt.Println(v)
				}
			}
		case "2":
			fmt.Printf("本次收入金额: ")
			fmt.Scanln(&income)
			fmt.Printf("本次收入说明: ")
			fmt.Scanln(&note)
			balance += income
			// float64保留2位小数.
			incomeStr := strconv.FormatFloat(income, 'f', 2, 64)
			balanceStr := strconv.FormatFloat(balance, 'f', 2, 64)
			stringSlice = append(stringSlice, "收入\t"+incomeStr+"\t\t"+balanceStr+"\t\t"+note)
			detailStr := fmt.Sprintf("收入:%v\t%v", income, note)
			fmt.Println("本次收支详情: ", detailStr)
			detailFlag = true
		case "3":
			fmt.Printf("本次支出金额: ")
			fmt.Scanln(&outcome)
			if outcome > balance {
				fmt.Println("对不起, 您的余额不足. ")
				// 注: switch...case默认结尾有break, 不需要显示写出, 但是如果中途退出, 则需要显示写出.
				break
			}
			fmt.Printf("本次支出说明: ")
			fmt.Scanln(&note)
			balance -= outcome
			// 注: 相较于case "2", 换一种简洁写法.
			detailStr := fmt.Sprintf("支出%v\t%v\t%v\n", outcome, balance, note)
			fmt.Println("本次收支详情: ", detailStr)
			detailStr2 := fmt.Sprintf("支出\t%v\t\t%v\t\t%v", outcome, balance, note)
			stringSlice = append(stringSlice, detailStr2)
			detailFlag = true
		case "4":
			fmt.Println("你确定退出(y/n)吗? 请输入: ")
			for {
				fmt.Scanln(&exitChar)
				if exitChar == "y" || exitChar == "n" {
					break
				}
				fmt.Println("你的输入有误, 请输入正确的字符: y/n")
			}
			if exitChar == "y" {
				flag = false
			}
		default:
			fmt.Println("请输入正确的选项.")
		}

		if !flag {
			// 退出程序
			break
		}
	}
	fmt.Println("您已退出了家庭收支记账软件, 谢谢使用. ")
}
