package main

import (
	"fmt"
)

func CalcFactory(operate string) func(int, int) int {
	switch operate {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "-":
		return func(a, b int) int {
			return a - b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	case "/":
		return func(a, b int) int {
			if b != 0 {
				return a / b
			}
			return 0
		}
	default:
		return nil
	}
}
func re() (int, int, string) {
	var x, y int
	var expr, op string
	fmt.Scan(&x, &expr)
	/*
		fmt.Print(x, ",", expr, "\n")
	*/
	for i, ch := range expr {
		// fmt.Println(i, " , ", string(ch))
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			op += string(expr[i])
		} else if ch >= '0' && ch <= '9' {
			y = (y << 3) + (y << 1) + int(ch^48)
		} else {
			continue
		}
	}
	return x, y, op
}
func main() {
	fmt.Println("请输入含两个数的四则运算")
	x, y, op := re()
	// fmt.Println(x, y, op)
	calc := CalcFactory(op)
	var ans int = calc(x, y)
	fmt.Printf("%d %s %d = %d\n", x, op, y, ans)
}
