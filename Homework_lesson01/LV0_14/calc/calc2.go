package calc

import (
	"fmt"
	"testing"
)

func TestAdd1(t *testing.T) {
	Add := CalculationFactory("+")
	got := Add(1, 2)
	if got != 3 {
		t.Errorf("Add(1,2) = %d,want 3", got)

	}
}
func TestAdd2(t *testing.T) {
	// 功能测试 test
	Add := CalculationFactory("+")
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"正数", 1, 2, 3},
		{"含负数", -1, 1, 0},
		{"零", 0, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Add(c.a, c.b); got != c.want {
				t.Errorf("Add(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
			}
		})
	}
}
func BenchmarkAdd(b *testing.B) {
	// 基准测试 benchmark
	Add := CalculationFactory("+")
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
func ExampleAdd() {
	Add := CalculationFactory("+")
	fmt.Println(Add(1, 2))
}
