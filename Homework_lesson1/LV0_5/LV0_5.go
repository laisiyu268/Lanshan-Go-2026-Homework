// LV0_5
package main

import "fmt"

func main() {
	//切片共享底层数组
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	s5 := s4[0:cap(s4)]
	s6 := append(s4, 9)
	fmt.Println(s3)
	// [1 2 3 4 5 6 9 8]
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

}
