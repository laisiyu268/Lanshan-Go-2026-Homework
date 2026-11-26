// 数组
package main

import "fmt"

func main() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3}
	c := [5]int{2: 100, 4: 200}
	var a2 = [3]int32{1, 2}
	var b2 = [...]int{1, 2, 3}
	var c2 [5]int
	fmt.Println(a, b, c)
	fmt.Printf("%#v\n%#v\n%#v\n\n", a2, b2, c2)
	/**/
	d := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	e := [...][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	f := [4][2]int{1: {20, 21}, 3: {40, 41}}
	g := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(d, e, f, g)
	var d2 = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Printf("%#v\n\n", d2)
	/**/
}
