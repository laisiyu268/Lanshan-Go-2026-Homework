// 结构体练习
package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var ThreeBody Books = Books{
		"ThreeBody",
		"CiXin Liu",
		"science fiction",
		1,
	}
	fmt.Printf("%+v\n", ThreeBody)
	fmt.Printf("%#v\n\n", ThreeBody)
	var C_design Books = Books{
		title:   "C_design",
		subject: "textbook",
	}
	fmt.Printf("%+v\n", C_design)
	fmt.Printf("%#v\n\n", C_design)
}
