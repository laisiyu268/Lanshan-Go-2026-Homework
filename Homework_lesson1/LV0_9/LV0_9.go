// LV0_9
package main

// import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := map[string]User{"u1": {"Tom", 18}}
	// map里的值不可取地址
	u := users["u1"]
	u.Age = 19
	users["u1"] = u
}
