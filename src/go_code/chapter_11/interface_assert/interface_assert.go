/*
 * This is description.
 * 需求:
 * 断言最佳实践1: 给Phone结构体增加一个call方法, 当USB接口接收的是一个Phone结构体时, 在start, end方法基础上, 新增打电话call的功能, 如果是camera结构体, 则正常输出start, end方法

 * @author Chris0710
 * @date 2021/2/11 12:25
 */
package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name + " start...")
}

func (p Phone) stop() {
	fmt.Println(p.Name + " stop...")
}

func (p Phone) call() {
	fmt.Println(p.Name + " call...")
}

func (c Camera) start() {
	fmt.Println(c.Name + " start...")
}

func (c Camera) stop() {
	fmt.Println(c.Name + " stop...")
}

func working(u Usb) {
	u.start()
	if p, ok := u.(Phone); ok {
		p.call()
	}
	u.stop()
}

func main() {
	// 需求1:
	var phone1 Usb = Phone{"iPhone"}
	working(phone1)
	var phone2 Usb = Phone{"xiaomi"}
	working(phone2)
	var camera1 Usb = Camera{"Canon"}
	working(camera1)
	/*
	   iPhone start...
	   iPhone call...
	   iPhone stop...
	   xiaomi start...
	   xiaomi call...
	   xiaomi stop...
	   Canon start...
	   Canon stop...
	*/

}
