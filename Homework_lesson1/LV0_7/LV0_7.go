// LV0_7
package main

import (
	"fmt"
)

func main() {
	//map
	//1. make
	ages := make(map[string]int)
	//2. 字面量

	grades := map[string]string{
		"Alice": "A",
		"Bob":   "B",
	}

	// 增加或修改
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 22

	// 查询
	aliceAge, exists := ages["Alice"]
	if exists {
		fmt.Printf("Alince's age: %d\n", aliceAge)
	} else {
		fmt.Println("Alice not found.")
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
	fmt.Println(len(ages))
	delete(ages, "Bob")
	fmt.Println(len(ages))
	fmt.Println(grades["Alice"])
}
