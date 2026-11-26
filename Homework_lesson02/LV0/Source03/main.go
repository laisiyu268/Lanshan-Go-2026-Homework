// 结构体学习
package main

import (
	"fmt"
)

type Info struct {
	Name string
	Age  int
}
type Lanshaner struct {
	Info
	Group string
}
type Man struct {
	Info
	Group string
}

func (l *Lanshaner) Love() {
	fmt.Printf("%v:Don't forget your love\n", l.Name)
}
func (l *Man) Set1(age int) {
	l.Age = age
	// (*l).Age=age
}
func (l Man) Set2(age int) {
	// 副本,不能实质改变变量在main的值
	l.Age = age
	p := &l
	p.Age = 0 // 副本的地址，一样修改不了
}
func main() {
	kq := &Lanshaner{
		Info: Info{
			Name: "Lsyu",
			Age:  18,
		},
		Group: "后端Go组",
	}
	fmt.Printf("%v\n", kq.Info.Name)
	fmt.Printf("%v\n", kq.Group)
	fmt.Printf("%+v\n", kq)
	fmt.Printf("%#v\n", kq)
	var david *Man = &Man{
		Info: Info{
			Name: "David",
			Age:  0,
		},
		Group: "None",
	}
	kq.Love()
	fmt.Println()
	david.Set2(39)
	fmt.Printf("%+v\n", david)
	david.Set1(40)
	fmt.Printf("%+v\n", david)

}
