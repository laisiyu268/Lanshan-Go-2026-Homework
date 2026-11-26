package main

import (
	"errors"
	"fmt"
	"sort"
)

type operate func(x, y int) int

func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

type calculateFunc func(x, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}
func main() {
	nums := []int{5, 2, 4}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	add := genCalculator(func(x, y int) int {
		return x + y
	})
	fmt.Println(add(1, 2))
}
