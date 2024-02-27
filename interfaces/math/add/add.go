package add

type Operator func(a, b int) int

func NewOperator() Operator {
	return func(a, b int) int {
		return a + b
	}
}
