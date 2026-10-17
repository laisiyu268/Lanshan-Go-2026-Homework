package main

import "fmt"

func main() {
	var a [10]int
	// 1
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
		fmt.Printf("a[%d]=%d\n", i, a[i])
	}
	// 2
	for i, v := range a {
		fmt.Println("a[", i, "]=", v)
	}
    // 3
    for _,v:=range a{
        fmt.Println(v)
    }
}
