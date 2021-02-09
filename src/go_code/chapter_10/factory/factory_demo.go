/*
 * This is description.
 *
 * @author Chris0710
 * @date 2021/2/9 14:56
 */
package factory

// 情况1: 当结构体的变量首字母为大写时, 通过<package_name>.<struct_name>可以进行调用.
// type People struct {
// 	Name string
// 	Score float64
// }

// 情况2: 当结构体的变量首字母为小写时, 可以通过创建工厂方法进行调用.
// type people struct {
// 	Name  string
// 	Score float64
// }
//
// func NewPeople(n string, s float64) *people {
// 	return &people{
// 		Name:  n,
// 		Score: s,
// 	}
// }

// 情况3: 当结构体的变量首字母为小写时, 且其中的Name字段也是小写.
type people struct {
	name  string
	Score float64
}

func NewPeople(n string, s float64) *people {
	return &people{
		name:  n,
		Score: s,
	}
}

func (p *people) GetName() string {
	return p.name
}
