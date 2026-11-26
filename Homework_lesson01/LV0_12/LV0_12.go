package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func CalculationFactory(operation string) func(int, int) int {
	switch operation {
	case "add":
		return func(a, b int) int {
			return a + b
		}

	case "substract":
		return func(a, b int) int {
			return a - b
		}
	case "multiply":
		return func(a, b int) int {
			return a * b
		}
	case "divide":
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
func intsIterator(nums []int) func() (int, bool) {
	i := 0
	return func() (int, bool) {
		if i >= len(nums) {
			return 0, false
		}
		v := nums[i]
		i++
		return v, true
	}
}

/*
	func GetData(data int, callback func(int, int) int) {
		go func() {
			res := data + 2
			fmt.Println(callback(res, res-2))
		}()
	}
*/
func main() {
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
	other := counter()
	fmt.Println(other())

	numAdd := CalculationFactory("add")
	fmt.Println(numAdd(1, 2))
	numSubstract := CalculationFactory("substract")
	fmt.Println(numSubstract(100, 2))

	cnt := intsIterator([]int{1, 2, 3})
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())

	//GetData(7, CalculationFactory("substract"))

}
