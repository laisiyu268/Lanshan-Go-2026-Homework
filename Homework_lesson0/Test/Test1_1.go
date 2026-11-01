package main

import (
	"fmt"
)

func main() {
	var a string = "first"
	_ = a
	var b, c int = 1, 2
	fmt.Printf("%d + %d = %d\n", b, c, b+c)
}
