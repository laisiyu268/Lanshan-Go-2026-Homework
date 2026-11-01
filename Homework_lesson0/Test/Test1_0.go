package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello,World")
	const N int = 1e5
	const Pi float64 = 3.1415926
	const a float64 = Pi / 2
	fmt.Println(N*N, math.Sin(a))

	var b, c = 1, 2
	if b == 1 {
		c++
		var d int = 3
		_ = d
	}
	//fmt.Println(d)

}
