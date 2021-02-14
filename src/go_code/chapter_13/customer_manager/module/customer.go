/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/13 17:15
 */
package module

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 使用工厂模式返回一个Customer的实例.
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (this Customer) GetCustomerInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t\n", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
