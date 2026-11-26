package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWork(t *testing.T) {
	cases := []struct {
		name string
		sl   []int
		want map[int]int
	}{
		{"常规", []int{1, 1, 2, 3, 3, 3}, map[int]int{1: 2, 2: 1, 3: 3}},
		{"空切片", []int{}, map[int]int{}},
		{"含零和负数", []int{0, -1, -1, 0}, map[int]int{0: 2, -1: 2}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Work(c.sl)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("Work(%v) = %#v, want %#v", c.sl, got, c.want)
			}
		})
	}
}
func BenchmarkWork(b *testing.B) {
	sl := []int{1, 1, 2, 3, 3, 3}
	for Q := 0; Q < b.N; Q++ {
		Work(sl)
	}
}
func ExampleWork() {
	sl := []int{1, 1, 2, 3, 3, 3}
	fmt.Printf("%#v", Work(sl))
	// Output: map[int]int{1:2, 2:1, 3:3}
}

/*
在终端敲 `go test`
*/
