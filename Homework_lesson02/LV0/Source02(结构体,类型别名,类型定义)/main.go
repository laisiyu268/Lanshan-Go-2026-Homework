// 结构体,类型别名,类型定义
package main

import (
	"fmt"
	// "unsafe"
)

type Myint int /* 将Myint定义为int类型,定义了一个全新的类型
Myint是一个全新的类型,只是有int的特性 */

type MyInt = int /* MyInt是int的别名
类型别名和原类型是同一种类型。自定义类型是一种全新的类型。
*/
func main() {
	/*
		alice_name:="Alice"
		alice_age:=16
	*/
	type People /*类型名*/ struct {
		Name string
		Age  int
	}
	alice := People{
		Name: "Alice",
		Age:  16,
	}
	var bob People
	bob.Name = "Bob"
	bob.Age = 16 //没有冒号
	david := People{
		"David",
		18,
	}
	fmt.Printf("%v\n", alice)
	fmt.Printf("%+v\n", bob)
	fmt.Printf("%#v\n", david)
	/*
		{Alice 16}
		{Name:Bob Age:16}
		main.People{Name:"David", Age:18}
	*/
}
