package main

import "fmt"

func main() {
	/*
		var s1 []int
		s2 := []int{}
		s3 := make([]int, 0)
		s4 := make([]int, 5)
		s5 := make([]int, 2, 5)
		s6 := []int{1, 2, 3}
	*/
	s := make([]int, 2, 5)
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println(len(s), cap(s))

	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s[:],
		s[2:5],
		s[:5],
		s[5:],
		s[2:5:7])

	/*
		s[low:high:max]
		len=high-low
		cap=max-low
	*/

}
