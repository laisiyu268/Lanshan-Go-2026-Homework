package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}
	fmt.Println("-------")
	for i := 0; i < 5; i++ {
		fmt.Println(i + '0')
	}
	for {
		fmt.Println("Hello")
		break
	}
	for i, c := range "hello" {
		fmt.Printf("i=%d,c=%c\n", i, c)
	}
	if (7 & 1) == 1 {
		fmt.Println("7是奇数")
	} else {
		fmt.Println("7是偶数")
	}
	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " is 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	var a int
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("other")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
