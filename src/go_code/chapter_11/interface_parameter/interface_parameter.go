/*
 * This is description.
 * 需求:
	1. 接口体现多态的两种形式:
	请写一个多态, 可以接收多种参数及多态数组的例子(如: Usb, Phone, Camera).
	* 多态参数
	* 多态数组
 * @author Chris0710
 * @date 2021/2/11 11:25
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
	fmt.Println("phone start...")
}

func (p Phone) stop() {
	fmt.Println("phone stop...")
}

func (c Camera) start() {
	fmt.Println("camera start...")
}

func (c Camera) stop() {
	fmt.Println("camera stop...")
}

func main() {
	// 1. 演示多态参数.
	var phone Usb = Phone{}
	phone.start()
	phone.stop()
	fmt.Println("---------")
	var camera Usb = Camera{}
	camera.start()
	camera.stop()
	/*
	   phone start...
	   phone stop...
	   ---------
	   camera start...
	   camera stop...
	*/

	// 2. 演示多态数组.
	var usbArr [3]Usb
	usbArr[0] = Phone{"iPhone"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"Nikon"}
	fmt.Println("usbArr: ", usbArr)
	/*
		usbArr:  [{iPhone} {xiaomi} {Nikon}]
	*/
}
