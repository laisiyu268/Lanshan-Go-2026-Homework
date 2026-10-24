package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World")
	/**/
	var a = 1
	var b, c = 1, 2

	// var d = 1
	// var e float64
	// var e = 1.0
	// var f = 114514
	// var g = 3 + 1415i
	/**/
	const s = "constant"
	const h = 5e8
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))

	a = 1
	b, c = 1, 2
	if b == 1 {
		c = c + 1
		// d := 3
	}
	// fmt.Printf("%d",d)
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	// a := 1
	if b := a; b == 1 {
		fmt.Print(b)
	}
	// fmt.Print(b)
	// a := 1
	fmt.Println("hello")
}
