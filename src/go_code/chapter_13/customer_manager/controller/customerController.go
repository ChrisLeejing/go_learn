/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/13 17:27
 */
package main

import (
	"fmt"
	"go_code/chapter_13/customer_manager/module"
	"go_code/chapter_13/customer_manager/service"
)

type customerController struct {
	key             string // 获取用户的输入值
	loop            bool   // 定义主界面菜单是否循环
	CustomerService *service.CustomerService
}

func main() {
	// &customerController{}: 结构体初始化.
	controller := &customerController{
		key:             "",
		loop:            true,
		CustomerService: service.NewCustomerService(),
	}
	controller.ShowMenu()
}

func (this *customerController) ShowMenu() {
	for {
		fmt.Println("---------------客户信息管理软件---------------")
		fmt.Printf("\t\t1 添加客户\n")
		fmt.Printf("\t\t2 修改客户\n")
		fmt.Printf("\t\t3 删除客户\n")
		fmt.Printf("\t\t4 客户列表\n")
		fmt.Printf("\t\t5 退    出\n")
		fmt.Printf("请选择(1 ~ 5): ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.AddCustomer()
		case "2":
			this.UpdateCustomer()
		case "3":
			this.DeleteCustomer()
		case "4":
			this.CustomerList()
		case "5":
			this.Exit()
		default:
			fmt.Println("您的输入有误, 请重新输入. ")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您已退出客户关系管理系统.")
}

func (this *customerController) AddCustomer() {
	fmt.Println("---------------添加客户---------------")
	Name := ""
	Gender := ""
	Age := 0
	Phone := ""
	Email := ""
	fmt.Printf("姓名: ")
	fmt.Scanln(&Name)
	fmt.Printf("性别: ")
	fmt.Scanln(&Gender)
	fmt.Printf("年龄: ")
	fmt.Scanln(&Age)
	fmt.Printf("电话: ")
	fmt.Scanln(&Phone)
	fmt.Printf("邮箱: ")
	fmt.Scanln(&Email)
	fmt.Println()

	customer := module.Customer{
		Name:   Name,
		Gender: Gender,
		Age:    Age,
		Phone:  Phone,
		Email:  Email,
	}
	// this.CustomerService = service.NewCustomerService()
	this.CustomerService.AddCustomer(customer)
	fmt.Println("---------------添加完成---------------")
}

func (this *customerController) CustomerList() {
	customerList := this.CustomerService.CustomerList()
	if len(customerList) == 0 {
		fmt.Println("客户列表为空, 请添加客户.")
		fmt.Println()
		return
	}
	fmt.Println("---------------客户列表---------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	for _, customer := range customerList {
		fmt.Println(customer.GetCustomerInfo())
	}
	fmt.Println("---------------客户列表完成---------------")
	fmt.Println()
}

func (this *customerController) UpdateCustomer() {
	id := -1
	fmt.Println("---------------修改客户---------------")
	fmt.Printf("请选择待修改客户编号(-1退出):")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	if this.CustomerService.UpdateCustomer(id) {
		fmt.Println("---------------修改完成---------------")
	} else {
		fmt.Printf("对不起, 没有查找到id为%d的客户, 修改失败.\n", id)
	}
}

func (this *customerController) DeleteCustomer() {
	Id := -1
	DeleteFlag := ""
	fmt.Println("---------------删除客户---------------")
	fmt.Printf("请选择待删除客户编号(-1退出):")
	fmt.Scanln(&Id)

	if Id == -1 {
		return
	}

	for {
		fmt.Printf("确认是否删除(Y/N):")
		fmt.Scanln(&DeleteFlag)
		if DeleteFlag == "Y" || DeleteFlag == "y" {
			break
		} else if DeleteFlag == "N" || DeleteFlag == "n" {
			return
		} else {
			fmt.Println("您的输入有误, 请重新输入(Y/N).")
		}
	}

	if this.CustomerService.DeleteCustomer(Id) {
		fmt.Println("---------------删除完成---------------")
	} else {
		fmt.Printf("对不起, 没有查找到id为%d的客户, 删除失败.\n", Id)
	}

}

func (this *customerController) Exit() {
	ExitFlag := ""
	for {
		fmt.Printf("你确定退出(y/n)该程序吗? 请正确输入(Y/N): ")
		fmt.Scanln(&ExitFlag)
		if ExitFlag == "Y" || ExitFlag == "y" {
			this.loop = false
			break
		} else if ExitFlag == "N" || ExitFlag == "n" {
			return
		}
	}
}
