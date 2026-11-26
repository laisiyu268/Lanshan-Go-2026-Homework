// 指针
package main

import (
	"fmt"
)

func Update(a *int) {
	*a++
}
func main() {
	a := 10
	b := &a
	c := &b
	fmt.Printf("a: %d\t\tptr: %p\n", a, &a)
	fmt.Printf("b: %p\ttype: %T\n", b, b)
	fmt.Printf("c: %p\ttype:%T\tval:%d \n", c, c, **c)
	/**/
	a = 0
	Update(&a)
	fmt.Printf(" a: %d\t\tptr: %p\n", a, &a)
	fmt.Printf(" b: %p\tc: %p\n", b, c)
	fmt.Printf("*b: %d\t*c: %d\t**c: %d\n", *b, *c, **c)
}
