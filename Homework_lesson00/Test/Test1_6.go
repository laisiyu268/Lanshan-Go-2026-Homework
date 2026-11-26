package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	var s string = "hello"
	var l int = len(s)
	p := point{1, 2}
	fmt.Println(s, l)
	fmt.Println(p)

	fmt.Printf("s=%v\n", s)
	//	%v 默认格式,打印任意类型
	fmt.Printf("l=%d\n", l)
	fmt.Printf("p=%v\n", p)
	// p={1 2} ;
	fmt.Printf("p=%+v\n", p)
	// p={x:1 y:2} ;%+v:结构体时会显示字段名和值
	fmt.Printf("p=%#v\n", p)
	// p=main.point{x:1, y:2} ;%#v:值的Go语法表示

	var name string
	var age int
	fmt.Print("Please input your name and age: ")
	fmt.Scanln(&name, &age)
	fmt.Printf("name=%s age=%d\n", name, age)
}
