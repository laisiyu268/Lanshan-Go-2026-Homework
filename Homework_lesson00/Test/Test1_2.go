package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int32 = 10, 7
	fmt.Printf("a+b=%d , a-b=%d"+
		"a*b=%d , a/b=%d", a+b, a-b, a*b, a/b)
	fmt.Print("\n----------\n")
	fmt.Println(a > b)
	fmt.Println(a == b)
	var c int32 = 5
	fmt.Println(c > a-b)
	var s string = strconv.Itoa(123789)
	n, _ := strconv.Atoi("12345")
	fmt.Print(n, " ", s, "\n-------\n")

	s = "hello"
	fmt.Println(len(s))
	s2 := s + "world"
	fmt.Println(s2)
	fmt.Println(s[0], s[len(s)-1])
	fmt.Println("----------")
	fmt.Println(s2[1:3])
	fmt.Println(s2[:len(s2)])
	fmt.Println(s2[2:])

}
