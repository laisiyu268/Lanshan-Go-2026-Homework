// LV0_8
package main

import "fmt"

func main() {
	m1 := map[string][]int{"a": {1, 2}}
	m2 := map[int]func() string{}
	fmt.Println(m1, m2)

	ages := make(map[string]int)
	v, ok := ages["NotExists"]
	_, _ = v, ok
	_ = m1
	_ = m2

	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 22
	ages["Alice"] = 26
	for k := range ages {
		fmt.Println(k, ages[k])
	}

}
