package main

import "fmt"

func average(sum int, count int) float64 {
	return float64(sum) / float64(count)
}
func main() {
	var n int = 1
	var sum int = 0
	var count int = 0
	for n != 0 {
		fmt.Print("请输入一个整数(输入0结束): ")
		fmt.Scanln(&n)
		count++
		sum += n
	}
	var ans = average(sum, count-1)
	var f bool = ans >= 60
	if f {
		fmt.Printf("平均成绩为 %.2f, 成绩合格", ans)
	} else {
		fmt.Printf("平均成绩为 %.2f, 成绩不合格", ans)
	}

}
