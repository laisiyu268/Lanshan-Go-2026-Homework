package main

import (
	"testing"
)

func TestCalcFactory(t *testing.T) {
	cases := []struct {
		name string
		op   string
		a, b int
		want int
	}{
		{"Add", "+", 1, 2, 3},
		{"Substract", "-", 108, 9, 99},
		{"Multiply", "*", 5, 7, 35},
		{"Divide", "/", 100, 4, 25},
		{"Divided by Zero", "/", 1, 0, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expr := CalcFactory(c.op)
			if expr == nil {
				t.Fatalf("CalcFactory(%q) 返回了 nil", c.op)
			}
			if got := expr(c.a, c.b); got != c.want {
				t.Errorf("%s(%d, %d) = %d, want %d", c.name, c.a, c.b, got, c.want)
			}
		})
	}
}
func TestCalcFactoryUnknownOp(t *testing.T) {
	if expr := CalcFactory("unknown"); expr != nil {
		t.Errorf("CalcFactory(\"unknown\") = %p, want nil", expr)
	}
}
func BenchmarkCalcFactory(b *testing.B) {
	Add := CalcFactory("+")
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

/*
func TestCalcFactory(t *testing.T) {
	Add := CalcFactory("+")
	Sub := CalcFactory("-")
	Mul := CalcFactory("*")
	Div := CalcFactory("/")
	casesAdd := []struct {
		name string
		a, b int
		want int
	}{
		{"正数", 1, 2, 3},
		{"含负数", -1, 1, 0},
		{"零", 0, 0, 0},
	}
	casesOp := map[string]string{
		"+": "Add",
		"-": "Sub",
		"*": "Mul",
		"/": "Div",
	}
	for _, c := range casesAdd {
		t.Run(c.name, func(t *testing.T) {
			if got := Add(c.a, c.b); got != c.want {
				t.Errorf("%s(%d, %d) = %d, want %d", casesOp["+"], c.a, c.b, got, c.want)
			}
		})
	}
	casesGot := []struct {
		op   string
		a, b int
		got  int
		want int
	}{
		{"+", 1, 2, Add(1, 2), 3},
		{"-", 108, 9, Sub(108, 9), 99},
		{"*", 5, 7, Mul(5, 7), 35},
		{"/", 100, 4, Div(100, 4), 25},
		{"/", 1, 0, Div(1, 0), 0},
	}
	for _, c := range casesGot {
		t.Run(c.op, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("%s(%d, %d) = %d, want %d", casesOp[c.op], c.a, c.b, c.got, c.want)
			}
		})
	}
	/*
		gots:=[]int{Add(1, 2), Sub(108, 9),Mul(5,7),Div(100,4),Div(1,0)}

		if gots[0] != 3 {
			t.Errorf("Add(1, 2) = %d, want 3", got1)
		}
		if got != 99 {
			t.Errorf("Sub(108,9) = %d, want 99", got2)
		}
		got3,got4,got5:=Mul(5,7),Div(100,4),Div(1,0)
		if got3 != 35 {
			t.Errorf("Mul(108,9) = %d, want 35", got3)
		}
		if got4 != 25{
			t.Errorf("Div(100, 4) = %d, want 25", got4)
		}
		if got5 !=
	* /
}
*/
