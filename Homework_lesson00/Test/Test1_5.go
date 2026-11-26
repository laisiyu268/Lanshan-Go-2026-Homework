package main

import (
	"fmt"
)

type mathfunc func(int, int) int

func add(x int, y int) int {
	return x + y
}
func calc(op func(int, int) int, a, b int) int {
	return op(a, b)
}
func getMultiplier() mathfunc {
	return func(a, b int) int {
		return a * b
	}
}
func greet(name string) {
	fmt.Println("Hello, ", name)
}

func main() {
	var name string
	fmt.Scanln(&name)
	greet(name)
	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(add(x, y))
	/*
		Scan可以忽略换行
		Scanln 类似于Scan,但换行\回车立马结束输入
		Scanf 不能换行输入
	*/
	var op mathfunc
	op = add
	res := op(3, 4)
	fmt.Println("3 + 4 =", res)
	multiplier := getMultiplier()
	res = multiplier(3, 7)
	fmt.Println("3 * 7 =", res)
}
