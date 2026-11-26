package main

import "fmt"

const (
	A = iota
	B
	C
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(A, B, C)
	fmt.Println(KB, MB, GB)
}
