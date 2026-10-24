// LV0_6
package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Println(dst, n)
}
