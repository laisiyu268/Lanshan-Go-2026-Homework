// LV0_4
package main

import "fmt"

func main() {
	var s []int
	prevCap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if cap(s) != prevCap {
			fmt.Printf("len=%2d cap=%2d %v\n",
				len(s), cap(s), s)
			prevCap = cap(s)
		}
	}
}
