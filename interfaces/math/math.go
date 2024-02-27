package math

type Operator func(a, b int) int

func NewOperator(op Operator) Operator {
	return func(a, b int) int {
		return op(a, b)
	}
}
