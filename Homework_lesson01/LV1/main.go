package main

import (
	"fmt"
)

func Work(s1 []int) map[int]int {
	m := make(map[int]int)
	delete(m, 0)
	for _, x := range s1 {
		m[x]++
	}

	return m
}
func main() {
	var n int
	fmt.Scan(&n)
	var s1 = make([]int, 0)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Scan(&x)
		s1 = append(s1, x)
	}
	ans := Work(s1)
	fmt.Printf("Map:\n  %#v\n", ans)
	for x, y := range ans {
		fmt.Printf("数字 %d 有 %d 个\n", x, y)
	}
}
