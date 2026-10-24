// LV0_3
package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1)
	s = append(s, 2, 3)
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	s2 := make([]int, 5, 5)
	fmt.Println(s2)
	s2 = append(s2, 6)
	fmt.Println(s2, len(s2), cap(s2))

	other := []int{7, 8}
	s = append(s, other...)

}
