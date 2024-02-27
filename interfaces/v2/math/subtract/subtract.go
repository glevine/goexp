package subtract

import "github.com/glevine/goexp/interfaces/v2/math"

func NewOperator() math.Operator {
	return func(a, b int) int {
		return a - b
	}
}
