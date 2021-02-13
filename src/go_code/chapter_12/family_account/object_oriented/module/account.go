/*
 * This is description.
 * @author Chris0710
 * @date 2021/2/13 14:27
 */
package module

import (
	"fmt"
	"strconv"
)

type account struct {
	balance     float64
	income      float64
	outcome     float64
	note        string
	stringSlice []string
}

func NewAccount(balance float64, income float64, outcome float64, note string, stringSlice []string) *account {
	return &account{
		balance:     balance,
		income:      income,
		outcome:     outcome,
		note:        note,
		stringSlice: stringSlice,
	}
}

func (account *account) GetBalance() float64 {
	return account.balance
}

func (account *account) SetBalance(balance float64) {
	account.balance = balance
}

func (account *account) GetIncome() float64 {
	return account.income
}

func (account *account) SetIncome(income float64) {
	account.income = income
}

func (account *account) GetOutcome() float64 {
	return account.outcome
}

func (account *account) SetOutcome(outcome float64) {
	account.outcome = outcome
}

func (account *account) GetNote() string {
	return account.note
}

func (account *account) SetNote(note string) {
	account.note = note
}

func (account *account) GetStringSlice() []string {
	return account.stringSlice
}

func (account *account) SetStringSlice(stringSlice []string) {
	account.stringSlice = stringSlice
}

func (account *account) ShowDetails(detailFlag bool) {
	fmt.Printf("--------------------当前收支明细记录--------------------\n\n")
	if !detailFlag {
		fmt.Println("提示: 目前, 您还没有任何收支明细, 请记录一笔吧.")
	} else {
		fmt.Printf("收支\t收支金额\t账户余额\t说明\n")
		// 遍历[]string
		for _, v := range account.stringSlice {
			fmt.Println(v)
		}
	}
}

func (account *account) ShowIncome() {
	fmt.Printf("本次收入金额: ")
	fmt.Scanln(&account.income)
	fmt.Printf("本次收入说明: ")
	fmt.Scanln(&account.note)
	account.SetIncome(account.income)
	account.SetBalance(account.balance + account.income)
	// float64保留2位小数.
	incomeStr := strconv.FormatFloat(account.income, 'f', 2, 64)
	balanceStr := strconv.FormatFloat(account.balance, 'f', 2, 64)
	account.stringSlice = append(account.stringSlice, "收入\t"+incomeStr+"\t\t"+balanceStr+"\t\t"+account.note)
	detailStr := fmt.Sprintf("收入:%v\t%v", account.income, account.note)
	fmt.Println("本次收支详情: ", detailStr)
}

func (account *account) ShowOutcome() {
	fmt.Printf("本次支出金额: ")
	fmt.Scanln(&account.outcome)
	if account.outcome > account.GetBalance() {
		fmt.Println("对不起, 您的余额不足. ")
	} else {
		fmt.Printf("本次支出说明: ")
		fmt.Scanln(&account.note)
		account.SetOutcome(account.outcome)
		account.SetBalance(account.balance - account.outcome)
		// 注: 相较于case "2", 换一种简洁写法.
		detailStr := fmt.Sprintf("支出%v\t%v\t%v\n", account.outcome, account.balance, account.note)
		fmt.Println("本次收支详情: ", detailStr)
		detailStr2 := fmt.Sprintf("支出\t%v\t\t%v\t\t%v", account.outcome, account.balance, account.note)
		account.stringSlice = append(account.stringSlice, detailStr2)
	}
}

func (account *account) Exit(flag *bool) {
	exitChar := ""
	fmt.Println("你确定退出(y/n)吗? 请输入: ")
	for {
		fmt.Scanln(&exitChar)
		if exitChar == "y" || exitChar == "n" {
			break
		}
		fmt.Println("你的输入有误, 请输入正确的字符: y/n")
	}
	if exitChar == "y" {
		*flag = false
	}
}
