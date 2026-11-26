package calc

func CalculationFactory(operation string) func(int, int) int {
	switch operation {
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

/*
Add:=CalculationFactory("+")
Sub:=CalculationFactory("-")
Mul:=CalculationFactory("*")
Div:=CalculationFactory("/")

*/
