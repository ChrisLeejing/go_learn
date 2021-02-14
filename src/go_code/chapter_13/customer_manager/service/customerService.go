/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/13 17:24
 */
package service

import (
	"fmt"
	"go_code/chapter_13/customer_manager/module"
	"strconv"
)

type CustomerService struct {
	Num int
	// CustomerSlice []*module.Customer // 修改为[]module.Customer类型, 因为切片本身就是引用类型, 不过修改slice中值后, 需要重新赋值给slice.
	CustomerSlice []module.Customer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		Num:           0,
		CustomerSlice: []module.Customer{},
	}
}

func (this *CustomerService) AddCustomer(customer module.Customer) {
	this.Num++
	customer.Id = this.Num
	this.CustomerSlice = append(this.CustomerSlice, customer)
}

func (this *CustomerService) CustomerList() []module.Customer {
	return this.CustomerSlice
}

func (this *CustomerService) FindById(id int) module.Customer {
	for _, customer := range this.CustomerSlice {
		if customer.Id != id {
			continue
		}
		return customer
	}

	fmt.Printf("对不起, 没有查找到id为%d的客户.\n", id)
	return module.Customer{}
}

// 通过ID查询到该客户在slice中的索引值.
func (this *CustomerService) FindById2(id int) int {
	Index := -1
	for i, customer := range this.CustomerSlice {
		if customer.Id == id {
			Index = i
		}
	}
	return Index
}

func (this *CustomerService) DeleteCustomer(Id int) bool {
	index := this.FindById2(Id)
	if index == -1 {
		return false
	}

	this.CustomerSlice = append(this.CustomerSlice[:index], this.CustomerSlice[index+1:]...)
	return true
}

func (this *CustomerService) UpdateCustomer(Id int) bool {
	Name := ""
	Gender := ""
	Age := 0
	Phone := ""
	Email := ""
	Index := this.FindById2(Id)
	if Index == -1 {
		return false
	}
	customer := this.CustomerSlice[Index]
	fmt.Printf("姓名(" + customer.Name + "):")
	fmt.Scanln(&Name)
	if Name != "" {
		customer.Name = Name
	}

	fmt.Printf("性别(" + customer.Gender + "):")
	fmt.Scanln(&Gender)
	if Gender != "" {
		customer.Gender = Gender
	}

	fmt.Printf("年龄(" + strconv.Itoa(customer.Age) + "):")
	fmt.Scanln(&Age)
	if Age != 0 {
		customer.Age = Age
	}

	fmt.Printf("电话(" + customer.Phone + "):")
	fmt.Scanln(&Phone)
	if Phone != "" {
		customer.Phone = Phone
	}

	fmt.Printf("邮箱(" + customer.Email + "):")
	fmt.Scanln(&Email)
	if Email != "" {
		customer.Email = Email
	}
	this.CustomerSlice[Index] = customer
	return true
}
