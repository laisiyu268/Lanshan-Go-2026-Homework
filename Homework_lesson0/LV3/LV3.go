package main

import "fmt"

func add(n int) int {
	var sum int = 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}
func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(add(n))
}
